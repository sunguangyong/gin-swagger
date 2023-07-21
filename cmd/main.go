package main

import (
	"github.com/gin-gonic/gin"
	swaggerFiles "github.com/swaggo/files"
	"xunjikeji.com.cn/gin-swagger/ginSwagger"
)

func main() {
	router := gin.Default()
	router.GET("/swagger/v1/*any", ginSwagger.WrapHandler(swaggerFiles.NewHandler(), ginSwagger.InstanceName("v1")))
	_ = router.Run(":8989")
}
