package v1

import (
	"github.com/AZRV17/goWEB/internal/service"
	"github.com/go-chi/chi/v5"
	"github.com/go-chi/render"
	"github.com/go-playground/validator/v10"
	"log"
	"net/http"
	"strconv"
)

type createAccountInput struct {
	Owner    string `json:"owner" validate:"required"`
	Balance  int64  `json:"balance" validate:"required,min=0"`
	Currency string `json:"currency" validate:"required,oneof=USD EUR RUB BYN CNY"`
}

func (h *Handler) initAccountRoutes(r chi.Router) {
	r.Route("/account", func(r chi.Router) {
		r.Post("/", h.createAccount)
		r.Get("/{id}", h.getAccount)
		r.Put("/{id}", h.updateAccount)
		r.Delete("/{id}", h.deleteAccount)
		r.Get("/", h.getAccounts)
	})
}

func (h *Handler) getAccount(w http.ResponseWriter, r *http.Request) {
	idParam := chi.URLParam(r, "id")

	id, err := strconv.Atoi(idParam)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	acc, err := h.service.AccountService.GetAccount(int64(id))
	if err != nil {
		if err.Error() == "record not found" {
			w.WriteHeader(http.StatusNotFound)
			render.JSON(w, r, map[string]interface{}{
				"message": "account with id " + idParam + " not found",
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
	render.JSON(w, r, acc)
}

func (h *Handler) createAccount(w http.ResponseWriter, r *http.Request) {
	inp := &createAccountInput{}
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

	if err := h.service.AccountService.CreateAccount(service.CreateAccountInput{
		Owner:    inp.Owner,
		Balance:  inp.Balance,
		Currency: inp.Currency,
	}); err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		render.JSON(w, r, map[string]interface{}{
			"message": err.Error(),
		})
		return
	}

	w.WriteHeader(http.StatusOK)
	render.JSON(w, r, map[string]interface{}{
		"message": "account created successfully",
	})
}

func (h *Handler) updateAccount(w http.ResponseWriter, r *http.Request) {
	idParam := chi.URLParam(r, "id")

	id, err := strconv.Atoi(idParam)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	inp := &createAccountInput{}

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

	if err := h.service.AccountService.UpdateAccount(service.UpdateAccountInput{
		ID:       int64(id),
		Owner:    inp.Owner,
		Balance:  inp.Balance,
		Currency: inp.Currency,
	}); err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		render.JSON(w, r, map[string]interface{}{
			"message": err.Error(),
		})
		return
	}

	w.WriteHeader(http.StatusOK)
	render.JSON(w, r, map[string]interface{}{
		"message": "account updated successfully",
		"id":      id,
	})
}

func (h *Handler) deleteAccount(w http.ResponseWriter, r *http.Request) {
	idParam := chi.URLParam(r, "id")

	id, err := strconv.Atoi(idParam)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		render.JSON(w, r, map[string]interface{}{
			"message": err.Error(),
		})
		return
	}

	log.Println(id)

	if err := h.service.AccountService.DeleteAccount(int64(id)); err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		render.JSON(w, r, map[string]interface{}{
			"message": err.Error(),
		})
		return
	}

	w.WriteHeader(http.StatusOK)
	render.JSON(w, r, map[string]interface{}{
		"message": "account deleted successfully",
	})
}

func (h *Handler) getAccounts(w http.ResponseWriter, r *http.Request) {
	accs, err := h.service.AccountService.GetAllAccounts()
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		render.JSON(w, r, map[string]interface{}{
			"message": err.Error(),
		})
		return
	}

	w.WriteHeader(http.StatusOK)
	render.JSON(w, r, accs)
}
