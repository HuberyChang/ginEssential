package common

import (
	"fmt"

	"github.com/spf13/viper"
	"gorm.io/gorm"

	"gorm.io/driver/mysql"
)

var DB *gorm.DB

func InitDB() {
	//host := "localhost"
	//port := "3306"
	//database := "ginessential"
	//username := "root"
	//pasword := "root"
	//charset := "utf8mb4"
	//loc := "Local"
	dsn := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=%s&parseTime=True&loc=%s",
		//username,
		//pasword,
		//host,
		//port,
		//database,
		//charset,
		//loc
		viper.GetString("dataSource.user"),
		viper.GetString("dataSource.password"),
		viper.GetString("dataSource.host"),
		viper.GetString("dataSource.port"),
		viper.GetString("dataSource.database"),
		viper.GetString("dataSource.charset"),
		viper.GetString("dataSource.loc"),
	)

	db, err := gorm.Open(mysql.New(mysql.Config{
		DSN:                       dsn,
		DefaultStringSize:         256,
		DisableDatetimePrecision:  true,
		DontSupportRenameIndex:    true,
		DontSupportRenameColumn:   true,
		SkipInitializeWithVersion: false,
	}), &gorm.Config{})
	if err != nil {
		panic("failed to connect database,err:" + err.Error())
	}
	DB = db
}

func GetDB() *gorm.DB {
	DB.Debug()
	return DB
}
