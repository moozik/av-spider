package controller

import (
	"av-spider/dao"
	"av-spider/util"
	"fmt"
	"github.com/gin-gonic/gin"
	"net/url"
)

func Index(c *gin.Context) {
	keyword := c.Param("keyword")
	pageNum := util.ToInt(c.Param("pageNum"))
	whereList := dao.SearchWhereBuild(keyword)
	avList := dao.SelectAvList(whereList, pageNum)

	var pageRoot string
	if keyword != "" {
		pageRoot = fmt.Sprintf("/search/%s", url.QueryEscape(keyword))
	} else {
		pageRoot = ""
	}

	c.HTML(200, "index.html", gin.H{
		"data": gin.H{
			"av_list": avList,
		},
		"frame_data": gin.H{
			"title":       "title",
			"placeholder": "placeholder",
			"origin_link": util.GetUrl("search", url.QueryEscape(keyword), pageNum),
			"page":        util.Pagination(pageRoot, pageNum, 100, 0),
		},
	})
}
