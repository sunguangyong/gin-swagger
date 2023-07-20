package main

import (
	"github.com/gin-gonic/gin"
	swaggerFiles "github.com/swaggo/files"
	"xunjikeji.com.cn/gin-swagger/ginSwagger"
	//ginSwagger "xunjikeji.com.cn/gin-swagger"
)

func main() {

	// New gin router
	router := gin.Default()

	//swag.Register(docs.SwaggerInfov1.InstanceName(), docs.GetSwag(path))
	router.GET("/swagger/v1/*any", ginSwagger.WrapHandler(swaggerFiles.NewHandler(), ginSwagger.InstanceName("v1")))

	_ = router.Run()

}
