package product

import (
	"ecommerce/rest/middleware"
	"net/http"
)

func (h *Handler) RegisterRoutes(
	mux *http.ServeMux,
	manager *middleware.Manager,
) {

	manager.Handle(
		mux,
		"GET /products",
		h.GetProducts,
	)

	manager.Handle(
		mux,
		"POST /products",
		h.CreateProduct,
		h.middleware.AuthenticateJwt,
	)

	manager.Handle(
		mux,
		"GET /products/{id}",
		h.GetProduct,
	)

	manager.Handle(
		mux,
		"PUT /products/{id}",
		h.UpdateProduct,
		h.middleware.AuthenticateJwt,
	)

	manager.Handle(
		mux,
		"DELETE /products/{id}",
		h.DeleteProduct,
		h.middleware.AuthenticateJwt,
	)

}
