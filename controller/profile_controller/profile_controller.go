package profilecontroller

import "github.com/gofiber/fiber/v2"

type ProfileController interface {
	Save(ctx *fiber.Ctx) error
	Find(ctx *fiber.Ctx) error
	Update(ctx *fiber.Ctx) error
	Delete(ctx *fiber.Ctx) error
}
