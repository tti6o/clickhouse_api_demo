package models

import (
	"github.com/tti6o/clickhouse_api_demo/database"
	"log"
)

//type HitUser struct {
//	ID       int64  `json:"id"`       // 列名为 `id`
//	Username string `json:"username"` // 列名为 `username`
//	Password string `json:"password"` // 列名为 `password`
//}

type HitUser struct {
	ID       	int64  `json:"WatchID"`
	UserID 	 	int64  `json:"UserID"`
	Age			int64  `json:"Age"`
	Sex			int64  `json:"Sex"`
	ClientIP 	string `json:"ClientIP"`
	WindowName 	string	`json:"WindowName"`
}

type HighUser struct {
	UserID 	 	int64  `json:"UserID"`
	Count 	 	int64  `json:"Count"`
}

////添加
//func (user HitUser) Insert() (id int64, err error) {
//
//	//添加数据
//	result := orm.Eloquent.Create(&user)
//	id =user.ID
//	if result.Error != nil {
//		err = result.Error
//		return
//	}
//	return
//}

//列表
func (user *HitUser) GetHitUsers(UserID int64) (users []HitUser, err error) {
	db := database.GetDB()
	rows, err1 := db.Query("SELECT WatchID,UserID,Age,Sex,ClientIP,WindowName FROM hits_v1 WHERE UserID = ?",UserID)
	if err1 != nil {
		log.Fatal(err1)
		return nil,err1
	}
	defer rows.Close()
	log.Println(rows.Columns())

	for rows.Next() {
		hitUser := HitUser{}
		if err := rows.Scan(&hitUser.ID,&hitUser.UserID,&hitUser.Age,&hitUser.Sex,&hitUser.ClientIP,&hitUser.WindowName); err != nil {
			log.Fatal(err)
		}
		log.Println("hitUser",hitUser)
		users = append(users,hitUser)
	}

	if err := rows.Err(); err != nil {
		log.Fatal(err)
		return nil,err
	}
	return
}

func (user *HitUser) GetHighUsers() (users []HighUser, err error) {
	db := database.GetDB()
	rows, err1 := db.Query("SELECT UserID, COUNT(1) AS `COUNT` FROM  tutorial.hits_v1 GROUP BY UserID HAVING `COUNT` > 1000 ORDER BY `COUNT` DESC")
	if err1 != nil {
		log.Fatal(err1)
		return nil,err1
	}
	defer rows.Close()
	log.Println(rows.Columns())

	for rows.Next() {
		highUser := HighUser{}
		if err := rows.Scan(&highUser.UserID,&highUser.Count,); err != nil {
			log.Fatal(err)
		}
		log.Println("highUser",highUser)
		users = append(users,highUser)
	}

	if err := rows.Err(); err != nil {
		log.Fatal(err)
		return nil,err
	}
	return
}

////修改
//func (user *HitUser) Update(id int64) (updateUser HitUser, err error) {
//
//	if err = orm.Eloquent.Select([]string{"id", "username"}).First(&updateUser, id).Error; err != nil {
//		return
//	}
//
//	//参数1:是要修改的数据
//	//参数2:是修改的数据
//	if err = orm.Eloquent.Model(&updateUser).Updates(&user).Error; err != nil {
//		return
//	}
//	return
//}
//
////删除数据
//func (user *HitUser) Destroy(id int64) (Result HitUser, err error) {
//
//	if err = orm.Eloquent.Select([]string{"id"}).First(&user, id).Error; err != nil {
//		return
//	}
//
//	if err = orm.Eloquent.Delete(&user).Error; err != nil {
//		return
//	}
//	Result = *user
//	return
//}