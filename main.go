package main

import (
	"lottery/handler"
	"lottery/util"
	"net/http"

	"github.com/gin-gonic/gin"
)

func Init() {
	util.InitLog("log")
}

func main() {
	Init()

	// GIN 线上发布模式
	gin.SetMode(gin.ReleaseMode)

	// 修改静态资源不需要重启GIN，刷新页面即可
	router := gin.Default()
	router.Static("/js", "views/js")
	router.Static("/img", "views/img")
	router.StaticFile("favicon.ico", "views/img/dqq.png")
	router.LoadHTMLFiles("views/lottery.html")

	router.GET("/", func(ctx *gin.Context) {
		ctx.HTML(http.StatusOK, "lottery.html", nil)
	})

	router.GET("/gifts", handler.GetAllLotteryGifts)

	router.Run("localhost:5678")
}

// command: go run ./main.go
