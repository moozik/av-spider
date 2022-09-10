package main

import (
	"av-spider/util"
	"av-spider/website"
	"av-spider/website/controller"
	"fmt"
	"github.com/gin-gonic/gin"
	"io"
	"os"
	"time"
)

func main() {
	if util.IsDebugging() {
		//命令行颜色
		gin.ForceConsoleColor()
		gin.SetMode(gin.DebugMode)
		// 如果需要同时将日志写入文件和控制台，请使用以下代码。
		gin.DefaultWriter = io.MultiWriter(os.Stdout)
	} else {
		//禁用命令行颜色
		gin.DisableConsoleColor()
		gin.SetMode(gin.ReleaseMode)

		if !util.PathExists("logs") {
			os.Mkdir("logs", os.ModePerm)
		}
		// 记录到文件。
		f, _ := os.Create(fmt.Sprintf("logs/gin.log.%s", time.Now().Format("20060102")))
		// 如果需要同时将日志写入文件和控制台，请使用以下代码。
		gin.DefaultWriter = io.MultiWriter(f, os.Stdout)
	}

	engine := gin.New()
	//html := template.Must(template.ParseFiles("file1", "file2"))
	//engine.SetHTMLTemplate(html)
	//engine.Delims("{%", "%}")
	engine.SetFuncMap(website.FuncMap)
	engine.Static("/static", "website/static")
	//engine.LoadHTMLGlob("website/templates/*")
	engine.LoadHTMLGlob("website/templates2/*")
	engine.Use(gin.Logger(), gin.Recovery())

	engine.GET("/test", func(c *gin.Context) {
		c.String(200, "ok")
	})
	{
		engine.GET("/", controller.Index)
		engine.GET("/page/:page_num", controller.Index)
		engine.GET("/search/:keyword", controller.Index)
		engine.GET("/search/:keyword/page/:pageNum", controller.Index)
	}
	fmt.Println("running")
	engine.Run(":5000")
}
