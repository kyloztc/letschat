package main

import (
	"letschat/common"
	usermodule "letschat/user_module"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

func regsiterRouter() *gin.Engine {
	router := gin.Default()
	corsConfig := cors.Default()
	router.Use(corsConfig)
	userHandler := new(usermodule.UserHandler)
	userHandler.AddTo(router)
	return router
}

func main() {
	configPath := "conf/properties.yaml"
	common.Init(configPath)

	router := regsiterRouter()
	if err := router.Run(":8090"); err != nil {
		common.LogUtil.Errorf("server start failed: %v", err)
	}
}
