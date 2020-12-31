/*
* @Author: HuberyChang
* @Date: 2020/12/31 14:56
 */

package vo

type CreateCategoryRequest struct {
	Name string `json:"name" binding:"required"`
}
