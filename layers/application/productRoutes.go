package application

import (
	"encoding/json"
	"errors"
	"net/http"

	"github.com/Hadis2971/go_web/layers/dataAccess"
	"github.com/Hadis2971/go_web/layers/domain"
	"github.com/Hadis2971/go_web/middlewares"
	"github.com/Hadis2971/go_web/models"
)

type ProductRoutes struct {
	mux *http.ServeMux
	productDomain *domain.ProductDomain
}

func NewProductRoutes(productDomain *domain.ProductDomain) *ProductRoutes {
	return &ProductRoutes{mux: http.NewServeMux(), productDomain: productDomain}
}

func (pr *ProductRoutes) HandleCreateProduct(w http.ResponseWriter, r *http.Request) {
	var createProductJsonBoby models.Product

	err := json.NewDecoder(r.Body).Decode(&createProductJsonBoby)

	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)

		return
	}

	err = pr.productDomain.HandleCreateProduct(createProductJsonBoby)
	

	if errors.Is(err, dataAccess.ErrorCreateProduct) {
		http.Error(w, dataAccess.ErrorCreateProduct.Error(), http.StatusInternalServerError)

		return 
	}

	if errors.Is(err, dataAccess.ErrorCreateProductMissingFields) {
		http.Error(w, dataAccess.ErrorCreateProductMissingFields.Error(), http.StatusBadRequest)

		return
	}

	w.WriteHeader(http.StatusOK)

}

func (pd *ProductRoutes) HandleGetAllProducts(w http.ResponseWriter, r *http.Request) {
	type Response struct {
		Products []models.Product `json:"products"`
	}

	products, err := pd.productDomain.HandleGetAllProducts()

	if errors.Is(err, dataAccess.ErrorGetAllProducts) {
		http.Error(w, dataAccess.ErrorGetAllProducts.Error(), http.StatusBadRequest)
	}

	productsJson, err := json.Marshal(&Response{Products: products})

	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)

		return
	}

	w.WriteHeader(http.StatusOK)
	w.Header().Set("Content-Type", "application/json")
	w.Write(productsJson)
}

func (pd *ProductRoutes) HandleGetProductById(w http.ResponseWriter, r *http.Request) {
	type Response struct {
		Product models.Product `json:"products"`
	}

	type GetProductByIdJsonBody struct {
		ID models.ProductId `json:"id"`
	}

	var getProductByIdJsonBody GetProductByIdJsonBody

	

	err := json.NewDecoder(r.Body).Decode(&getProductByIdJsonBody)

	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)

		return
	}

	product, err := pd.productDomain.HandleGetProductById(getProductByIdJsonBody.ID)

	if errors.Is(err, dataAccess.ErrorGetProductById) {
		http.Error(w, dataAccess.ErrorGetProductById.Error(), http.StatusInternalServerError)

		return 
	}

	if errors.Is(err, dataAccess.ErrorGetProductByIdMissingId) {
		http.Error(w, dataAccess.ErrorGetProductByIdMissingId.Error(), http.StatusBadRequest)
		
		return 
	}

	productJson, err := json.Marshal(&product)

	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)

		return
	}

	w.WriteHeader(http.StatusOK)
	w.Header().Set("Content-Type", "application/json")
	w.Write(productJson)
}

func (pr *ProductRoutes) HandleDeleteProduct(w http.ResponseWriter, r *http.Request) {
	type DeleteProductJsonBody struct {
		ID models.ProductId `json:"id"`
	}

	var deleteProductJsonBody DeleteProductJsonBody

	err := json.NewDecoder(r.Body).Decode(&deleteProductJsonBody)

	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)

		return
	}

	err = pr.productDomain.HandleDeleteProduct(deleteProductJsonBody.ID)

	if errors.Is(err, dataAccess.ErrorDeleteProduct) {
		http.Error(w, dataAccess.ErrorDeleteProduct.Error(), http.StatusInternalServerError)

		return 
	}

	if errors.Is(err, dataAccess.ErrorDeleteProductMissingId) {
		http.Error(w, dataAccess.ErrorDeleteProductMissingId.Error(), http.StatusBadRequest)
		
		return 
	}

	w.WriteHeader(http.StatusOK)
}

func (pr *ProductRoutes) HandleUpdateProduct(w http.ResponseWriter, r *http.Request) {
	var updateProductJsonBody models.Product

	err := json.NewDecoder(r.Body).Decode(&updateProductJsonBody)

	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)

		return
	}

	err = pr.productDomain.HandleUpdateProduct(updateProductJsonBody)

	if errors.Is(err, dataAccess.ErrorUpdateProduct) {
		http.Error(w, dataAccess.ErrorUpdateProduct.Error(), http.StatusInternalServerError)

		return 
	}

	if errors.Is(err, dataAccess.ErrorUpdateProductMissingFields) {
		http.Error(w, dataAccess.ErrorUpdateProductMissingFields.Error(), http.StatusBadRequest)
		
		return 
	}

	w.WriteHeader(http.StatusOK)
}

func (pr *ProductRoutes) RegisterRoutes() *http.ServeMux {
	authMiddleware := middlewares.NewAuthMiddleware()

	pr.mux.HandleFunc("POST /create/", authMiddleware.WithHttpRouthAuthentication(pr.HandleCreateProduct))
	pr.mux.HandleFunc("POST /list/", authMiddleware.WithHttpRouthAuthentication(pr.HandleGetAllProducts))
	pr.mux.HandleFunc("POST /list/product", authMiddleware.WithHttpRouthAuthentication(pr.HandleGetProductById))
	pr.mux.HandleFunc("POST /delete/", authMiddleware.WithHttpRouthAuthentication(pr.HandleDeleteProduct))
	pr.mux.HandleFunc("POST /update/", authMiddleware.WithHttpRouthAuthentication(pr.HandleUpdateProduct))


	return pr.mux
}