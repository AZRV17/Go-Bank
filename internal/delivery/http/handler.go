package http

import (
	v1 "github.com/AZRV17/goWEB/internal/delivery/http/v1"
	"github.com/AZRV17/goWEB/internal/service"
	"github.com/go-chi/chi/v5"
)

type Handler struct {
	service service.Service
}

func NewHandler(service service.Service) *Handler {
	return &Handler{
		service: service,
	}
}

func (h *Handler) Init(r chi.Router) {
	v1 := v1.NewHandler(h.service)
	v1.Init(r)
}
