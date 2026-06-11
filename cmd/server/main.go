package main

import (
	"fmt"
	"net/http"

	"product-system/internal/database"
	"product-system/internal/product"
)

func main() {

	db, err := database.NewMySQL()

	if err != nil {
		panic(err)
	}

	repo := product.NewRepository(db)

	service := product.NewService(repo)

	handler := product.NewHandler(service)

	http.HandleFunc(
		"POST /products",
		handler.CreateProduct,
	)

	http.HandleFunc(
		"/health",
		func(w http.ResponseWriter, r *http.Request) {
			fmt.Fprintln(w, "OK")
		},
	)

	fmt.Println("server started on :8080")

	http.ListenAndServe(":8080", nil)
}