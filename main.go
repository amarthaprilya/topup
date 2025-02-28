package main

import (
	_ "camera-rent/docs"
	"camera-rent/handler"
)

func main() {
	// @title Sweager Midtrans API
	// @description Sweager service API in Go using Gin framework
	// @host https://rental-camera-6a4d1d4fd45f.herokuapp.com
	// @securitydefinitions.apikey BearerAuth
	// @in header
	// @name Authorization
	handler.StartApp()
}
