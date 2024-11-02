package service

import (
	"context"
	"fmt"
	"github.com/bmdavis419/the-better-backend/models"
	"github.com/uptrace/bun"
	"time"
)

type IReportService interface {
	GenerateReport(ctx context.Context, userID int, month int, year int) (*models.Report, error)
}

type ReportService struct {
	db *bun.DB
}

func NewReportService(db *bun.DB) IReportService {
	return &ReportService{db}
}

func (s ReportService) GenerateReport(ctx context.Context, userID int, month int, year int) (*models.Report, error) {
	var report models.Report
	report.UserID = userID
	report.Month = month
	report.Year = year

	user := models.User{}
	err := s.db.NewSelect().Model(&user).Where("id = ?", report.UserID).Scan(ctx)

	if err != nil {
		return nil, fmt.Errorf("failed to retrieve user: %w", err)
	}
	report.User = &user

	// Verilen ay ve yıl için tüm gelirleri alın
	var incomes []models.Income
	err = s.db.NewSelect().Model(&incomes).Where("user_id = ? AND EXTRACT(MONTH FROM date) = ? AND EXTRACT(YEAR FROM date) = ?", userID, month, year).Scan(ctx)
	if err != nil {
		return nil, err
	}

	//Verilen ay ve yıl için tüm harcamaları alın
	var expenses []models.Expense
	err = s.db.NewSelect().Model(&expenses).Where("user_id = ? AND EXTRACT(MONTH FROM date) = ? AND EXTRACT(YEAR FROM date) = ?", userID, month, year).Scan(ctx)
	if err != nil {
		return nil, err
	}

	// toplam geliri hesapla
	var totalIncome float64
	for _, income := range incomes {
		totalIncome += income.Amount
	}
	report.TotalIncome = totalIncome

	// toplam masraf/gideri hesapla
	var totalExpense float64
	for _, expense := range expenses {
		totalExpense += expense.Amount
	}
	report.TotalExpense = totalExpense

	// bakiyeyi hesapla
	report.Balance = totalIncome - totalExpense
	report.GeneratedAt = time.Now()

	_, err = s.db.NewInsert().Model(&report).Exec(ctx)
	if err != nil {
		return nil, err
	}

	return &report, nil
}
