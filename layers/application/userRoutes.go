package application

import (
	"encoding/json"
	"errors"
	"fmt"
	"net/http"

	"github.com/Hadis2971/go_web/layers/dataAccess"
	"github.com/Hadis2971/go_web/layers/domain"
	"github.com/Hadis2971/go_web/middlewares"
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

	fmt.Println(deleteUserJsonBody)
	
	if err := json.NewDecoder(r.Body,).Decode(&deleteUserJsonBody); err != nil {
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
		// golang is a bit different here. If you want, you can change your else/if together, they will execute in sequence until they reach an error or the end
		// This would be different styling in the same project though, so I would choose something and be consistent.
		// A more complex example here: https://github.com/SilicalNZ/wikia/blob/master/services/discord/controllers/entrypoint/main.go#L96

		// Hadis => I didnt quite understand this comment maybe a call would be nice here.
		http.Error(w, err.Error(), http.StatusInternalServerError)

		return
	}

	w.WriteHeader(http.StatusOK)
}

func (ur *UserRouteHandler) RegisterRoutes() *http.ServeMux {
	authMiddleware := middlewares.NewAuthMiddleware()

	ur.mux.HandleFunc("POST /delete/", authMiddleware.WithHttpRouthAuthentication(ur.HandleDeleteUser))
	ur.mux.HandleFunc("POST /update/", authMiddleware.WithHttpRouthAuthentication(ur.HandleUpdateUser))

	return ur.mux
}
