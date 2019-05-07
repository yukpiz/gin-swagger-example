package main

import (
	"log"

	"github.com/gin-gonic/gin"
	ginSwagger "github.com/swaggo/gin-swagger"
	"github.com/swaggo/gin-swagger/swaggerFiles"

	_ "github.com/yukpiz/gin-swagger-example/docs"
)

// @title gin-swaggerサンプルAPI
// @version 1
// @description gin-swaggerを試すためのAPI
// @termsOfService http://swagger.io/terms/

// @contact.name @yukpiz
// @contact.url https://twitter.com/yukpiz
// @contact.email yukpiz@gmail.com

// @license.name WTFPL
// @license.url http://www.wtfpl.net/

// @host localhost:8081
// @BasePath /
func main() {
	r := gin.Default()
	config := &ginSwagger.Config{
		URL: "http://localhost:8081/swagger/doc.json",
	}

	r.GET("/swagger/*any", ginSwagger.CustomWrapHandler(config, swaggerFiles.Handler))

	r.GET("/hello1", Hello1)

	r.Run(":8081")
}

// Hello1 godoc
// @summary APIのサマリー説明
// @description ログに"Hello"と出力します
// @tags タグ1
// @accept application/json
// @produce application/json
// @param user_id query int true "パラメータの説明"
// @param merchant_id query int true "パラメータの説明"
// @success 200 {string} string "成功レスポンスの説明"
// @failure 400 {string} string "失敗レスポンスの説明"
// @router /hello [get]
func Hello1(gc *gin.Context) {
	log.Println("hello!")
}
