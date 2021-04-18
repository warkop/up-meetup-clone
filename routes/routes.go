package routes

import (
	"github.com/gofiber/fiber/v2"
	user_v1 "github.com/warkop/up-meetup-clone/api/v1/user"
	user_route_v1 "github.com/warkop/up-meetup-clone/api/v1/user/gateway/route"
	"gorm.io/gorm"
)

func Endpoints(app *fiber.App) {
	user_route_v1.NewUserRoute(user_v1.ProvideUserHandler(&gorm.DB{})).Init(app)
}
