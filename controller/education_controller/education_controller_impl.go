package educationcontroller

import (
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
	panic("not implemented") // TODO: Implement
}

func (controller *EducationControllerImpl) Update(ctx *fiber.Ctx) error {
	panic("not implemented") // TODO: Implement
}

func (controller *EducationControllerImpl) Find(ctx *fiber.Ctx) error {
	panic("not implemented") // TODO: Implement
}

func (controller *EducationControllerImpl) FindId(ctx *fiber.Ctx) error {
	panic("not implemented") // TODO: Implement
}

func (controller *EducationControllerImpl) Delete(ctx *fiber.Ctx) error {
	panic("not implemented") // TODO: Implement
}
