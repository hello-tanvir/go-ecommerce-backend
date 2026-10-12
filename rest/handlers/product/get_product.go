package product

import (
	"ecommerce/util"
	"net/http"
	"strconv"
)

func (h *Handler) GetProduct(w http.ResponseWriter, r *http.Request) {
	productID := r.PathValue("id")

	pId, err := strconv.Atoi(productID)

	if err != nil {
		http.Error(w, "Plz give me a valid product id ", 400)
		return
	}

	product, err := h.productRepo.Get(pId)
	if err != nil {
		http.Error(w, "Internal server Error ", http.StatusInternalServerError)
		return
	}

	if product == nil {
		util.SendError(w, 404, "Product not found")
		return
	}

	util.SendData(w, product, 200)

}
