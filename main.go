package main

import (
	"github.com/anthdm/weavebox"
	"github.com/esojlm5/ggcommerce/api"
)

func main() {
	app := weavebox.New()
	admin := app.Box("/admin")

	productHandler := &api.ProductHandler{}

	admin.Get("/product", productHandler.HandleGetProduct)
	admin.Post("/product", productHandler.HandlePostProduct)

	app.Serve(3001)
}
