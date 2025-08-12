package server

import (
	"ecommerce/models"
	"encoding/json"
	"io"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
)

type HttpHandlers struct {
	repository models.Repository
}

func NewHttpHandlers(repo models.Repository) *HttpHandlers {
	return &HttpHandlers{
		repository: repo,
	}
}

func (hh *HttpHandlers) _saveProduct(w http.ResponseWriter, r *http.Request) {
	b, err := io.ReadAll(r.Body)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte(err.Error()))
		return
	}
	var pModel = models.NewProductModel()
	err = json.Unmarshal(b, pModel)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte(err.Error()))
		return
	}

	err = models.ValidateProduct(pModel)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte(err.Error()))
		return
	}

	err = hh.repository.SaveProduct(pModel)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte(err.Error()))
		return
	}

	pJson, err := json.Marshal(pModel)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte(err.Error()))
		return
	}
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write(pJson)
}

func (hh *HttpHandlers) GetProductByIdHandler(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	var id, err = strconv.ParseInt(params["id"], 10, 64)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte(err.Error()))
		return
	}
	pModel, err := hh.repository.GetProductById(uint64(id))
	if err != nil {
		w.WriteHeader(http.StatusNotFound)
		w.Write([]byte(err.Error()))
		return
	}

	pJson, err := json.Marshal(pModel)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte(err.Error()))
		return
	}
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write(pJson)
}

func (hh *HttpHandlers) _getProducts(byCat string, w http.ResponseWriter, r *http.Request) {
	var pList []*models.ProductModel
	var err error
	if byCat != "" {
		pList, err = hh.repository.GetProductsByCat(byCat)
	} else {
		pList, err = hh.repository.GetProducts()
	}
	if err != nil {
		w.WriteHeader(http.StatusNotFound)
		w.Write([]byte(err.Error()))
		return
	}

	pJson, err := json.Marshal(pList)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte(err.Error()))
		return
	}
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write(pJson)
}

func (hh *HttpHandlers) ProductsHandler(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case http.MethodGet:
		hh._getProducts(r.URL.Query().Get("category"), w, r)
	case http.MethodPost:
		hh._saveProduct(w, r)
	}

}
