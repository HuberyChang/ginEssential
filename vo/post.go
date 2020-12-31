/*
* @Author: HuberyChang
* @Date: 2020/12/31 16:33
 */

package vo

type CreatePostRequest struct {
	CategoryID uint   `json:"category_id" binding:"required"`
	Title      string `json:"title" binding:"required,max=10"`
	HeadImg    string `json:"head_img"`
	Content    string `json:"content" binding:"required"`
}
