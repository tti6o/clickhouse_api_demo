package main

import (
	"fmt"
	"github.com/360EntSecGroup-Skylar/excelize"
	"github.com/parnurzeal/gorequest"
	"net/http"
	"strconv"
	"time"
)

type UserData struct {
	WatchID     int64  `json:"WatchID"`
	UserID 	 	int64  `json:"UserID"`
	Age			int64  `json:"Age"`
	Sex			int64  `json:"Sex"`
	ClientIP 	string `json:"ClientIP"`
	WindowName 	string	`json:"WindowName"`
}

type GetUserResp struct {
	Code  int64       	`json:"code"`
	Data  []UserData		`json:"data"`
}

type HighUserData struct {
	UserID 	 	int64  `json:"UserID"`
	Count		int64  `json:"Count"`
}

type GetHighUserResp struct {
	Code  int64       		`json:"code"`
	Data  []HighUserData	`json:"data"`
}


func main() {
	f := excelize.NewFile()
	GetUsers(f)
	index := f.NewSheet("Sheet2")
	GetHighUsers(f)
	f.SetActiveSheet(index)
	if err := f.SaveAs("PivotTable.xlsx"); err != nil {
		fmt.Println(err)
	}
}

//获取指定userid的用户列表
func GetUsers(f *excelize.File ) {
	getUserUrl := "http://47.56.239.148:8000/users?user_id=610708775678702928"
	userResp := GetUserResp{}
	ret, _, errs := gorequest.New().
		Timeout(time.Second * 5).
		Get(getUserUrl).
		EndStruct(&userResp)
	if errs != nil || ret.StatusCode != http.StatusOK {
		fmt.Println(errs)
		return
	}
	fmt.Println("GetUserResp:",userResp)
	list := []string {"WatchID","UserID","Age","Sex","ClientIP","WindowName",}
	for i:= 0; i < len(list); i++ {
		cell, _ := excelize.CoordinatesToCellName(i+1,1)
		f.SetCellValue("Sheet1", cell, list[i])
	}
	for i:= 0; i < len(userResp.Data); i++ {
		data := userResp.Data[i]
		rowNum := i + 2
		cell1, _ := excelize.CoordinatesToCellName(1,rowNum)
		f.SetCellValue("Sheet1", cell1, strconv.FormatInt(data.WatchID,10))
		cell2, _ := excelize.CoordinatesToCellName(2,rowNum)
		f.SetCellValue("Sheet1", cell2, strconv.FormatInt(data.UserID,10))
		cell3, _ := excelize.CoordinatesToCellName(3,rowNum)
		f.SetCellValue("Sheet1", cell3, data.Age)
		cell4, _ := excelize.CoordinatesToCellName(4,rowNum)
		f.SetCellValue("Sheet1", cell4, data.Sex)
		cell5, _ := excelize.CoordinatesToCellName(5,rowNum)
		f.SetCellValue("Sheet1", cell5, data.ClientIP)
		cell6, _ := excelize.CoordinatesToCellName(6,rowNum)
		f.SetCellValue("Sheet1", cell6, data.WindowName)
	}
	if err := f.SaveAs("PivotTable.xlsx"); err != nil {
		fmt.Println(err)
	}
}

//获取记录数>1000的用户ID列表（从高到低排序）
func GetHighUsers(f *excelize.File) {
	url := "http://47.56.239.148:8000/high_user_ids"
	userResp := GetHighUserResp{}
	ret, _, errs := gorequest.New().
		Timeout(time.Second * 5).
		Get(url).
		EndStruct(&userResp)
	if errs != nil || ret.StatusCode != http.StatusOK {
		fmt.Println(errs)
		return
	}
	fmt.Println("GetHighUserResp:",userResp)
	list := []string {"UserID","Count",}
	for i:= 0; i < len(list); i++ {
		cell, _ := excelize.CoordinatesToCellName(i+1,1)
		f.SetCellValue("Sheet2", cell, list[i])
	}
	for i:= 0; i < len(userResp.Data); i++ {
		data := userResp.Data[i]
		rowNum := i + 2
		cell1, _ := excelize.CoordinatesToCellName(1,rowNum)
		f.SetCellValue("Sheet2", cell1, strconv.FormatInt(data.UserID,10))
		cell2, _ := excelize.CoordinatesToCellName(2,rowNum)
		f.SetCellValue("Sheet2", cell2, data.Count)
	}
}