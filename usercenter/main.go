package main

import (
	"usercenter/router"
)

func main() {
	engine := router.SetupRouter()
	_ = engine.Run(":8081")
}
