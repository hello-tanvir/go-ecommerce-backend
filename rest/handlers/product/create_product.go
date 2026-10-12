package product

import (
	"ecommerce/repo"
	"ecommerce/util"
	"encoding/json"
	"net/http"
)

type ReqCreateProduct struct {
	Id         int     `json:"id"`
	Title      string  `json:"title"`
	Dscription string  `json:"dscription"`
	Price      float64 `json:"price"`
	ImgUrl     string  `json:"image"`
}

func (h *Handler) CreateProduct(w http.ResponseWriter, r *http.Request) {

	var req ReqCreateProduct

	decoder := json.NewDecoder(r.Body)
	err := decoder.Decode(&req)

	if err != nil {
		http.Error(w, "Plz give me valid json", 400)
		return
	}

	created, err := h.productRepo.Create(repo.Product{
		Title:      req.Title,
		Dscription: req.Dscription,
		Price:      req.Price,
		ImgUrl:     req.ImgUrl,
	})

	if err != nil {
		http.Error(w, "Internal Server Error", http.StatusInternalServerError)

	}

	util.SendData(w, created, http.StatusCreated)

}
