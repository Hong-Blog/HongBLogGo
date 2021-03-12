package main

import (
	"github.com/gin-gonic/gin"
	_ "hong-blog/docs"
	"hong-blog/router"
)

// @title 用户中心API
// @version 1.0
func main() {
	engine := router.SetupUserRouter()

	//url := ginSwagger.URL("./doc.json")
	//engine.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler, url))

	addr := ":80"
	if gin.Mode() == gin.DebugMode {
		addr = ":8081"
	}

	_ = engine.Run(addr)
}
