package service

import (
	"context"
	"fmt"
	"github.com/bmdavis419/the-better-backend/models"
	"github.com/uptrace/bun"
)

type IExpenseService interface {
	GetAllExpenses(ctx context.Context) ([]models.Expense, error)
	GetExpenseByID(ctx context.Context, id int) (*models.Expense, error)
	CreateExpense(ctx context.Context, expense *models.Expense) error
	UpdateExpense(ctx context.Context, expense *models.Expense) error
	DeleteExpense(ctx context.Context, id int) error
}

type ExpenseService struct {
	db *bun.DB
}

func NewExpenseService(db *bun.DB) IExpenseService {
	return &ExpenseService{db}
}

func (s ExpenseService) GetAllExpenses(ctx context.Context) ([]models.Expense, error) {
	var incomes []models.Expense

	if err := s.db.NewSelect().Model(&incomes).Scan(ctx); err != nil {
		return nil, err
	}
	return incomes, nil
}

func (s ExpenseService) GetExpenseByID(ctx context.Context, id int) (*models.Expense, error) {
	var income models.Expense

	if err := s.db.NewSelect().Model(&income).Where("id = ?", id).Scan(ctx); err != nil {
		return nil, err
	}

	return &income, nil
}

func (s ExpenseService) CreateExpense(ctx context.Context, income *models.Expense) error {
	_, err := s.db.NewInsert().Model(income).Exec(ctx)
	if err != nil {
		fmt.Println(err.Error())
		return err
	}
	return nil
}

func (s ExpenseService) UpdateExpense(ctx context.Context, income *models.Expense) error {
	_, err := s.db.NewUpdate().Model(income).Where("id = ?", income.ID).Exec(ctx)
	if err != nil {
		return err
	}
	return nil
}

func (s ExpenseService) DeleteExpense(ctx context.Context, id int) error {
	var income models.Expense

	_, err := s.db.NewDelete().Model(&income).Where("id = ?", id).Exec(ctx)
	if err != nil {
		return err
	}
	return nil
}
