package router

import (
	"github.com/bmdavis419/the-better-backend/handlers"
	"github.com/bmdavis419/the-better-backend/pkg/database"
	"github.com/bmdavis419/the-better-backend/service"
	"github.com/gofiber/fiber/v2"
)

func SetupRoutes(app *fiber.App) {
	db := database.GetDB()

	userService := service.NewUserService(db)
	userHandler := handlers.NewUserHandler(userService)

	incomeService := service.NewIncomeService(db)
	incomeHandler := handlers.NewIncomeHandler(incomeService)

	expenseService := service.NewExpenseService(db)
	expenseHandler := handlers.NewExpenseHandler(expenseService)

	budgetService := service.NewBudgetService(db)
	budgetHandler := handlers.NewBudgetHandler(budgetService)

	reportService := service.NewReportService(db)
	reportHandler := handlers.NewReportHandler(reportService)

	// User routes
	app.Get("api/users", userHandler.GetAllUser)
	app.Post("api/user", userHandler.CreateUser)
	app.Get("api/user/:id", userHandler.GetUserByID)
	app.Delete("api/user/:id", userHandler.Delete)

	// Income routes
	app.Get("api/incomes", incomeHandler.GetAllIncome)
	app.Post("api/income", incomeHandler.CreateIncome)
	app.Get("api/income/:id", incomeHandler.GetIncomeByID)
	app.Put("api/income/:id", incomeHandler.UpdateIncome)
	app.Delete("api/income/:id", incomeHandler.DeleteIncome)

	// Expense routes
	app.Get("api/expenses", expenseHandler.GetAllExpense)
	app.Post("api/expense", expenseHandler.CreateExpense)
	app.Get("api/expense/:id", expenseHandler.GetExpenseByID)
	app.Put("api/expense/:id", expenseHandler.UpdateExpense)
	app.Delete("api/expense/:id", expenseHandler.DeleteExpense)

	// Budget routes
	app.Get("api/budgets", budgetHandler.GetAllBudget)
	app.Post("api/budget", budgetHandler.CreateBudget)
	app.Get("api/budget/:id", budgetHandler.GetBudgetByID)
	app.Put("api/budget/:id", budgetHandler.UpdateBudget)
	app.Delete("api/budget/:id", budgetHandler.DeleteBudget)

	// Report routes
	app.Post("api/report", reportHandler.GenerateReport)
}
