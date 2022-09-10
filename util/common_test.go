package util

import (
	"av-spider/dao"
	"av-spider/entity"
	"fmt"
	"reflect"
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

func getConfig(s string) string {
	return s
}

func TestReflect(t *testing.T) {
	a := getConfig
	//fmt.Printf("%v", a)
	aa := reflect.ValueOf(a)

	fmt.Printf("%v", aa.String())
}
