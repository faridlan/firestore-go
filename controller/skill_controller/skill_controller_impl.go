package skillcontroller

import (
	"github.com/faridlan/firestore-go/model/web"
	skillservice "github.com/faridlan/firestore-go/service/skill_service"
	"github.com/gofiber/fiber/v2"
)

type SkillControllerImpl struct {
	SkillService skillservice.SkillService
}

func NewSkillController(skillService skillservice.SkillService) SkillController {
	return &SkillControllerImpl{
		SkillService: skillService,
	}
}

func (controller *SkillControllerImpl) Save(ctx *fiber.Ctx) error {

	skillCreate := new(web.SkillCreate)
	err := ctx.BodyParser(skillCreate)
	if err != nil {
		return err
	}

	skillResponse, err := controller.SkillService.Save(ctx.Context(), skillCreate)
	if err != nil {
		return err
	}

	webREsponse := web.WebResponse{
		Code:   200,
		Status: "OK",
		Data:   skillResponse,
	}

	return ctx.JSON(webREsponse)

}

func (controller *SkillControllerImpl) Update(ctx *fiber.Ctx) error {

	ID := ctx.Params("skillId")

	skillUpdate := new(web.SkillCreate)
	err := ctx.BodyParser(skillUpdate)
	if err != nil {
		return err
	}

	skillUpdate.ID = ID

	skillResponse, err := controller.SkillService.Update(ctx.Context(), skillUpdate)
	if err != nil {
		return err
	}

	webREsponse := web.WebResponse{
		Code:   200,
		Status: "OK",
		Data:   skillResponse,
	}

	return ctx.JSON(webREsponse)

}

func (controller *SkillControllerImpl) Find(ctx *fiber.Ctx) error {

	skillResponse, err := controller.SkillService.Find(ctx.Context())
	if err != nil {
		return err
	}

	webREsponse := web.WebResponse{
		Code:   200,
		Status: "OK",
		Data:   skillResponse,
	}

	return ctx.JSON(webREsponse)

}

func (controller *SkillControllerImpl) FindId(ctx *fiber.Ctx) error {

	ID := ctx.Params("skillId")

	skillResponse, err := controller.SkillService.FindId(ctx.Context(), ID)
	if err != nil {
		return err
	}

	webREsponse := web.WebResponse{
		Code:   200,
		Status: "OK",
		Data:   skillResponse,
	}

	return ctx.JSON(webREsponse)

}

func (controller *SkillControllerImpl) Delete(ctx *fiber.Ctx) error {

	ID := ctx.Params("skillId")

	err := controller.SkillService.Delete(ctx.Context(), ID)
	if err != nil {
		return err
	}

	webREsponse := web.WebResponse{
		Code:   200,
		Status: "OK",
	}

	return ctx.JSON(webREsponse)

}
