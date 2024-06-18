package config

import (
	experiencecontroller "github.com/faridlan/firestore-go/controller/experience_controller"
	profilecontroller "github.com/faridlan/firestore-go/controller/profile_controller"
	"github.com/faridlan/firestore-go/exception"
	"github.com/gofiber/fiber/v2"
)

type Router struct {
	ProfileContrroler    profilecontroller.ProfileController
	ExperienceContrroler experiencecontroller.ExperienceController
}

func NewRouter(router Router) *fiber.App {

	app := fiber.New(
		fiber.Config{
			ErrorHandler: exception.ExceptionError,
		},
	)

	app.Post("/api/profiles", router.ProfileContrroler.Save)
	app.Get("/api/profiles/:profileId", router.ProfileContrroler.Find)
	app.Delete("/api/profiles/:profileId", router.ProfileContrroler.Delete)
	app.Put("/api/profiles/:profileId", router.ProfileContrroler.Update)

	app.Post("/api/experience", router.ExperienceContrroler.Save)
	app.Get("/api/experience/:experienceId", router.ExperienceContrroler.FindId)
	app.Get("/api/experience/", router.ExperienceContrroler.Find)
	app.Delete("/api/experience/:experienceId", router.ExperienceContrroler.Delete)
	app.Put("/api/experience/:experienceId", router.ExperienceContrroler.Update)

	return app

}
