package database

import (
	"database/sql"
	"github.com/ClickHouse/clickhouse-go"
	//"github.com/jinzhu/gorm"
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

	//_, err = db.Exec(`
	//	CREATE TABLE IF NOT EXISTS example (
	//		country_code FixedString(2),
	//		os_id        UInt8,
	//		browser_id   UInt8,
	//		categories   Array(Int16),
	//		action_day   Date,
	//		action_time  DateTime
	//	) engine=Memory
	//`)
	//
	//if err != nil {
	//	log.Fatal(err)
	//}
	//var (
	//	tx, _   = db.Begin()
	//	stmt, _ = tx.Prepare("INSERT INTO example (country_code, os_id, browser_id, categories, action_day, action_time) VALUES (?, ?, ?, ?, ?, ?)")
	//)
	//defer stmt.Close()
	//
	//for i := 0; i < 100; i++ {
	//	if _, err := stmt.Exec(
	//		"RU",
	//		10+i,
	//		100+i,
	//		clickhouse.Array([]int16{1, 2, 3}),
	//		time.Now(),
	//		time.Now(),
	//	); err != nil {
	//		log.Fatal(err)
	//	}
	//}
	//
	//if err := tx.Commit(); err != nil {
	//	log.Fatal(err)
	//}

	rows, err := db.Query("SELECT * FROM hits_v1 limit 1")
	if err != nil {
		log.Fatal(err)
		return db,err
	}
	defer rows.Close()
	log.Println(rows.Columns())

	//for rows.Next() {
	//	var (
	//		country               string
	//		os, browser           uint8
	//		categories            []int16
	//		actionDay, actionTime time.Time
	//	)
	//	if err := rows.Scan(&country, &os, &browser, &categories, &actionDay, &actionTime); err != nil {
	//		log.Fatal(err)
	//	}
	//	log.Printf("country: %s, os: %d, browser: %d, categories: %v, action_day: %s, action_time: %s", country, os, browser, categories, actionDay, actionTime)
	//}

	//if err := rows.Err(); err != nil {
	//	log.Fatal(err)
	//	return db,err
	//}

	//if _, err := db.Exec("DROP TABLE example"); err != nil {
	//	log.Fatal(err)
	//}
	return db,nil
}