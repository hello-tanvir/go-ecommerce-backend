package user

import (
	"ecommerce/repo"
	"ecommerce/util"
	"encoding/json"
	"net/http"
	"strings"
)

type ReqCreateUser struct {
	FirstName   string `json:"first_name"`
	LastName    string `json:"last_name"`
	Email       string `json:"email"`
	Password    string `json:"password"`
	IsShopOwner bool   `json:"is_shop_owner"`
}

func (h *Handler) CreateUser(w http.ResponseWriter, r *http.Request) {
	var req ReqCreateUser
	decoder := json.NewDecoder(r.Body)

	err := decoder.Decode(&req)

	if err != nil {

		http.Error(w, "Invalid Request Data", http.StatusBadRequest)
		return

	}

	use, err := h.userRepo.Create(repo.User{
		FirstName:   req.FirstName,
		LastName:    req.LastName,
		Email:       req.Email,
		Password:    req.Password,
		IsShopOwner: req.IsShopOwner,
	})
	if err != nil {
		if strings.Contains(err.Error(), "users_email_key") {
			http.Error(w, "Email already exists", http.StatusConflict)
			return
		}

		http.Error(w, "Internal Server Error", http.StatusInternalServerError)
		return
	}

	util.SendData(w, use, http.StatusCreated)

}
