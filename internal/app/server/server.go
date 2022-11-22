package server

import (
	"net/http"

	"github.com/devkekops/ctf-signature/internal/app/config"
	"github.com/devkekops/ctf-signature/internal/app/handlers"
	"github.com/devkekops/ctf-signature/internal/app/storage"
)

func Serve(cfg *config.Config) error {
	var paymentRepo storage.PaymentRepository
	paymentRepo = storage.NewPaymentRepo(cfg.Flag)

	var baseHandler = handlers.NewBaseHandler(paymentRepo, cfg.SecretKey)

	server := &http.Server{
		Addr:    cfg.ServerAddress,
		Handler: baseHandler,
	}

	return server.ListenAndServe()
}
