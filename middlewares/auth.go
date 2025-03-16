package middlewares

import (
	"fmt"
	"net/http"

	"github.com/Hadis2971/go_web/util"
	"github.com/golang-jwt/jwt/v5"
)


func verifyToken(tokenString string) error {
	secret := util.GetEnvVariable("JWT_LOGIN_TOKEN_SECRET")

	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
	   return secret, nil
	})
   
	if err != nil {
	   return err
	}
   
	if !token.Valid {
	   return fmt.Errorf("invalid token")
	}
   
	return nil
 }

type AuthMiddleware struct {}


func NewAuthMiddleware () *AuthMiddleware {
	return &AuthMiddleware{}
}

func (am *AuthMiddleware) Authorized(httpHandler http.HandlerFunc) http.HandlerFunc {
	return func (w http.ResponseWriter, r *http.Request) {
		authorization := r.Header.Get("Authorization")

		if err := verifyToken(authorization); err != nil {
			http.Error(w, err.Error(), http.StatusUnauthorized)

			return
		}

		httpHandler(w, r);
	}
}


