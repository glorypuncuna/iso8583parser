package main

import (
	"mappertest/handler"
	"mappertest/service"

	"github.com/gin-gonic/gin"
)

func main() {

	bitmapBitService := service.NewBitmapBitService()
	mtiService := service.NewMTIService()
	isoService := service.NewISOService()

	bitmapBitHandler := handler.NewBitmapBitHandler(bitmapBitService)
	mtiHandler := handler.NewMTIHandler(mtiService)
	isoHandler := handler.NewISOHandler(isoService)

	router := gin.Default()
	api := router.Group("/v1/")

	api.POST("/bitmapBit/:input", bitmapBitHandler.ResultField)
	api.POST("/mti/:input", mtiHandler.MTIConverter)
	api.POST("/iso/coba/", isoHandler.TestDulu)

	router.Run()
}
