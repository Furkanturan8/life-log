package handlers

import (
	"github.com/bmdavis419/the-better-backend/models"
	"github.com/bmdavis419/the-better-backend/service"
	"github.com/gofiber/fiber/v2"
	"strconv"
)

type BudgetHandler struct {
	budgetService service.IBudgetService
}

func NewBudgetHandler(s service.IBudgetService) BudgetHandler {
	return BudgetHandler{budgetService: s}
}

func (h BudgetHandler) GetAllBudget(ctx *fiber.Ctx) error {
	budgets, err := h.budgetService.GetAllBudgets(ctx.Context())
	if err != nil {
		return ctx.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": "Error fetching budgets"})
	}
	return ctx.Status(fiber.StatusOK).JSON(budgets)
}

func (h BudgetHandler) GetBudgetByID(ctx *fiber.Ctx) error {
	param := ctx.Params("id")
	id, err := strconv.Atoi(param)
	if err != nil {
		return ctx.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "Invalid input"})
	}

	budget, err := h.budgetService.GetBudgetByID(ctx.Context(), id)
	if err != nil {
		return ctx.Status(fiber.StatusNotFound).JSON(fiber.Map{"error": "Budget not found"})
	}
	return ctx.Status(fiber.StatusOK).JSON(budget)
}

func (h BudgetHandler) CreateBudget(ctx *fiber.Ctx) error {
	var budget models.Budget

	if err := ctx.BodyParser(&budget); err != nil {
		return ctx.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "Invalid input"})
	}

	if err := h.budgetService.CreateBudget(ctx.Context(), &budget); err != nil {
		return ctx.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": "Error creating budget"})
	}
	return ctx.Status(fiber.StatusOK).JSON(fiber.Map{"message": "budget created successfully!"})
}

func (h BudgetHandler) UpdateBudget(ctx *fiber.Ctx) error {
	var budget models.Budget

	if err := ctx.BodyParser(&budget); err != nil {
		return ctx.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "Invalid input"})
	}

	param := ctx.Params("id")
	id, err := strconv.Atoi(param)
	if err != nil {
		return ctx.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "Invalid input"})
	}
	budget.ID = id
	if err = h.budgetService.UpdateBudget(ctx.Context(), &budget); err != nil {
		return ctx.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": "Error updating budget"})
	}
	return ctx.Status(fiber.StatusOK).JSON(fiber.Map{"message": "Budget updated successfully!"})
}

func (h BudgetHandler) DeleteBudget(ctx *fiber.Ctx) error {
	param := ctx.Params("id")
	id, err := strconv.Atoi(param)
	if err != nil {
		return ctx.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "Invalid input"})
	}

	if err = h.budgetService.DeleteBudget(ctx.Context(), id); err != nil {
		return ctx.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": "Error deleting budget"})
	}
	return ctx.Status(fiber.StatusOK).JSON(fiber.Map{"message": "Budget deleted successfully!"})
}
