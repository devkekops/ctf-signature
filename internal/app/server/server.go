package server

import (
	"net/http"

	"github.com/devkekops/ctf-signature/internal/app/config"
	"github.com/devkekops/ctf-signature/internal/app/handlers"
	"github.com/devkekops/ctf-signature/internal/app/storage"
)

func Serve(cfg *config.Config) error {
	var paymentRepo storage.PaymentRepository
	paymentRepo = storage.NewPaymentRepo()

	var baseHandler = handlers.NewBaseHandler(paymentRepo)

	server := &http.Server{
		Addr:    cfg.ServerAddress,
		Handler: baseHandler,
	}

	return server.ListenAndServe()
}