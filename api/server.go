package api

import (
	"net/http"

	"github.com/anthdm/weavebox"
	"github.com/esojlm5/ggcommerce/types"
)

type ProductHandler struct {
}

func (h *ProductHandler) HandleGetProduct(c *weavebox.Context) error {
	return c.JSON(http.StatusOK, &types.Product{SKU: "shoe-111"})
}
