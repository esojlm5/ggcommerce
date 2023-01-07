package api

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/anthdm/weavebox"
	"github.com/esojlm5/ggcommerce/types"
)

type CreateProductRequest struct {
	SKU  string `json:"SKU"`
	Name string `json:"name"`
}

type ProductStorer interface {
	Insert(*types.Product) error
	GetById(string) (*types.Product, error)
}

type ProductHandler struct {
	store ProductStorer
}

func NewProductHandler(pStore ProductStorer) *ProductHandler {
	return &ProductHandler{
		store: pStore,
	}
}

func (h *ProductHandler) HandlePostProduct(c *weavebox.Context) error {
	productReq := &CreateProductRequest{}
	if err := json.NewDecoder(c.Request().Body).Decode(productReq); err != nil {
		return err
	}
	fmt.Println("CreateProductRequest", productReq)
	return c.JSON(http.StatusOK, productReq)
}

func (h *ProductHandler) HandleGetProduct(c *weavebox.Context) error {
	return c.JSON(http.StatusOK, &types.Product{SKU: "shoe-111"})
}
