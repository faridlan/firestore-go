package experiencecontroller

import "github.com/gofiber/fiber/v2"

type ExperienceController interface {
	Save(ctx *fiber.Ctx) error
	Update(ctx *fiber.Ctx) error
	FindId(ctx *fiber.Ctx) error
	Find(ctx *fiber.Ctx) error
	Delete(ctx *fiber.Ctx) error
}
