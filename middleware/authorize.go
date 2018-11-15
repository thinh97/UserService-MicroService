package middleware

import (
	"context"
	"errors"
	"log"
	"net/http"
	"strings"

	"../common"
)

func Authorize(w http.ResponseWriter, r *http.Request, next http.HandlerFunc) {
	tokenString, err := tokenFromAuthHeader(r)
	if err != nil {
		log.Println("Get token error: ", err)
		common.RespondWithError(w, http.StatusUnauthorized, "Invalid Token")
		return
	}
	valid, appClaims := common.ValidateToken(tokenString)
	if valid {
		ctx := context.WithValue(r.Context(), "UserName", appClaims.UserName)
		next(w, r.WithContext(ctx))
	} else {
		common.RespondWithError(w, http.StatusUnauthorized, "Token Expired")
	}
}

func tokenFromAuthHeader(r *http.Request) (string, error) {
	ah := r.Header.Get("Authorization")
	if ah != "" {
		if len(ah) > 6 && strings.ToUpper(ah[0:6]) == "BEARER" {
			return ah[7:], nil
		}
	}
	return "", errors.New("No token in the HTTP request")
}
