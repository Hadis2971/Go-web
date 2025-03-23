package middlewares

import (
	"errors"
	"net/http"

	"github.com/Hadis2971/go_web/util"
	"github.com/golang-jwt/jwt/v5"
	"golang.org/x/net/websocket"
)

var (
	ErrorInvalidToken = errors.New("Invalid Token!!!")
	ErrorParsingToken = errors.New("Error Parsing Token!!!")
	ErrorMissingToken = errors.New("Missing Token!!!")
)

func verifyToken(tokenString string) error {
	secret := util.GetEnvVariable("JWT_LOGIN_TOKEN_SECRET")

	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
	   return []byte(secret), nil
	})

	if (token == nil) {
		return ErrorMissingToken
	}
   
	if (!token.Valid) {
	   return ErrorInvalidToken
	}

	if (err != nil) {
		return ErrorParsingToken
	 }
   
	return nil
 }

type AuthMiddleware struct {}


func NewAuthMiddleware () *AuthMiddleware {
	return &AuthMiddleware{}
}

func (am *AuthMiddleware) WithHttpRouthAuthentication(httpHandler http.HandlerFunc) http.HandlerFunc {
	return func (w http.ResponseWriter, r *http.Request) {

		authorization := r.Header.Get("Authorization")

		err := verifyToken(authorization)

		if (errors.Is(err, ErrorMissingToken)) {
			http.Error(w, ErrorMissingToken.Error(), http.StatusUnauthorized)

			return
		}

		if (errors.Is(err, ErrorInvalidToken)) {
			http.Error(w, ErrorInvalidToken.Error(), http.StatusUnauthorized)

			return
		}

		if (errors.Is(err, ErrorParsingToken)) {
			http.Error(w, ErrorParsingToken.Error(), http.StatusUnauthorized)

			return
		}

		httpHandler(w, r);
	}
}

func (am *AuthMiddleware) WithWebsocketRouthAuthentication(handler websocket.Handler) websocket.Handler {
	return func (conn *websocket.Conn) {

		r := conn.Request()

		authorization := r.URL.Query().Get("Authorization")

		err := verifyToken(authorization)

		if (errors.Is(err, ErrorMissingToken) || errors.Is(err, ErrorInvalidToken) || errors.Is(err, ErrorParsingToken)) {
			conn.Close()

			return
		}


		handler(conn)
	}
}


