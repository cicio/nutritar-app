package repository

import (
	"backend/internal/models"
	"database/sql"
)

type DatabaseRepo interface {
	AllFoods() ([]*models.Food, error)
	Connection() *sql.DB
	GetUserByEmail(email string) (*models.User, error)
	GetUserByID(id int) (*models.User, error)
}