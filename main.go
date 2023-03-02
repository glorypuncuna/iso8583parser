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
	//Parse with all the plain text
	api.POST("/iso/mode1/", isoHandler.Mode1)
	//Parse the hexdump
	api.POST("/iso/mode2/", isoHandler.Mode2)
	api.POST("/iso/mode3/", isoHandler.Mode3)

	router.Run()
}
