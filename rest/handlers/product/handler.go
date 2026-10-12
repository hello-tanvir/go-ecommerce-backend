package product

import (
	"ecommerce/repo"
	"ecommerce/rest/middleware"
)

type Handler struct {
	middleware  *middleware.Middlewares
	productRepo repo.ProductRepo
}

func NewHandler(middleware *middleware.Middlewares, productRepo repo.ProductRepo) *Handler {
	return &Handler{
		middleware:  middleware,
		productRepo: productRepo,
	}
}
