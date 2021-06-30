package main

import (
	"fmt"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

func corsConfig() middleware.CORSConfig {
	cc := middleware.CORSConfig{
		AllowHeaders:     []string{echo.HeaderAccept, echo.HeaderAcceptEncoding, echo.HeaderContentLength, echo.HeaderContentType, echo.HeaderOrigin},
		AllowCredentials: true,
		ExposeHeaders:    []string{echo.HeaderAccept, echo.HeaderAcceptEncoding, echo.HeaderContentLength, echo.HeaderContentType, echo.HeaderOrigin},
		AllowOrigins:     []string{"*"},
	}

	return cc
}

func webserviceStart() {
	fmt.Println("starting webservices ...")

	e := echo.New()

	//e.Use(middleware.CORSWithConfig(corsConfig()))

	r := e.Group("/earthquake")

	r.GET("/", func(c echo.Context) error {
		return earthquakeList(c, db)
	})

	r.GET("/:id/", func(c echo.Context) error {
		return earthquakeByID(c, db)
	})

	e.Logger.Fatal(e.Start(":4000"))
}
