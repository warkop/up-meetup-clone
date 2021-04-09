package route

import (
	"github.com/labstack/echo/v4"
	"github.com/warkop/up-meetup-clone/api/v1/user/gateway/handler/http"
)

type UserRouteProto interface {
	Init(e *echo.Echo)
}

type UserRoute struct {
	Handler *http.UserHandler
}

func NewUserRoute(handler http.UserHandlerProto) UserRouteProto {
	return &UserRoute{
		Handler: handler.(*http.UserHandler),
	}
}

func (ur *UserRoute) Init(e *echo.Echo) {
	user := e.Group("/api/v1")

	user.GET("/user", ur.Handler.Fetch)
	user.GET("/user/:id", ur.Handler.FetchByID)
	user.POST("/user", ur.Handler.Create)
	user.PUT("/user/:id", ur.Handler.Update)
	user.DELETE("/user/:id", ur.Handler.Delete)
}
