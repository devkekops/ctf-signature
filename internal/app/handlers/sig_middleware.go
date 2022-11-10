package handlers

import (
	"crypto/md5"
	"encoding/hex"
	"log"
	"net/http"
)

func checkSignature(payload string, signature string, secretKey string) bool {
	sigBytes := md5.Sum([]byte(payload + "secret_key=" + secretKey))
	sig := hex.EncodeToString(sigBytes[:])

	if sig == signature {
		return true
	} else {
		return false
	}
}

func sigHandle(secretKey string) (sh func(http.Handler) http.Handler) {
	sh = func(h http.Handler) http.Handler {
		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			values := r.URL.Query()
			payload := ""
			for k, v := range values {
				payload = payload + k + "=" + v[0]
			}
			signature := r.Header.Get("X-SIG-TOKEN")

			result := checkSignature(payload, signature, secretKey)

			if result {
				h.ServeHTTP(w, r)
			} else {
				w.WriteHeader(http.StatusBadRequest)
				log.Println("invalid signature")
				return
			}
		})
	}
	return
}
