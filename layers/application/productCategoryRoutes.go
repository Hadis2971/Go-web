package application

import (
	"encoding/json"
	"net/http"

	"github.com/Hadis2971/go_web/layers/domain"
	"github.com/Hadis2971/go_web/models"
)

type ProductCategoryRoutes struct {
	mux *http.ServeMux
	productCategoryDomain *domain.ProductCategoryDomain
}

func NewProductCategoryRoutes(productCategoryDomain *domain.ProductCategoryDomain) *ProductCategoryRoutes {
	return &ProductCategoryRoutes{mux: http.NewServeMux(), productCategoryDomain: productCategoryDomain}
}

func (pcr ProductCategoryRoutes) HandleGetAllProductCategory(w http.ResponseWriter, r *http.Request) {
	productCategories, err := pcr.productCategoryDomain.HandleGetAllProductCategory()

	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)

		return
	}

	productCategoriesJSON, err := json.Marshal(productCategories)

	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)

		return
	}

	w.WriteHeader(http.StatusOK)
	w.Header().Set("Content-Type", "application/json")
	w.Write(productCategoriesJSON)
}

func (pcr ProductCategoryRoutes) HandleCreateProductCategory(w http.ResponseWriter, r *http.Request) {
	

	var request models.NewProductCategoryRequst

	err := json.NewDecoder(r.Body).Decode(&request)

	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)

		return
	}

	err = pcr.productCategoryDomain.HandleCreateProductCategory(request)

	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)

		return
	}

	w.WriteHeader(http.StatusOK)
}

func (pcr ProductCategoryRoutes) HandleUpdateProductCategory(w http.ResponseWriter, r *http.Request) {

	var request models.ProductCategory

	err := json.NewDecoder(r.Body).Decode(&request)

	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)

		return
	}

	err = pcr.productCategoryDomain.HandleUpdateProductCategory(request)

	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)

		return
	}

	w.WriteHeader(http.StatusOK)
}

func (pcr ProductCategoryRoutes) HandleDeleteProductCategory(w http.ResponseWriter, r *http.Request) {

	var productCategoryId models.ProductCategoryId

	err := json.NewDecoder(r.Body).Decode(&productCategoryId)

	err = pcr.productCategoryDomain.HandleDeleteProductCategory(productCategoryId)

	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)

		return
	}

	w.WriteHeader(http.StatusOK)
}

func (pcr ProductCategoryRoutes) RegisterRoutes() *http.ServeMux {
	pcr.mux.HandleFunc("POST /create", pcr.HandleCreateProductCategory)
	pcr.mux.HandleFunc("GET /list", pcr.HandleGetAllProductCategory)
	pcr.mux.HandleFunc("POST /update", pcr.HandleUpdateProductCategory)
	pcr.mux.HandleFunc("POST /delete", pcr.HandleDeleteProductCategory)

	return pcr.mux
}