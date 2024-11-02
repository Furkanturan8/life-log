package handlers

import (
	"errors"
	"github.com/bmdavis419/the-better-backend/models"
	"github.com/bmdavis419/the-better-backend/service"
	"github.com/gofiber/fiber/v2"

	"strconv"
)

type UserHandler struct {
	userService service.IUserService
}

func NewUserHandler(s service.IUserService) UserHandler {
	return UserHandler{userService: s}
}

func (h UserHandler) CreateUser(ctx *fiber.Ctx) error {
	var user models.User

	if err := ctx.BodyParser(&user); err != nil {
		return ctx.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "Invalid input"})
	}

	if err := h.userService.Create(ctx.Context(), &user); err != nil {
		return ctx.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": "Error creating user"})
	}
	return ctx.Status(fiber.StatusOK).JSON(fiber.Map{"message": "User created successfully!"})
}

func (h UserHandler) GetAllUser(ctx *fiber.Ctx) error {
	users, err := h.userService.GetAll(ctx.Context())
	if err != nil {
		return ctx.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": "Error fetching users"})
	}
	return ctx.Status(fiber.StatusOK).JSON(users)
}

func (h UserHandler) GetUserByID(ctx *fiber.Ctx) error {
	param := ctx.Params("id")
	id, err := strconv.Atoi(param)
	if err != nil {
		return ctx.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "Invalid input"})
	}

	user, err := h.userService.GetByID(ctx.Context(), id)
	if err != nil {
		return ctx.Status(fiber.StatusNotFound).JSON(fiber.Map{"error": "User not found"})
	}

	return ctx.Status(fiber.StatusOK).JSON(user)
}

func (h UserHandler) Delete(ctx *fiber.Ctx) error {
	id := ctx.Params("id")
	userID, err := strconv.Atoi(id)
	if err != nil {
		return ctx.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "Invalid input"})
	}

	if err = h.userService.Delete(ctx.Context(), userID); err != nil {
		if errors.Is(err, service.ErrUserNotFound) {
			return ctx.Status(fiber.StatusNotFound).JSON(fiber.Map{"error": "User not found!"})
		}
		return ctx.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": "Error deleting user"})
	}
	return ctx.Status(fiber.StatusOK).JSON(fiber.Map{"message": "User deleted successfully!"})
}
