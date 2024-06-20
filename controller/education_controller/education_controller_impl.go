package educationcontroller

import (
	"github.com/faridlan/firestore-go/model/web"
	educationservice "github.com/faridlan/firestore-go/service/education_service"
	"github.com/gofiber/fiber/v2"
)

type EducationControllerImpl struct {
	EducationService educationservice.EducationService
}

func NewEducationController(educationService educationservice.EducationService) EducationController {
	return &EducationControllerImpl{
		EducationService: educationService,
	}
}

func (controller *EducationControllerImpl) Save(ctx *fiber.Ctx) error {

	educationCreate := new(web.EducationCreate)
	err := ctx.BodyParser(educationCreate)
	if err != nil {
		return err
	}

	educationResponse, err := controller.EducationService.Save(ctx.Context(), educationCreate)
	if err != nil {
		return err
	}

	webResponse := web.WebResponse{
		Code:   200,
		Status: "OK",
		Data:   educationResponse,
	}

	return ctx.JSON(webResponse)

}

func (controller *EducationControllerImpl) Update(ctx *fiber.Ctx) error {

	ID := ctx.Params("eduId")

	educationUpdate := new(web.EducationCreate)
	err := ctx.BodyParser(educationUpdate)
	if err != nil {
		return err
	}

	educationUpdate.ID = ID

	educationResponse, err := controller.EducationService.Update(ctx.Context(), educationUpdate)
	if err != nil {
		return err
	}

	webResponse := web.WebResponse{
		Code:   200,
		Status: "OK",
		Data:   educationResponse,
	}

	return ctx.JSON(webResponse)

}

func (controller *EducationControllerImpl) Find(ctx *fiber.Ctx) error {

	educationResponses, err := controller.EducationService.Find(ctx.Context())
	if err != nil {
		return err
	}

	webResponse := web.WebResponse{
		Code:   200,
		Status: "OK",
		Data:   educationResponses,
	}

	return ctx.JSON(webResponse)

}

func (controller *EducationControllerImpl) FindId(ctx *fiber.Ctx) error {

	ID := ctx.Params("eduId")

	educationResponse, err := controller.EducationService.FindId(ctx.Context(), ID)
	if err != nil {
		return err
	}

	webResponse := web.WebResponse{
		Code:   200,
		Status: "OK",
		Data:   educationResponse,
	}

	return ctx.JSON(webResponse)

}

func (controller *EducationControllerImpl) Delete(ctx *fiber.Ctx) error {

	ID := ctx.Params("eduId")

	err := controller.EducationService.Delete(ctx.Context(), ID)
	if err != nil {
		return err
	}

	webResponse := web.WebResponse{
		Code:   200,
		Status: "OK",
	}

	return ctx.JSON(webResponse)

}
