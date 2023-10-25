package v1

import (
	"github.com/AZRV17/goWEB/internal/service"
	"github.com/go-chi/chi/v5"
	"github.com/go-chi/render"
	"github.com/go-playground/validator/v10"
	"net/http"
	"strconv"
)

type transferInput struct {
	FromAccountID int64 `json:"from_account_id" validate:"required,min=1"`
	ToAccountID   int64 `json:"to_account_id" validate:"required,min=1"`
	Amount        int64 `json:"amount" validate:"required,gt=0"`
}

func (h *Handler) initTransferRoutes(r chi.Router) {
	r.Route("/transfer", func(r chi.Router) {
		r.Post("/", h.createTransfer)
		r.Get("/{id}", h.getTransfer)
		r.Get("/", h.getTransfers)
	})
}

func (h *Handler) getTransfer(w http.ResponseWriter, r *http.Request) {
	idParam := chi.URLParam(r, "id")

	id, err := strconv.Atoi(idParam)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		render.JSON(w, r, map[string]interface{}{
			"message": err.Error(),
		})
		return
	}

	transfer, err := h.service.TransferService.GetTransfer(int64(id))
	if err != nil {
		if err.Error() == "record not found" {
			w.WriteHeader(http.StatusNotFound)
			render.JSON(w, r, map[string]interface{}{
				"message": "transfer with id " + idParam + " not found",
			})
			return
		}

		w.WriteHeader(http.StatusInternalServerError)
		render.JSON(w, r, map[string]interface{}{
			"message": err.Error(),
		})
		return
	}

	w.WriteHeader(http.StatusOK)
	render.JSON(w, r, transfer)
}

func (h *Handler) createTransfer(w http.ResponseWriter, r *http.Request) {
	inp := &transferInput{}
	if err := render.DecodeJSON(r.Body, inp); err != nil {
		w.WriteHeader(http.StatusBadRequest)
		render.JSON(w, r, map[string]interface{}{
			"message": err.Error(),
		})
		return
	}

	validate := validator.New(validator.WithRequiredStructEnabled())
	if err := validate.Struct(inp); err != nil {
		w.WriteHeader(http.StatusBadRequest)
		render.JSON(w, r, map[string]interface{}{
			"message": err.Error(),
		})
		return
	}

	if err := h.service.TransferService.CreateTransfer(service.CreateTransferInput{
		FromAccountID: inp.FromAccountID,
		ToAccountID:   inp.ToAccountID,
		Amount:        inp.Amount,
	}); err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		render.JSON(w, r, map[string]interface{}{
			"message": err.Error(),
		})
		return
	}

	w.WriteHeader(http.StatusOK)
	render.JSON(w, r, map[string]interface{}{
		"message": "transfer created successfully",
	})
}

func (h *Handler) getTransfers(w http.ResponseWriter, r *http.Request) {
	transfers, err := h.service.TransferService.GetAllTransfers()
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		render.JSON(w, r, map[string]interface{}{
			"message": err.Error(),
		})
		return
	}

	w.WriteHeader(http.StatusOK)
	render.JSON(w, r, transfers)
}
