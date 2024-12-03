package main

import (
	"ivandhitya/docker/presenter/handler"

	"github.com/labstack/echo/v4"
)

func main() {
	e := echo.New()
	if err := handler.RoutingRestAPI(e); err != nil {
		e.Logger.Error(err)
	}
	e.Logger.Fatal(e.Start(":8080"))
}
