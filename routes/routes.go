package routes

import (
	"github.com/labstack/echo/v4"
	user_v1 "github.com/warkop/up-meetup-clone/api/v1/user"
	user_route_v1 "github.com/warkop/up-meetup-clone/api/v1/user/gateway/route"
)

func Endpoints(e *echo.Echo) {
	user_route_v1.NewUserRoute(user_v1.ProvideUserHandler()).Init(e)
}
