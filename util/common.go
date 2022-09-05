package util

import (
	"av-spider/define"
	"fmt"
	"github.com/gin-gonic/gin"
	_ "github.com/mattn/go-sqlite3"
	"math"
	"os"
	"strconv"
)

// PathExists 判断所给路径文件/文件夹是否存在
func PathExists(path string) bool {
	_, err := os.Stat(path)
	if err == nil {
		return true
	}
	//IsNotExist来判断，是不是不存在的错误
	if os.IsNotExist(err) { //如果返回的错误类型使用os.isNotExist()判断为true，说明文件或者文件夹不存在
		return false
	}
	return false //如果有错误了，但是不是不存在的错误，所以把这个错误原封不动的返回
}

func IsDebugging() bool {
	return os.Getenv("OK_DEBUG") == "1"
}

func ToInt(s string) int {
	i, _ := strconv.Atoi(s)
	return i
}

func GetUrl(pageType, keyword string, pageNum int) string {
	ret := "https://avmoo.click/cn"
	if pageType != "" {
		ret = fmt.Sprintf("%s/%s", ret, pageType)
	}
	if keyword != "" {
		ret = fmt.Sprintf("%s/%s", ret, keyword)
	}
	if pageNum > 0 {
		ret = fmt.Sprintf("%s/page/%d", ret, pageNum)
	}
	return ret
}

func GetLocalUrl(pageType, keyword string, pageNum int) string {
	ret := fmt.Sprintf("http://%s:%d", define.LocalIp, define.ServerPort)
	if pageType == "popular" {
		return ret
	}
	if pageType != "" {
		ret = fmt.Sprintf("%s/%s", ret, pageType)
	}
	if keyword != "" {
		ret = fmt.Sprintf("%s/%s", ret, keyword)
	}
	if pageNum > 0 {
		ret = fmt.Sprintf("%s/page/%d", ret, pageNum)
	}
	return ret
}

func Pagination(pageRoot string, pageNum, count, pageLimit int) gin.H {
	if pageLimit == 0 {
		pageLimit = define.PageLimit
	}
	pageCount := int(math.Ceil(float64(count / pageLimit)))
	totalMax := 8
	p1 := pageNum - totalMax
	p2 := pageNum + totalMax

	var pageList []int
	for i := p1; 0 < i && i < p2+1; i += 1 {
		pageList = append(pageList, i)
	}

	pageLeft := 0
	pageRight := 0
	if pageCount != 0 && pageNum != pageCount {
		pageRight = pageNum + 1
	}
	if pageNum != 1 {
		pageLeft = pageNum - 1
	}

	pageHead := 0
	pageTail := 0
	if len(pageList) > 0 {
		if pageList[0] != 1 {
			pageHead = 1
		}
		if pageList[len(pageList)-1] != pageCount {
			pageTail = pageCount
		}
	}
	return gin.H{
		"now":      pageNum,
		"left":     pageLeft,
		"right":    pageRight,
		"list":     pageList,
		"head":     pageHead,
		"tail":     pageTail,
		"pageRoot": pageRoot,
		"count":    count,
	}
}
