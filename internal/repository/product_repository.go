package repository

import "github.com/destafajri/golang-fiber/internal/entity"


type ProductRepository interface {
	Insert(product entity.Product)

	FindAll() (products []entity.Product)

	DeleteAll()
}
