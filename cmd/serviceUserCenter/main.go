package main

import (
	_ "hong-blog/docs"
	"hong-blog/router"
)

// @title 用户中心API
// @version 1.0
func main() {
	engine := router.SetupUserRouter()

	//url := ginSwagger.URL("./doc.json")
	//engine.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler, url))

	_ = engine.Run(":8081")
}
