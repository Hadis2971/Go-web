package application

import (
	"encoding/json"
	"errors"
	"io"
	"net/http"

	"github.com/Hadis2971/go_web/layers/dataAccess"
	"github.com/Hadis2971/go_web/layers/domain"
	"github.com/Hadis2971/go_web/middlewares"
	"github.com/Hadis2971/go_web/models"
)

type UserRouteHandler struct {
	mux        *http.ServeMux
	userDomain *domain.UserDomain
}

func NewUserRouteHandler(userDomain *domain.UserDomain) *UserRouteHandler {
	return &UserRouteHandler{mux: http.NewServeMux(), userDomain: userDomain}
}

func (ur UserRouteHandler) HandleDeleteUser(w http.ResponseWriter, r *http.Request) {
	type DeleteUserJsonBody struct {
		ID int `json:"id"`
	}

	const MAX_SIZE_FOR_REQUEST_PAYLOAD = 64
	var deleteUserJsonBody DeleteUserJsonBody
	
	if err := json.NewDecoder(io.LimitReader(r.Body, MAX_SIZE_FOR_REQUEST_PAYLOAD)).Decode(&deleteUserJsonBody); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)

		return
	}

	if err := ur.userDomain.HandleDeleteUser(deleteUserJsonBody.ID); err != nil {
		
		if errors.Is(err, dataAccess.InternalServerError) {
			http.Error(w, err.Error(), http.StatusInternalServerError)

		}

		if errors.Is(err, dataAccess.ErrorMissingID) {
			http.Error(w, err.Error(), http.StatusBadRequest)

		}

		return
	}

	w.WriteHeader(http.StatusOK)
}

func (ur UserRouteHandler) HandleUpdateUser(w http.ResponseWriter, r *http.Request) {
	var updateUserRequestJsonBody dataAccess.UpdateUserRequest

	if err := json.NewDecoder(r.Body).Decode(&updateUserRequestJsonBody); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)

		return
	} else if err = ur.userDomain.HandleUpdateUser(updateUserRequestJsonBody); err != nil { 
		
		http.Error(w, err.Error(), http.StatusInternalServerError)

		return
	}

	w.WriteHeader(http.StatusOK)
}

func (ur UserRouteHandler) HandleGetAllUsersAndTheirOrders(w http.ResponseWriter, r *http.Request) {
	type Response struct {
		UserWithOrders []models.UserWithOrders `json:"user_and_orders"`
	}

	userWithOrders, err := ur.userDomain.HandleGetAllUsersAndTheirOrders()

	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)

		return
	}

	
	userAndOrdersJson, err := json.Marshal(&Response{UserWithOrders: userWithOrders})


	w.WriteHeader(http.StatusOK)
	w.Header().Set("Content-Type", "application/json")
	w.Write(userAndOrdersJson)
}

func (ur *UserRouteHandler) RegisterRoutes() *http.ServeMux {
	authMiddleware := middlewares.NewAuthMiddleware()

	ur.mux.HandleFunc("POST /delete", authMiddleware.WithHttpRouthAuthentication(ur.HandleDeleteUser))
	ur.mux.HandleFunc("POST /update", authMiddleware.WithHttpRouthAuthentication(ur.HandleUpdateUser))
	ur.mux.HandleFunc("GET /product_orders", authMiddleware.WithHttpRouthAuthentication(ur.HandleGetAllUsersAndTheirOrders))

	return ur.mux
}