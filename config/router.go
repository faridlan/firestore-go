package config

import (
	profilecontroller "github.com/faridlan/firestore-go/controller/profile_controller"
	"github.com/faridlan/firestore-go/exception"
	"github.com/gofiber/fiber/v2"
)

type Router struct {
	ProfileContrroler profilecontroller.ProfileController
}

func NewRouter(router Router) *fiber.App {

	app := fiber.New(
		fiber.Config{
			ErrorHandler: exception.ExceptionError,
		},
	)

	app.Post("/api/profiles", router.ProfileContrroler.Save)

	return app

}
