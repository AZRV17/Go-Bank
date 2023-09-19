package handler

import (
	"encoding/json"
	"log"
	"net/http"

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
		m, err := h.service.AccountService.GetAccount(1)
		if err != nil {
			log.Fatal(err)
		}

		json.NewEncoder(w).Encode(m)
	})
}
