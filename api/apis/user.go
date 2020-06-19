package apis

import (
	"github.com/gin-gonic/gin"
	model "github.com/tti6o/clickhouse_api_demo/api/models"
	"net/http"
)

type HitUserReq struct {
	UserID   int64 `form:"user_id" binding:"required"`
}

//列表数据
func Users(c *gin.Context) {
	hitUserReq := HitUserReq{}
	err := c.ShouldBindQuery(&hitUserReq)
	if err != nil {
		c.JSON(http.StatusOK, gin.H{
			"code":    -1,
			"message": "param error",
		})
	}

	var hitUser model.HitUser
	users, err2 := hitUser.GetHitUsers(hitUserReq.UserID)

	if err2 != nil {
		c.JSON(http.StatusOK, gin.H{
			"code":    -2,
			"message": "not exist",
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"code": 1,
		"data":  users,
	})
}

//列表数据
func GetHighUsers(c *gin.Context) {
	var hitUser model.HitUser
	users, err2 := hitUser.GetHighUsers()

	if err2 != nil {
		c.JSON(http.StatusOK, gin.H{
			"code":    -2,
			"message": "not exist",
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"code": 1,
		"data":  users,
	})
}
//
////添加数据
//func Store(c *gin.Context) {
//	var user model.User
//	user.Username = c.Request.FormValue("username")
//	user.Password = c.Request.FormValue("password")
//	id, err := user.Insert()
//
//	if err != nil {
//		c.JSON(http.StatusOK, gin.H{
//			"code":    -1,
//			"message": "添加失败",
//		})
//		return
//	}
//	c.JSON(http.StatusOK, gin.H{
//		"code":  1,
//		"message": "添加成功",
//		"data":    id,
//	})
//}
//
////修改数据
//func Update(c *gin.Context) {
//	var user model.User
//	id, err := strconv.ParseInt(c.Param("id"), 10, 64)
//	user.Password = c.Request.FormValue("password")
//	result, err := user.Update(id)
//	if err != nil || result.ID == 0 {
//		c.JSON(http.StatusOK, gin.H{
//			"code":    -1,
//			"message": "修改失败",
//		})
//		return
//	}
//	c.JSON(http.StatusOK, gin.H{
//		"code":  1,
//		"message": "修改成功",
//	})
//}
//
////删除数据
//func Destroy(c *gin.Context) {
//	var user model.User
//	id, err := strconv.ParseInt(c.Param("id"), 10, 64)
//	result, err := user.Destroy(id)
//	if err != nil || result.ID == 0 {
//		c.JSON(http.StatusOK, gin.H{
//			"code":    -1,
//			"message": "删除失败",
//		})
//		return
//	}
//	c.JSON(http.StatusOK, gin.H{
//		"code":  1,
//		"message": "删除成功",
//	})
//}