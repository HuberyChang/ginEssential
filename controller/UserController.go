package controller

import (
	"ginEssential/common"
	"ginEssential/dto"
	"ginEssential/model"
	"ginEssential/response"
	"ginEssential/util"
	"log"
	"net/http"

	"golang.org/x/crypto/bcrypt"

	"github.com/gin-gonic/gin"

	"gorm.io/gorm"
)

func Register(ctx *gin.Context) {
	db := common.GetDB()

	// 使用map获取请求的参数
	//var requestMap = make(map[string]string)
	//json.NewDecoder(ctx.Request.Body).Decode(&requestMap)

	// 使用结构体获取参数
	//var requestUser = model.User{}
	//json.NewDecoder(ctx.Request.Body).Decode(&requestUser)
	//name := ctx.PostForm("name")
	//telephone := ctx.PostForm("telephone")
	//password := ctx.PostForm("password")

	// gin框架
	var requestUser = model.User{}
	ctx.Bind(&requestUser)
	name := requestUser.Name
	telephone := requestUser.Telephone
	password := requestUser.Password

	//数据验证
	if len(telephone) != 11 {
		response.Response(ctx, http.StatusUnprocessableEntity, 422, nil, "手机号必须为11位!")
		//ctx.JSON(http.StatusUnprocessableEntity, gin.H{"code": 422, "msg": "手机号必须为11位!"})
		return
	}
	if len(password) < 6 {
		response.Response(ctx, http.StatusUnprocessableEntity, 422, nil, "密码必须大于6位")
		//ctx.JSON(http.StatusUnprocessableEntity, gin.H{"code": 422, "msg": "密码必须大于6位"})
		return
	}

	//如果名称没有传，给一个10位的随机字符串
	if len(name) == 0 {
		name = util.RandomString(10)
	}
	log.Println(name, telephone, password)

	//判断手机号
	if isTelephoneExist(db, telephone) {
		response.Response(ctx, http.StatusUnprocessableEntity, 422, nil, "用户已经存在")
		//ctx.JSON(http.StatusUnprocessableEntity, gin.H{"code": 422, "msg": "用户已经存在"})
		return
	}

	//创建用户
	hasePassword, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err != nil {
		response.Response(ctx, http.StatusUnprocessableEntity, 422, nil, "密码加密错误")
		//ctx.JSON(http.StatusInternalServerError, gin.H{"code": 500, "msg": "密码加密错误"})
		return
	}
	newUser := model.User{
		Name:      name,
		Telephone: telephone,
		Password:  string(hasePassword),
	}
	db.Create(&newUser)

	//发放token
	token, err := common.ReleaseToken(newUser)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"code": 500, "msg": "系统异常"})
		log.Printf("token generate error: %v", err)
		return

	}

	//返回结果
	response.Success(ctx, gin.H{"token": token}, "用户注册成功")
	//ctx.JSON(200, gin.H{
	//	"code":    200,
	//	"message": "用户注册成功",
	//})
}

func Login(ctx *gin.Context) {
	DB := common.GetDB()

	var requestUser = model.User{}
	ctx.Bind(&requestUser)
	telephone := requestUser.Telephone
	password := requestUser.Password

	// 数据验证
	if len(telephone) != 11 {
		response.Response(ctx, http.StatusUnprocessableEntity, 422, nil, "手机号必须为11位!")
		// ctx.JSON(http.StatusUnprocessableEntity, gin.H{"code": 422, "msg": "手机号必须为11位!"})
		return
	}
	if len(password) < 6 {
		response.Response(ctx, http.StatusUnprocessableEntity, 422, nil, "密码必须大于6位")
		// ctx.JSON(http.StatusUnprocessableEntity, gin.H{"code": 422, "msg": "密码必须大于6位"})
		return
	}

	// 判断手机是否存在
	var user model.User
	DB.Where("telephone=?", telephone).First(&user)
	if user.ID == 0 {
		ctx.JSON(http.StatusUnprocessableEntity, gin.H{"code": 422, "msg": "用户不存在！"})
		return
	}

	// 判断密码是否正确
	if err := bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(password)); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"code": 400, "msg": "密码错误"})
		return
	}

	// 发放token
	token, err := common.ReleaseToken(user)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"code": 500, "msg": "系统异常"})
		log.Printf("token generate error: %v", err)
		return

	}

	// 返回结果
	response.Success(ctx, gin.H{"token": token}, "登录成功")
	// ctx.JSON(200, gin.H{
	// 	"code": 200,
	//	"data": gin.H{"token": token},
	//	"msg":  "登录成功",
	//})

}

func Info(ctx *gin.Context) {
	user, _ := ctx.Get("user")

	ctx.JSON(http.StatusOK, gin.H{"code": 200, "data": gin.H{"user": dto.ToUserDto(user.(model.User))}})
}

func isTelephoneExist(db *gorm.DB, telephone string) bool {
	var user model.User
	db.Where("telephone = ?", telephone).First(&user)
	if user.Model.ID != 0 {
		return true
	}
	return false
}
