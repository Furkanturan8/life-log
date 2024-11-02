package service

import (
	"context"
	"fmt"
	"github.com/bmdavis419/the-better-backend/models"
	"github.com/uptrace/bun"
)

type IBudgetService interface {
	GetAllBudgets(ctx context.Context) ([]models.Budget, error)
	GetBudgetByID(ctx context.Context, id int) (*models.Budget, error)
	CreateBudget(ctx context.Context, budget *models.Budget) error
	UpdateBudget(ctx context.Context, budget *models.Budget) error
	DeleteBudget(ctx context.Context, id int) error
}

type BudgetService struct {
	db *bun.DB
}

func NewBudgetService(db *bun.DB) IBudgetService {
	return &BudgetService{db}
}

func (s BudgetService) GetAllBudgets(ctx context.Context) ([]models.Budget, error) {
	var incomes []models.Budget

	if err := s.db.NewSelect().Model(&incomes).Scan(ctx); err != nil {
		return nil, err
	}
	return incomes, nil
}

func (s BudgetService) GetBudgetByID(ctx context.Context, id int) (*models.Budget, error) {
	var income models.Budget

	if err := s.db.NewSelect().Model(&income).Where("id = ?", id).Scan(ctx); err != nil {
		return nil, err
	}

	return &income, nil
}

func (s BudgetService) CreateBudget(ctx context.Context, income *models.Budget) error {
	_, err := s.db.NewInsert().Model(income).Exec(ctx)
	if err != nil {
		fmt.Println(err.Error())
		return err
	}
	return nil
}

func (s BudgetService) UpdateBudget(ctx context.Context, income *models.Budget) error {
	_, err := s.db.NewUpdate().Model(income).Where("id = ?", income.ID).Exec(ctx)
	if err != nil {
		return err
	}
	return nil
}

func (s BudgetService) DeleteBudget(ctx context.Context, id int) error {
	var income models.Budget

	_, err := s.db.NewDelete().Model(&income).Where("id = ?", id).Exec(ctx)
	if err != nil {
		return err
	}
	return nil
}
