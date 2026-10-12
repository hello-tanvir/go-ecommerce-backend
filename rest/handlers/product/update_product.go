package product

import (
	"ecommerce/repo"
	"ecommerce/util"
	"encoding/json"
	"net/http"
	"strconv"
)

type ReqUpdateProduct struct {
	Id         int     `json:"id"`
	Title      string  `json:"title"`
	Dscription string  `json:"dscription"`
	Price      float64 `json:"price"`
	ImgUrl     string  `json:"image"`
}

func (h *Handler) UpdateProduct(w http.ResponseWriter, r *http.Request) {
	productID := r.PathValue("id")

	pId, err := strconv.Atoi(productID)
	if err != nil {
		http.Error(w, "Plz give me a valid product id", 400)
		return
	}

	var req ReqUpdateProduct
	decoder := json.NewDecoder(r.Body)
	err = decoder.Decode(&req)
	if err != nil {
		http.Error(w, "Plz give me valid json", http.StatusBadRequest)
		return
	}
	updateProduct, err := h.productRepo.Update(repo.Product{
		Id:         pId,
		Title:      req.Title,
		Dscription: req.Dscription,
		Price:      req.Price,
		ImgUrl:     req.ImgUrl,
	})

	if err != nil {
		http.Error(w, "Invalid Server Error", http.StatusInternalServerError)
		return

	}

	util.SendData(w, updateProduct, 200)
}
