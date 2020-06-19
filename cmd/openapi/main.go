package main

import (
	//"flag"
	"fmt"
	"github.com/tti6o/clickhouse_api_demo/api/router"
	"github.com/tti6o/clickhouse_api_demo/database"
	//"github.com/tti6o/clickhouse_api_demo/config"
)

func main() {
	db,err := database.InitDB()
	if err != nil{
		fmt.Println("InitDBf failed.", err.Error())
		return
	}
	defer db.Close()

	router := router.InitRouter()
	router.Run(":8000")
}
