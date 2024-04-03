package router

import (
	"github.com/labstack/echo/v4"
	"go-example/src/service"
)

func HelloApi(group *echo.Group) {
	// hellos 라우팅
	g := group.Group("/hellos")

	// 라우팅에 대한 핸들러 함수 호출
	g.GET("", service.Hello)
}
