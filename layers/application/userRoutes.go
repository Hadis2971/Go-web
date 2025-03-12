package application

import (
	"encoding/json"
	"net/http"

	"github.com/Hadis2971/go_web/layers/dataAccess"
	"github.com/Hadis2971/go_web/layers/domain"
)

type UserRouteHandler struct {
	mux        *http.ServeMux
	userDomain *domain.UserDomain
}

type DeleteUserJsonBody struct {
	ID int `json:"id"`
}

func NewUserRouteHandler(mux *http.ServeMux, userDomain *domain.UserDomain) *UserRouteHandler {
	return &UserRouteHandler{mux: mux, userDomain: userDomain}
}

func (ur UserRouteHandler) HandleDeleteUser(w http.ResponseWriter, r *http.Request) {
	var deleteUserJsonBody DeleteUserJsonBody

	// There is a security issue here. You're only reading ID from r.Body, but r.Body can contain whatever the end-user sends. Which could be TB of data!
	// Use io.LimitReader to protect yourself
	if err := json.NewDecoder(r.Body).Decode(&deleteUserJsonBody); err != nil { // Using a mix of `if err != nil` and `; err != nil`. Better to keep the project consistent
		http.Error(w, err.Error(), http.StatusBadRequest)
	}

	err := ur.userDomain.HandleDeleteUser(deleteUserJsonBody.ID)

	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest) // Need to differentiate between a bad request and a fatal error. Like the database was offline. Otherwise the frontender will be confused. You can differentiate with an error enum and use errors.Is()
	}

	w.WriteHeader(http.StatusOK)
}

func (ur UserRouteHandler) HandleUpdateUser(w http.ResponseWriter, r *http.Request) {
	var updateUserRequestJsonBody dataAccess.UpdateUserRequest

	if err := json.NewDecoder(r.Body).Decode(&updateUserRequestJsonBody); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)

		return
	} else if err := ur.userDomain.HandleUpdateUser(updateUserRequestJsonBody); err != nil { // golang is a bit different here. If you want, you can change your else/if together, they will execute in sequence until they reach an error or the end
		// This would be different styling in the same project though, so I would choose something and be consistent.
		// A more complex example here: https://github.com/SilicalNZ/wikia/blob/master/services/discord/controllers/entrypoint/main.go#L96
		http.Error(w, err.Error(), http.StatusInternalServerError)

		return
	}

	w.WriteHeader(http.StatusOK)
}

func (ur UserRouteHandler) RegisterRoutes() {
	// It would be more idiomatic to have (ur *UserRouteHandler) * means you're modifying the struct
	// It's also more performant, as * is using the pointer. Without the * its creating a copy everytime. Which you don't need a copy.
	ur.mux.HandleFunc("POST /delete/", ur.HandleDeleteUser) // There's no authentication for these endpoints
	ur.mux.HandleFunc("POST /update/", ur.HandleUpdateUser) // Can always do DELETE and UPDATE methods instead of full endpoints. Fine either way.
}
