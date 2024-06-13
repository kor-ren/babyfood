package auth

import (
	"crypto"
	"encoding/base64"
	"net/http"
	"time"
)

const (
	cookieName = "c184fe111a304773"
)

func Middleware(tokenHash string, secureCookie string, next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		cookie, err := r.Cookie(cookieName)

		if err != nil {
			w.WriteHeader(http.StatusUnauthorized)
			return
		}

		if cookie.Valid() == nil && cookie.Value == tokenHash {
			setCookie(w, tokenHash, secureCookie)
			next.ServeHTTP(w, r)
			return
		}

		w.WriteHeader(http.StatusForbidden)
	})
}

func LoginHandler(token string, tokenHash string, secure string) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		queryParam := r.URL.Query()["token"]

		if len(queryParam) != 1 {
			w.WriteHeader(http.StatusUnauthorized)
			return
		}

		if queryParam[0] != token {
			w.WriteHeader(http.StatusUnauthorized)
			return
		}

		setCookie(w, tokenHash, secure)
		http.Redirect(w, r, "/", http.StatusFound)
	}
}

func setCookie(w http.ResponseWriter, tokenHash string, secure string) {
	http.SetCookie(w, &http.Cookie{
		Name:     cookieName,
		HttpOnly: true,
		Secure:   secure == "true",
		Path:     "/",
		Value:    tokenHash,
		Expires:  time.Now().Add(365 * 24 * time.Hour),
	})
}

func GetTokenHash(token string) string {
	sha := crypto.SHA512.New()
	sha.Write([]byte(token))
	return base64.URLEncoding.EncodeToString(sha.Sum(nil))
}
