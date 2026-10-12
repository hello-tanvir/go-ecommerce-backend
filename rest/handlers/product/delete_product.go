package product

import (
	"ecommerce/util"
	"net/http"
	"strconv"
)

func (h *Handler) DeleteProduct(w http.ResponseWriter, r *http.Request) {

	productID := r.PathValue("id")

	pId, err := strconv.Atoi(productID)
	if err != nil {
		http.Error(w, "Plz give me a valid product id", 400)
		return
	}

	err = h.productRepo.Delete(pId)
	if err != nil {
		http.Error(w, "Internal Server Error", http.StatusInternalServerError)
		return
	}

	util.SendData(w, map[string]string{
		"message": "Product deleted successfully",
	}, 200)
}
