package server

import (
	"fmt"
	"net/http"

	"github.com/gorilla/mux"
)

func RunServer(port int, handlers *HttpHandlers) {
	fmt.Printf("running http server on :%d\n", port)
	r := mux.NewRouter()
	r.HandleFunc("/greet", func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusOK)
		w.Write([]byte("Hello requester!"))
	})

	r.HandleFunc("/products/{id}", handlers.GetProductByIdHandler).Methods(http.MethodGet)
	r.HandleFunc("/products", handlers.ProductsHandler).Methods(http.MethodGet, http.MethodPost)
	fmt.Println("http path [get]: /products/{id}")
	fmt.Println("http path [get, post]: /products?category={category}")
	http.Handle("/", r)
	var err = http.ListenAndServe(fmt.Sprintf(":%d", port), nil)
	if err != nil {
		panic(err)
	}
}
