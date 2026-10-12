package middleware

import (
	"crypto/hmac"
	"crypto/sha256"
	"encoding/base64"
	"fmt"
	"net/http"
	"strings"
)

func (m *Middlewares) AuthenticateJwt(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {

		header := r.Header.Get("Authorization")

		if header == "" {
			http.Error(w, "Unauthorized", http.StatusUnauthorized)
			return
		}

		headerArry := strings.Split(header, " ")

		if len(headerArry) != 2 {
			http.Error(w, "Unauthorized", http.StatusUnauthorized)
			return
		}

		accessToken := headerArry[1]

		fmt.Println("<<<AccessToken>>>", accessToken)

		tokenParts := strings.Split(accessToken, ".")

		if len(tokenParts) != 3 {
			http.Error(w, "Unauthorized", http.StatusUnauthorized)
			return
		}
		jwtHeader := tokenParts[0]
		jwtPayload := tokenParts[1]
		signature := tokenParts[2]

		message := jwtHeader + "." + jwtPayload

		byteArrSecret := []byte(m.cnf.JwtSecretkey)

		byteArrMessage := []byte(message)

		h := hmac.New(sha256.New, byteArrSecret)

		h.Write(byteArrMessage)

		hash := h.Sum(nil)

		newSignature := base64UrlEncode(hash)

		if newSignature != signature {
			http.Error(w, "Unauthorized", http.StatusUnauthorized)
			return
		}

		next.ServeHTTP(w, r)

	})
}
func base64UrlEncode(data []byte) string {
	return base64.URLEncoding.WithPadding(base64.NoPadding).EncodeToString(data)
}
