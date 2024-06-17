package profilecontroller

import (
	"github.com/faridlan/firestore-go/model/web"
	profileservice "github.com/faridlan/firestore-go/service/profile_service"
	"github.com/gofiber/fiber/v2"
)

type ProfileControllerImpl struct {
	ProfileService profileservice.ProfileService
}

func NewProfileController(profileService profileservice.ProfileService) ProfileController {
	return &ProfileControllerImpl{
		ProfileService: profileService,
	}
}

func (controller *ProfileControllerImpl) Save(ctx *fiber.Ctx) error {

	profileCreate := new(web.Profile)
	err := ctx.BodyParser(profileCreate)
	if err != nil {
		return err
	}

	profileResponse, err := controller.ProfileService.Save(ctx.Context(), profileCreate)
	if err != nil {
		return err
	}

	webRespons := web.WebResponse{
		Code:   200,
		Status: "OK",
		Data:   profileResponse,
	}

	return ctx.JSON(webRespons)

}

func (controller *ProfileControllerImpl) Find(ctx *fiber.Ctx) error {

	profileid := ctx.Params("profileId")

	profileResposne, err := controller.ProfileService.Find(ctx.Context(), profileid)
	if err != nil {
		return err
	}

	webResponse := web.WebResponse{
		Code:   fiber.StatusOK,
		Status: "OK",
		Data:   profileResposne,
	}

	return ctx.JSON(webResponse)

}

func (controller *ProfileControllerImpl) Update(ctx *fiber.Ctx) error {

	profileid := ctx.Params("profileId")

	profileUpdate := new(web.Profile)
	err := ctx.BodyParser(profileUpdate)
	if err != nil {
		return err
	}

	profileUpdate.ID = profileid

	profileResposne, err := controller.ProfileService.Update(ctx.Context(), profileUpdate)
	if err != nil {
		return err
	}

	webResponse := web.WebResponse{
		Code:   fiber.StatusOK,
		Status: "OK",
		Data:   profileResposne,
	}

	return ctx.JSON(webResponse)

}

func (controller *ProfileControllerImpl) Delete(ctx *fiber.Ctx) error {

	profileid := ctx.Params("profileId")

	err := controller.ProfileService.Delete(ctx.Context(), profileid)
	if err != nil {
		return err
	}

	webResponse := web.WebResponse{
		Code:   fiber.StatusOK,
		Status: "OK",
	}

	return ctx.JSON(webResponse)

}
