package repository

import (
	"example.com/project/pkg/models"
	"database/sql"
)

// UserRepository provides user database operations
// @Description user repository
// @receiver ur
// @return UserRepository
func NewUserRepository() *UserRepository {
	return &UserRepository{db: db.GetDB()}
}

type UserRepository struct {
	db *sql.DB
}

func (ur *UserRepository) FindAllUsers() ([]models.User, error) {
	// implement user database query
	return nil, nil
}