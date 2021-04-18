package route

import (
	"github.com/gofiber/fiber/v2"
	"github.com/warkop/up-meetup-clone/api/v1/user/gateway/handler/http"
)

type UserRouteProto interface {
	Init(app *fiber.App)
}

type UserRoute struct {
	Handler *http.UserHandler
}

func NewUserRoute(handler http.UserHandlerProto) UserRouteProto {
	return &UserRoute{
		Handler: handler.(*http.UserHandler),
	}
}

func (ur *UserRoute) Init(app *fiber.App) {
	user := app.Group("/api/v1")

	user.Get("/user", ur.Handler.Fetch)
	user.Get("/user/:id", ur.Handler.FetchByID)
	user.Post("/user", ur.Handler.Create)
	user.Put("/user/:id", ur.Handler.Update)
	user.Delete("/user/:id", ur.Handler.Delete)
}
