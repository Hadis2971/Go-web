package application

import (
	"encoding/json"
	"errors"
	"net/http"
	"strconv"

	"github.com/Hadis2971/go_web/layers/dataAccess"
	"github.com/Hadis2971/go_web/layers/domain"
	"github.com/Hadis2971/go_web/layers/service"
	"github.com/Hadis2971/go_web/middlewares"
	"github.com/Hadis2971/go_web/models"
)

type ProductRoutes struct {
	mux *http.ServeMux
	productDomain *domain.ProductDomain
	wsProductDomain *domain.WsProductDomain
}

func NewProductRoutes(productDomain *domain.ProductDomain, wsProductDomain *domain.WsProductDomain) *ProductRoutes {
	return &ProductRoutes{mux: http.NewServeMux(), productDomain: productDomain, wsProductDomain: wsProductDomain}
}

func (pr ProductRoutes) HandleCreateProduct(w http.ResponseWriter, r *http.Request) {
	var createProductJsonBody models.ProductReqPayload

	err := json.NewDecoder(r.Body).Decode(&createProductJsonBody)

	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)

		return
	}

	newProduct, err := pr.productDomain.HandleCreateProduct(createProductJsonBody)
	

	if errors.Is(err, dataAccess.ErrorCreateProduct) {
		http.Error(w, dataAccess.ErrorCreateProduct.Error(), http.StatusInternalServerError)

		return 
	}

	if errors.Is(err, dataAccess.ErrorCreateProductMissingFields) {
		http.Error(w, dataAccess.ErrorCreateProductMissingFields.Error(), http.StatusBadRequest)

		return
	}

	wsMessage := service.ProductMessage{ID: createProductJsonBody.ID, Topic: "product_update_message", Product: *newProduct}

	pr.wsProductDomain.HandleWsProductBroadcastMsg(wsMessage)

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

	id := strconv.Itoa(int(deleteProductJsonBody.ID))

	wsMessage := service.ProductMessage{ID: id, Topic: "product_delete_message"}

	pr.wsProductDomain.HandleWsProductBroadcastMsg(wsMessage)

	w.WriteHeader(http.StatusOK)
}

func (pr *ProductRoutes) HandleUpdateProduct(w http.ResponseWriter, r *http.Request) {
	var updateProductJsonBody models.ProductReqPayload

	err := json.NewDecoder(r.Body).Decode(&updateProductJsonBody)

	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)

		return
	}

	updatedProduct, err := pr.productDomain.HandleUpdateProduct(updateProductJsonBody)

	if errors.Is(err, dataAccess.ErrorUpdateProduct) {
		http.Error(w, dataAccess.ErrorUpdateProduct.Error(), http.StatusInternalServerError)

		return 
	}

	if errors.Is(err, dataAccess.ErrorUpdateProductMissingFields) {
		http.Error(w, dataAccess.ErrorUpdateProductMissingFields.Error(), http.StatusBadRequest)
		
		return 
	}

	wsMessage := service.ProductMessage{ID: updateProductJsonBody.ID, Topic: "product_update_message", Product: *updatedProduct}

	pr.wsProductDomain.HandleWsProductBroadcastMsg(wsMessage)

	w.WriteHeader(http.StatusOK)
}

func (pr *ProductRoutes) RegisterRoutes() *http.ServeMux {
	authMiddleware := middlewares.NewAuthMiddleware()

	pr.mux.HandleFunc("POST /create", authMiddleware.WithHttpRouthAuthentication(pr.HandleCreateProduct))
	pr.mux.HandleFunc("GET /list", authMiddleware.WithHttpRouthAuthentication(pr.HandleGetAllProducts))
	pr.mux.HandleFunc("POST /list", authMiddleware.WithHttpRouthAuthentication(pr.HandleGetProductById))
	pr.mux.HandleFunc("POST /delete", authMiddleware.WithHttpRouthAuthentication(pr.HandleDeleteProduct))
	pr.mux.HandleFunc("POST /update", authMiddleware.WithHttpRouthAuthentication(pr.HandleUpdateProduct))


	return pr.mux
}