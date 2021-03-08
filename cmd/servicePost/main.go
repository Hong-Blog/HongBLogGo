package main

import (
	"github.com/gin-gonic/gin"
	"hong-blog/router"
)

func main() {
	engine := gin.Default()

	router.SetupPostRouter(engine)

	addr := ":80"
	if gin.Mode() == gin.DebugMode {
		addr = ":8083"
	}
	_ = engine.Run(addr)
}
