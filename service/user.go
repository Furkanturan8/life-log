package service

import (
	"context"
	"errors"
	"github.com/bmdavis419/the-better-backend/models"
	"github.com/uptrace/bun"
	"log"
)

type IUserService interface {
	Create(context context.Context, user *models.User) error
	GetAll(ctx context.Context) ([]models.User, error)
	GetByID(ctx context.Context, id int) (*models.User, error)
	Delete(ctx context.Context, id int) error
}

type UserService struct {
	DB *bun.DB
}

func NewUserService(db *bun.DB) IUserService {
	return &UserService{DB: db}
}

var ErrUserNotFound = errors.New("no user found")

func (s UserService) Create(context context.Context, user *models.User) error {
	_, err := s.DB.NewInsert().Model(user).Exec(context)
	if err != nil {
		log.Printf("Error inserting user: %v", err) // Hata kaydı
		return err
	}
	return nil
}

func (s UserService) GetAll(ctx context.Context) ([]models.User, error) {
	var users []models.User

	if err := s.DB.NewSelect().Model(&users).Scan(ctx); err != nil {
		return nil, err
	}
	return users, nil
}

func (s UserService) GetByID(ctx context.Context, id int) (*models.User, error) {
	var user models.User

	if err := s.DB.NewSelect().Model(&user).Where("id = ?", id).Scan(ctx); err != nil {
		return nil, err
	}

	return &user, nil
}

func (s UserService) Delete(ctx context.Context, id int) error {
	var user models.User

	result, err := s.DB.NewDelete().Model(&user).Where("id = ?", id).Exec(ctx)
	if err != nil {
		return err
	}

	rowsAffected, _ := result.RowsAffected()
	if rowsAffected == 0 {
		return ErrUserNotFound
	}
	return nil
}
