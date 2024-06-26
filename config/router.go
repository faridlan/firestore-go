package config

import (
	educationcontroller "github.com/faridlan/firestore-go/controller/education_controller"
	experiencecontroller "github.com/faridlan/firestore-go/controller/experience_controller"
	profilecontroller "github.com/faridlan/firestore-go/controller/profile_controller"
	skillcontroller "github.com/faridlan/firestore-go/controller/skill_controller"
	"github.com/faridlan/firestore-go/exception"
	"github.com/gofiber/fiber/v2"
)

type Router struct {
	ProfileContrroler    profilecontroller.ProfileController
	ExperienceContrroler experiencecontroller.ExperienceController
	EducationController  educationcontroller.EducationController
	SkillController      skillcontroller.SkillController
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

	app.Post("/api/educations", router.EducationController.Save)
	app.Get("/api/educations/:eduId", router.EducationController.FindId)
	app.Get("/api/educations/", router.EducationController.Find)
	app.Delete("/api/educations/:eduId", router.EducationController.Delete)
	app.Put("/api/educations/:eduId", router.EducationController.Update)

	app.Post("/api/skills", router.SkillController.Save)
	app.Get("/api/skills/:skillId", router.SkillController.FindId)
	app.Get("/api/skills/", router.SkillController.Find)
	app.Delete("/api/skills/:skillId", router.SkillController.Delete)
	app.Put("/api/skills/:skillId", router.SkillController.Update)

	return app

}
