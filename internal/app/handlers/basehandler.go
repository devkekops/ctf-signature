package handlers

import (
	"encoding/json"
	"log"
	"net/http"
	"os"
	"path/filepath"
	"strconv"

	"github.com/go-chi/chi/middleware"
	"github.com/go-chi/chi/v5"
	"github.com/go-chi/cors"

	"github.com/devkekops/ctf-signature/internal/app/storage"
)

type BaseHandler struct {
	*chi.Mux
	fs          http.Handler
	paymentRepo storage.PaymentRepository
}

func NewBaseHandler(paymentRepo storage.PaymentRepository) *BaseHandler {
	cwd, _ := os.Getwd()
	root := filepath.Join(cwd, "/static")
	fs := http.FileServer(http.Dir(root))

	bh := &BaseHandler{
		Mux:         chi.NewMux(),
		fs:          fs,
		paymentRepo: paymentRepo,
	}

	bh.Use(middleware.Logger)

	bh.Use(cors.Handler(cors.Options{
		AllowedOrigins:   []string{"https://*", "http://*"},
		AllowedMethods:   []string{"GET", "POST", "PUT", "DELETE", "OPTIONS"},
		AllowedHeaders:   []string{"Accept", "Authorization", "Content-Type", "X-CSRF-Token"},
		ExposedHeaders:   []string{"Link"},
		AllowCredentials: false,
		MaxAge:           300, // Maximum value not ignored by any of major browsers
	}))

	bh.Handle("/*", fs)
	bh.Get("/getPayments", bh.getPayments())
	bh.Get("/getPayment", bh.getPayment())

	return bh
}

func (bh *BaseHandler) getPayments() http.HandlerFunc {
	return func(w http.ResponseWriter, req *http.Request) {
		idString := req.URL.Query().Get("offset")
		id, err := strconv.ParseInt(idString, 10, 64)
		if err != nil {
			w.WriteHeader(http.StatusBadRequest)
			log.Println(err)
			return
		}

		payments, err := bh.paymentRepo.GetPayments(id)
		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			log.Println(err)
			return
		}

		buf, err := json.Marshal(payments)
		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			log.Println(err)
			return
		}

		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)
		_, err = w.Write(buf)
		if err != nil {
			log.Println(err)
		}
	}
}

func (bh *BaseHandler) getPayment() http.HandlerFunc {
	return func(w http.ResponseWriter, req *http.Request) {
		id := req.URL.Query().Get("id")
		payments, err := bh.paymentRepo.GetPayment(id)
		if err != nil {
			w.WriteHeader(http.StatusNotFound)
			log.Println(err)
			return
		}

		buf, err := json.Marshal(payments)
		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			log.Println(err)
			return
		}

		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)
		_, err = w.Write(buf)
		if err != nil {
			log.Println(err)
		}
	}
}
