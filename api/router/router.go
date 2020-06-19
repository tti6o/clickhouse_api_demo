package router

import (
	"github.com/gin-gonic/gin"
	. "github.com/tti6o/clickhouse_api_demo/api/apis"
)

func InitRouter() *gin.Engine {
	router := gin.Default()

	router.GET("/users", Users)

	router.GET("/high_user_ids", GetHighUsers)

	return router
}