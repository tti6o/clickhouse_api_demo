package apis

import (
	"github.com/gin-gonic/gin"
	"github.com/tti6o/clickhouse_api_demo/api/models"
	"net/http"
)

type HitUserReq struct {
	UserID   int64 `form:"user_id" binding:"required"`
}

//获取指定userid的用户列表
func Users(c *gin.Context) {
	hitUserReq := HitUserReq{}
	err := c.ShouldBindQuery(&hitUserReq)
	if err != nil {
		c.JSON(http.StatusOK, gin.H{
			"code":    -1,
			"message": "param error",
		})
	}
	var hitUser models.HitUser
	users, err2 := hitUser.GetHitUsers(hitUserReq.UserID)

	if err2 != nil {
		c.JSON(http.StatusOK, gin.H{
			"code":    -2,
			"message": "not exist",
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"code": 0,
		"data":  users,
	})
}

//获取记录数>1000的用户ID列表（从高到低排序）
func GetHighUsers(c *gin.Context) {
	var hitUser models.HitUser
	users, err2 := hitUser.GetHighUsers()

	if err2 != nil {
		c.JSON(http.StatusOK, gin.H{
			"code":    -2,
			"message": "not exist",
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"code": 0,
		"data":  users,
	})
}