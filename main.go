package main

import (
	"github.com/labstack/echo/v4"
	"go-example/src/router"
	"net/http"
)

func main() {
	// Echo 인스턴스 생
	e := echo.New()

	// /hello 엔드포인트에 대한 핸들러 함수 등록
	e.GET("/", func(c echo.Context) error {
		// HTTP 응답
		return c.String(http.StatusOK, "Hello, World!")
	})

	// v1 라우팅 그룹 생성
	g := e.Group("/v1")

	router.HelloApi(g)
	router.WorldApi(g)

	// 서버 실행
	e.Logger.Fatal(e.Start(":3333"))
}
