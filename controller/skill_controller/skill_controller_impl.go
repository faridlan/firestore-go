package skillcontroller

import (
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
	panic("not implemented") // TODO: Implement
}

func (controller *SkillControllerImpl) Update(ctx *fiber.Ctx) error {
	panic("not implemented") // TODO: Implement
}

func (controller *SkillControllerImpl) Find(ctx *fiber.Ctx) error {
	panic("not implemented") // TODO: Implement
}

func (controller *SkillControllerImpl) FindId(ctx *fiber.Ctx) error {
	panic("not implemented") // TODO: Implement
}

func (controller *SkillControllerImpl) Delete(ctx *fiber.Ctx) error {
	panic("not implemented") // TODO: Implement
}
