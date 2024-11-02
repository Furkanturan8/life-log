package service

import (
	"context"
	"fmt"
	"github.com/bmdavis419/the-better-backend/models"
	"github.com/uptrace/bun"
)

type IIncomeService interface {
	GetAllIncomes(ctx context.Context) ([]models.Income, error)
	GetIncomeByID(ctx context.Context, id int) (*models.Income, error)
	CreateIncome(ctx context.Context, income *models.Income) error
	UpdateIncome(ctx context.Context, income *models.Income) error
	DeleteIncome(ctx context.Context, id int) error
}

type IncomeService struct {
	db *bun.DB
}

func NewIncomeService(db *bun.DB) IIncomeService {
	return &IncomeService{db}
}

func (s IncomeService) GetAllIncomes(ctx context.Context) ([]models.Income, error) {
	var incomes []models.Income

	if err := s.db.NewSelect().Model(&incomes).Scan(ctx); err != nil {
		return nil, err
	}
	return incomes, nil
}

func (s IncomeService) GetIncomeByID(ctx context.Context, id int) (*models.Income, error) {
	var income models.Income

	if err := s.db.NewSelect().Model(&income).Where("id = ?", id).Scan(ctx); err != nil {
		return nil, err
	}

	return &income, nil
}

func (s IncomeService) CreateIncome(ctx context.Context, income *models.Income) error {
	_, err := s.db.NewInsert().Model(income).Exec(ctx)
	if err != nil {
		fmt.Println(err.Error())
		return err
	}
	return nil
}

func (s IncomeService) UpdateIncome(ctx context.Context, income *models.Income) error {
	_, err := s.db.NewUpdate().Model(income).Where("id = ?", income.ID).Exec(ctx)
	if err != nil {
		return err
	}
	return nil
}

func (s IncomeService) DeleteIncome(ctx context.Context, id int) error {
	var income models.Income

	_, err := s.db.NewDelete().Model(&income).Where("id = ?", id).Exec(ctx)
	if err != nil {
		return err
	}
	return nil
}
