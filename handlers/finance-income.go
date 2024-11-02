package handlers

import (
	"github.com/bmdavis419/the-better-backend/models"
	"github.com/bmdavis419/the-better-backend/service"
	"github.com/gofiber/fiber/v2"
	"strconv"
)

type IncomeHandler struct {
	incomeService service.IIncomeService
}

func NewIncomeHandler(s service.IIncomeService) IncomeHandler {
	return IncomeHandler{incomeService: s}
}

func (h IncomeHandler) GetAllIncome(ctx *fiber.Ctx) error {
	incomes, err := h.incomeService.GetAllIncomes(ctx.Context())
	if err != nil {
		return ctx.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": "Error fetching incomes"})
	}
	return ctx.Status(fiber.StatusOK).JSON(incomes)
}

func (h IncomeHandler) GetIncomeByID(ctx *fiber.Ctx) error {
	param := ctx.Params("id")
	id, err := strconv.Atoi(param)
	if err != nil {
		return ctx.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "Invalid input"})
	}

	income, err := h.incomeService.GetIncomeByID(ctx.Context(), id)
	if err != nil {
		return ctx.Status(fiber.StatusNotFound).JSON(fiber.Map{"error": "Income not found"})
	}
	return ctx.Status(fiber.StatusOK).JSON(income)
}

func (h IncomeHandler) CreateIncome(ctx *fiber.Ctx) error {
	var income models.Income

	if err := ctx.BodyParser(&income); err != nil {
		return ctx.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "Invalid input"})
	}

	if err := h.incomeService.CreateIncome(ctx.Context(), &income); err != nil {
		return ctx.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": "Error creating income"})
	}
	return ctx.Status(fiber.StatusOK).JSON(fiber.Map{"message": "Income created successfully!"})
}

func (h IncomeHandler) UpdateIncome(ctx *fiber.Ctx) error {
	var income models.Income

	if err := ctx.BodyParser(&income); err != nil {
		return ctx.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "Invalid input"})
	}

	param := ctx.Params("id")
	id, err := strconv.Atoi(param)
	if err != nil {
		return ctx.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "Invalid input"})
	}
	income.ID = id
	if err = h.incomeService.UpdateIncome(ctx.Context(), &income); err != nil {
		return ctx.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": "Error updating income"})
	}
	return ctx.Status(fiber.StatusOK).JSON(fiber.Map{"message": "Income updated successfully!"})
}

func (h IncomeHandler) DeleteIncome(ctx *fiber.Ctx) error {
	param := ctx.Params("id")
	id, err := strconv.Atoi(param)
	if err != nil {
		return ctx.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "Invalid input"})
	}

	if err := h.incomeService.DeleteIncome(ctx.Context(), id); err != nil {
		return ctx.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": "Error deleting income"})
	}
	return ctx.Status(fiber.StatusOK).JSON(fiber.Map{"message": "Income deleted successfully!"})
}
