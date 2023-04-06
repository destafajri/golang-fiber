package repository

import (
	"database/sql"

	"github.com/destafajri/golang-fiber/internal/entity"
)

type UserRepository interface {
	Register(users *entity.UserEntity) error
	GetData(phone string) (*entity.UserEntity, error)
	Login(phone string) (*entity.UserEntity, error)
}

type userImplementation struct {
	db *sql.DB
}

func NewUserRepository(db *sql.DB) UserRepository{
	return &userImplementation{
		db: db,
	}
}