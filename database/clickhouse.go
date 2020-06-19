package database

import (
	"database/sql"
	"github.com/ClickHouse/clickhouse-go"
	"fmt"
	"log"
)

var db *sql.DB

func GetDB() *sql.DB {
	return db
}

func InitDB() (*sql.DB,error) {
	var err error
	db, err = sql.Open("clickhouse", "tcp://127.0.0.1:9000?username=default&password=root&database=tutorial")
	if err != nil {
		log.Fatal(err)
		return db,err
	}
	if err := db.Ping(); err != nil {
		if exception, ok := err.(*clickhouse.Exception); ok {
			fmt.Printf("[%d] %s \n%s\n", exception.Code, exception.Message, exception.StackTrace)
		} else {
			fmt.Println(err)
		}
		return db,err
	}
	//test
	//rows, err := db.Query("SELECT * FROM hits_v1 limit 1")
	//if err != nil {
	//	log.Fatal(err)
	//	return db,err
	//}
	//defer rows.Close()
	//log.Println(rows.Columns())

	return db,nil
}