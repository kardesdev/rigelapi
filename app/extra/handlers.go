package extra

import (
	"github.com/gofiber/fiber/v2"
	"github.com/vmkevv/rigelapi/app/models"
)

type ExtraHandler struct {
	app  *fiber.App
	repo ExtraRepository
}

func NewExtraHandler(app *fiber.App, repo ExtraRepository) ExtraHandler {
	return ExtraHandler{app, repo}
}

func (eh *ExtraHandler) handle() {
	eh.app.Get("/years", eh.handleYearsData)
	eh.app.Get("/static", eh.handleStaticData)
}

func (eh *ExtraHandler) handleYearsData(ctx *fiber.Ctx) error {
	years, err := eh.repo.GetYearsData()
	if err != nil {
		return err
	}
	return ctx.JSON(years)
}

func (eh *ExtraHandler) handleStaticData(ctx *fiber.Ctx) error {
	grades, err := eh.repo.GetGrades()
	if err != nil {
		return err
	}
	subjects, err := eh.repo.GetSubjects()
	if err != nil {
		return err
	}
	return ctx.JSON(StaticDataRes{grades, subjects})
}

type StaticDataRes struct {
	Grades   []models.Grade   `json:"grades"`
	Subjects []models.Subject `json:"subjects"`
}