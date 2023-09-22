package handler

import (
	"encoding/json"
	"log"
	"net/http"
	"time"

	"github.com/AZRV17/goWEB/internal/service"
)

type Handler struct {
	service service.Service
}

func NewHandler(service service.Service) *Handler {
	return &Handler{
		service: service,
	}
}

func (h *Handler) Init(mux *http.ServeMux) {
	mux.Handle("/v1", h.v1())
}

func (h *Handler) v1() http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")

		acc := service.CreateAccountInput{
			Owner:     "John Doe",
			Balance:   100,
			Currency:  "USD",
			CreatedAt: time.Now(),
		}

		err := h.service.AccountService.CreateAccount(acc)
		if err != nil {
			log.Fatal(err)
		}

		json.NewEncoder(w).Encode("status:success")
	})
}
