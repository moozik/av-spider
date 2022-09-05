package util

import (
	"av-spider/dao"
	"av-spider/entity"
	"fmt"
	"testing"
)

func TestGetDb(t *testing.T) {
	db := dao.GetDb()
	var avList []entity.AvList
	db.Table("av_list").Limit(10).Scan(&avList)
	fmt.Printf("row:%+v, err:%+v", avList, db.Error)
}

func TestSelectAvList(t *testing.T) {
	ret := dao.SelectAvList([]string{}, 1)
	fmt.Println(ret)
}
