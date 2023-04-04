package cmd

import (
	"app/internal/controller"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

func Server() {
	// インスタンスを作成
	e := echo.New()
	e.Use(middleware.CORS())

	// ルートを設定
	api := e.Group("/api")
	auth_api := api.Group("")
	gac := controller.GoogleOnlyAuthedController{}
	ac := controller.AuthedController{}
	auth_api.Use(ac.Auth)
	auth_api.Use(gac.Auth)

	auth_api.POST("/user", gac.RegisterUser)
	auth_api.GET("/user_info", gac.GetUserInfo)
	auth_api.POST("/dog", ac.RegisterDog)
	auth_api.POST("/check_in", ac.CheckIn)
	auth_api.GET("/check_in", ac.GetCheckIn)
	auth_api.POST("/check_out", ac.CheckOut)

	//e.GET("/user", controllers.GetUser)

	// public_api := e.Group("/public")

	// サーバーをポート番号1323で起動
	e.Logger.Fatal(e.Start(":1323"))
}
