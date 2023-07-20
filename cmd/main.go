package main

import (
	"github.com/gin-gonic/gin"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
	v1 "github.com/swaggo/gin-swagger/example/multiple/api/v1"
	_ "github.com/swaggo/gin-swagger/example/multiple/docs"
)

func main() {

	// New gin router
	router := gin.Default()

	// Register api/v1 endpoints
	v1.Register(router)
	//swag.Register(docs.SwaggerInfov1.InstanceName(), docs.GetSwag(path))
	router.GET("/swagger/v1/*any", ginSwagger.WrapHandler(swaggerFiles.NewHandler(), ginSwagger.InstanceName("v1")))

	// Register api/v2 endpoints
	//v2.Register(router)
	//router.GET("/swagger/v2/*any", ginSwagger.WrapHandler(swaggerFiles.NewHandler(), ginSwagger.InstanceName("v2")))
	// Listen and Server in
	_ = router.Run()

}
