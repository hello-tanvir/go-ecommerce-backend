package user

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
		"POST /users",
		h.CreateUser,
	)

	manager.Handle(
		mux,
		"POST /users/login",
		h.Login,
	)

}
