package main

import (
	"github.com/labstack/echo/v4"
	echoSwagger "github.com/swaggo/echo-swagger"
	_ "go-example/docs" // swag cli에 의해 생성되므로 가져와야 합니다.
	"go-example/src/router"
	"net/http"
)

// @title							go-example
// @version							1.0
// @description						echo-swagger 통한 API 문서화
// @BasePath						/v1
func main() {
	// Echo 인스턴스 생
	e := echo.New()

	swagDocList := echoSwagger.DocExpansion("none") // // API 문서 목록을 펼쳐 보여주지 않는 옵션

	e.GET("/v1/*", echoSwagger.EchoWrapHandler(swagDocList))

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
