package experiencecontroller

import (
	"github.com/faridlan/firestore-go/model/web"
	experienceservice "github.com/faridlan/firestore-go/service/experience_service"
	"github.com/gofiber/fiber/v2"
)

type ExperienceControllerImpl struct {
	ExperienceService experienceservice.ExperienceService
}

func NewExperienceController(experienceService experienceservice.ExperienceService) ExperienceController {
	return &ExperienceControllerImpl{
		ExperienceService: experienceService,
	}
}

func (controller *ExperienceControllerImpl) Save(ctx *fiber.Ctx) error {

	experienceCreate := new(web.ExperienceCreate)
	err := ctx.BodyParser(experienceCreate)
	if err != nil {
		return err
	}

	experienceResponse, err := controller.ExperienceService.Save(ctx.Context(), experienceCreate)
	if err != nil {
		return err
	}

	webResponse := web.WebResponse{
		Code:   200,
		Status: "OK",
		Data:   experienceResponse,
	}

	return ctx.JSON(webResponse)

}

func (controller *ExperienceControllerImpl) Update(ctx *fiber.Ctx) error {

	id := ctx.Params("experienceId")

	experienceCreate := new(web.ExperienceCreate)
	err := ctx.BodyParser(experienceCreate)
	if err != nil {
		return err
	}

	experienceCreate.ID = id

	experienceResponse, err := controller.ExperienceService.Update(ctx.Context(), experienceCreate)
	if err != nil {
		return err
	}

	webResponse := web.WebResponse{
		Code:   200,
		Status: "OK",
		Data:   experienceResponse,
	}

	return ctx.JSON(webResponse)

}

func (controller *ExperienceControllerImpl) FindId(ctx *fiber.Ctx) error {

	id := ctx.Params("experienceId")

	experienceResponse, err := controller.ExperienceService.FindId(ctx.Context(), id)
	if err != nil {
		return err
	}

	webResponse := web.WebResponse{
		Code:   200,
		Status: "OK",
		Data:   experienceResponse,
	}

	return ctx.JSON(webResponse)

}

func (controller *ExperienceControllerImpl) Find(ctx *fiber.Ctx) error {

	experienceResponse, err := controller.ExperienceService.Find(ctx.Context())
	if err != nil {
		return err
	}

	webResponse := web.WebResponse{
		Code:   200,
		Status: "OK",
		Data:   experienceResponse,
	}

	return ctx.JSON(webResponse)

}

func (controller *ExperienceControllerImpl) Delete(ctx *fiber.Ctx) error {

	id := ctx.Params("experienceId")

	err := controller.ExperienceService.Delete(ctx.Context(), id)
	if err != nil {
		return err
	}

	webResponse := web.WebResponse{
		Code:   200,
		Status: "OK",
	}

	return ctx.JSON(webResponse)

}
