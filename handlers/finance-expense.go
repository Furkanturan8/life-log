package handlers

import (
	"github.com/bmdavis419/the-better-backend/models"
	"github.com/bmdavis419/the-better-backend/service"
	"github.com/gofiber/fiber/v2"
	"strconv"
)

type ExpenseHandler struct {
	expenseService service.IExpenseService
}

func NewExpenseHandler(s service.IExpenseService) ExpenseHandler {
	return ExpenseHandler{expenseService: s}
}

func (h ExpenseHandler) GetAllExpense(ctx *fiber.Ctx) error {
	expenses, err := h.expenseService.GetAllExpenses(ctx.Context())
	if err != nil {
		return ctx.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": "Error fetching expenses"})
	}
	return ctx.Status(fiber.StatusOK).JSON(expenses)
}

func (h ExpenseHandler) GetExpenseByID(ctx *fiber.Ctx) error {
	param := ctx.Params("id")
	id, err := strconv.Atoi(param)
	if err != nil {
		return ctx.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "Invalid input"})
	}

	expense, err := h.expenseService.GetExpenseByID(ctx.Context(), id)
	if err != nil {
		return ctx.Status(fiber.StatusNotFound).JSON(fiber.Map{"error": "Expense not found"})
	}
	return ctx.Status(fiber.StatusOK).JSON(expense)
}

func (h ExpenseHandler) CreateExpense(ctx *fiber.Ctx) error {
	var expense models.Expense

	if err := ctx.BodyParser(&expense); err != nil {
		return ctx.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "Invalid input"})
	}

	if err := h.expenseService.CreateExpense(ctx.Context(), &expense); err != nil {
		return ctx.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": "Error creating expense"})
	}
	return ctx.Status(fiber.StatusOK).JSON(fiber.Map{"message": "expense created successfully!"})
}

func (h ExpenseHandler) UpdateExpense(ctx *fiber.Ctx) error {
	var expense models.Expense

	if err := ctx.BodyParser(&expense); err != nil {
		return ctx.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "Invalid input"})
	}

	param := ctx.Params("id")
	id, err := strconv.Atoi(param)
	if err != nil {
		return ctx.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "Invalid input"})
	}
	expense.ID = id
	if err = h.expenseService.UpdateExpense(ctx.Context(), &expense); err != nil {
		return ctx.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": "Error updating expense"})
	}
	return ctx.Status(fiber.StatusOK).JSON(fiber.Map{"message": "Expense updated successfully!"})
}

func (h ExpenseHandler) DeleteExpense(ctx *fiber.Ctx) error {
	param := ctx.Params("id")
	id, err := strconv.Atoi(param)
	if err != nil {
		return ctx.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "Invalid input"})
	}

	if err = h.expenseService.DeleteExpense(ctx.Context(), id); err != nil {
		return ctx.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": "Error deleting expense"})
	}
	return ctx.Status(fiber.StatusOK).JSON(fiber.Map{"message": "Expense deleted successfully!"})
}
