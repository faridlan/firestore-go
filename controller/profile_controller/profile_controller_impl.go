package profilecontroller

import (
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
	panic("not implemented") // TODO: Implement
}

func (controller *ProfileControllerImpl) Find(ctx *fiber.Ctx) error {
	panic("not implemented") // TODO: Implement
}

func (controller *ProfileControllerImpl) Update(ctx *fiber.Ctx) error {
	panic("not implemented") // TODO: Implement
}

func (controller *ProfileControllerImpl) Delete(ctx *fiber.Ctx) error {
	panic("not implemented") // TODO: Implement
}
