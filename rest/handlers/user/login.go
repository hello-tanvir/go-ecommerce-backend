package user

import (
	"ecommerce/util"
	"encoding/json"
	"net/http"
)

type ReqLogin struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}

func (h *Handler) Login(w http.ResponseWriter, r *http.Request) {

	var req ReqLogin

	decoder := json.NewDecoder(r.Body)

	err := decoder.Decode(&req)

	if err != nil {
		http.Error(w, "Invalid Request Data", http.StatusBadRequest)
		return
	}

	usr, err := h.userRepo.Find(req.Email, req.Password)
	if err != nil {
		util.SendError(w, http.StatusUnauthorized, "Unauthorizes")
		return
	}

	accessToken, err := util.CreateJwt(h.cnf.JwtSecretkey, util.Payload{
		Sub:       usr.ID,
		FirstName: usr.FirstName,
		LastName:  usr.LastName,
		Email:     usr.Email,
	})

	if err != nil {
		http.Error(w, "Internal server Error", http.StatusInternalServerError)
		return
	}

	util.SendData(w, accessToken, http.StatusCreated)
}
