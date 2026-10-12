package product

import (
	"ecommerce/util"
	"net/http"
)

func (h *Handler) GetProducts(w http.ResponseWriter, r *http.Request) {

	productList, err := h.productRepo.List()
	if err != nil {
		http.Error(w, "Invalid Server Error", http.StatusInternalServerError)
		return
	}

	util.SendData(w, productList, http.StatusOK)

}
