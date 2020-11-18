package main

import (
	"fmt"

	openapi "github.com/birnamwood/go-open-api/openapi/gen"
	"github.com/labstack/echo/v4"
)

func main() {
	e := echo.New()
	han := handle.userHandler{}
	// 定義した struct を登録
	openapi.RegisterHandlers(e, han)
	e.Logger.Fatal(e.Start(fmt.Sprintf("0.0.0.0:%d", 8000)))
}
