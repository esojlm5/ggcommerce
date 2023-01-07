package main

import (
	"github.com/anthdm/weavebox"
	"github.com/esojlm5/ggcommerce/api"
)

func main() {
	app := weavebox.New()

	product := &api.ProductHandler{}
	app.Get("/product", product.HandleGetProduct)

	app.Serve(3001)
}
