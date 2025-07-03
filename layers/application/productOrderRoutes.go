package application

import (
	"encoding/json"
	"errors"
	"net/http"

	"github.com/google/uuid"

	"github.com/Hadis2971/go_web/layers/domain"
	"github.com/Hadis2971/go_web/middlewares"
	"github.com/Hadis2971/go_web/models"
)

type ProductOrderRoutes struct {
	productOrderDomain *domain.ProductOrderDomain
	mux *http.ServeMux
}

var (
	ErrorHandleCreateProductOrder = errors.New("Server Error Create Product")
)

func NewProductOrderRoutes(productOrderDomain *domain.ProductOrderDomain) *ProductOrderRoutes {
	return &ProductOrderRoutes{productOrderDomain: productOrderDomain, mux: http.NewServeMux()}
}

func (por *ProductOrderRoutes) HandleCreateProductOrder(w http.ResponseWriter, r *http.Request) {
	var productOrderReqPayload models.ProductOrderRequestPayload

	err := json.NewDecoder(r.Body).Decode(&productOrderReqPayload)

	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	orderId := models.OrderId(uuid.New().String())

	newProductOrder := models.ProductOrder{OrderId: orderId, Quantity: productOrderReqPayload.Quantity, UserId: models.UserId(productOrderReqPayload.UserId), ProductId: models.ProductId(productOrderReqPayload.ProductId)}

	err = por.productOrderDomain.HandleCreateProductOrder(newProductOrder)

	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	w.WriteHeader(http.StatusOK)
}

func (por *ProductOrderRoutes) HandleCreateProductOrderWithMultipleProducts(w http.ResponseWriter, r *http.Request) {
	var productOrderReqPayload []models.ProductOrderRequestPayload
	newProductOrders := []models.ProductOrder{}
	orderId := models.OrderId(uuid.New().String())

	err := json.NewDecoder(r.Body).Decode(&productOrderReqPayload)

	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)

		return
	}

	for _, val := range productOrderReqPayload {
		newProductOrders = append(newProductOrders, models.ProductOrder{OrderId: orderId, Quantity: val.Quantity, UserId: models.UserId(val.UserId), ProductId: models.ProductId(val.ProductId)})
	}

	err = por.productOrderDomain.HandleCreateProductOrderWithMultipleProducts(newProductOrders)

	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
	}

	w.WriteHeader(http.StatusOK)

}


func (por *ProductOrderRoutes) HandleGetProuctOrdersByUserId(w http.ResponseWriter, r *http.Request) {
	type ReqPayload struct {
		UserId int `json:"user_id"`
	}

	type Response struct {
		ProductOrders []models.ProductOrder `json:"product_orders"`
	}

	var reqPayload ReqPayload

	json.NewDecoder(r.Body).Decode(&reqPayload)

	productOrders, err := por.productOrderDomain.HandleGetProuctOrdersByUserId(models.UserId(reqPayload.UserId))

	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	productOrdersJson, err := json.Marshal(&Response{ProductOrders: productOrders})

	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusOK)
	w.Header().Set("Content-Type", "application/json")
	w.Write(productOrdersJson)
}

func (por *ProductOrderRoutes) HandleGetProuctOrdersByOrderId(w http.ResponseWriter, r *http.Request) {
	type ReqPayload struct {
		OrderId models.OrderId `json:"order_id"`
	}

	type Response struct {
		ProductOrders []models.ProductOrder `json:"product_orders"`
	}

	var reqPayload ReqPayload

	json.NewDecoder(r.Body).Decode(&reqPayload)

	productOrders, err := por.productOrderDomain.HandleGetProductOrdersByOrderId(models.OrderId(reqPayload.OrderId))

	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	productOrdersJson, err := json.Marshal(&Response{ProductOrders: productOrders})

	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusOK)
	w.Header().Set("Content-Type", "application/json")
	w.Write(productOrdersJson)
}

func (por *ProductOrderRoutes) HandleGetProductOrdersByUserIdAndOrderId(w http.ResponseWriter, r *http.Request) {
	type ReqPayload struct {
		UserId models.UserId `json:"user_id"`
		OrderId models.OrderId `json:"order_id"`
	}

	type Response struct {
		ProductOrdersAndUser []models.ProductAndUser `json:"product_orders_and_user"`
	}

	var reqPayload ReqPayload

	json.NewDecoder(r.Body).Decode(&reqPayload)

	productOrdersAndUser, err := por.productOrderDomain.HandleGetProductOrdersByUserIdAndOrderId(reqPayload.UserId, reqPayload.OrderId)

	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)

		return
	}

	jsonProductOrdersAndUser, err := json.Marshal(&Response{ProductOrdersAndUser: productOrdersAndUser})

	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)

		return
	}

	w.WriteHeader(http.StatusOK)
	w.Header().Set("Content-Type", "application/json")
	w.Write(jsonProductOrdersAndUser)
}

func (por *ProductOrderRoutes) HandleUpdateProductOrder(w http.ResponseWriter, r *http.Request) {
	var productOrderReqPayload models.ProductOrder

	err := json.NewDecoder(r.Body).Decode(&productOrderReqPayload)

	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	newProductOrder := models.ProductOrder{ID: productOrderReqPayload.ID, Quantity: productOrderReqPayload.Quantity, UserId: models.UserId(productOrderReqPayload.UserId), ProductId: models.ProductId(productOrderReqPayload.ProductId)}

	err = por.productOrderDomain.HandleUpdateProductOrder(newProductOrder)

	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	w.WriteHeader(http.StatusOK)
}

func (por *ProductOrderRoutes) HandleDeleteProductOrder(w http.ResponseWriter, r *http.Request) {
	type ReqPayload struct {
		ID int `json:"id"`
	}

	var reqPayload ReqPayload

	err := json.NewDecoder(r.Body).Decode(&reqPayload)

	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	

	err = por.productOrderDomain.HandleDeleteProductOrder(models.ProductOrderId(reqPayload.ID))

	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	w.WriteHeader(http.StatusOK)
}

func (por *ProductOrderRoutes) RegisterRoutes() *http.ServeMux {
	authMiddleware := middlewares.NewAuthMiddleware()

	por.mux.HandleFunc("POST /create", authMiddleware.WithHttpRouthAuthentication(por.HandleCreateProductOrder))
	por.mux.HandleFunc("POST /create_multiple", authMiddleware.WithHttpRouthAuthentication(por.HandleCreateProductOrderWithMultipleProducts))
	por.mux.HandleFunc("POST /list/user", authMiddleware.WithHttpRouthAuthentication(por.HandleGetProuctOrdersByUserId))
	por.mux.HandleFunc("GET /list/order", authMiddleware.WithHttpRouthAuthentication(por.HandleGetProuctOrdersByOrderId))
	por.mux.HandleFunc("POST /update", authMiddleware.WithHttpRouthAuthentication(por.HandleUpdateProductOrder))
	por.mux.HandleFunc("POST /delete", authMiddleware.WithHttpRouthAuthentication(por.HandleDeleteProductOrder))
	por.mux.HandleFunc("POST /list/user_order", authMiddleware.WithHttpRouthAuthentication(por.HandleGetProductOrdersByUserIdAndOrderId))

	return por.mux
}