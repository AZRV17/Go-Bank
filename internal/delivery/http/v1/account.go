package v1

import (
	"net/http"
	"strconv"

	"github.com/AZRV17/goWEB/internal/service"
	"github.com/go-chi/chi/v5"
	"github.com/go-chi/render"
)

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
	var body []byte
	if _, err := r.Body.Read(body); err != nil {
		w.WriteHeader(http.StatusBadRequest)
		render.JSON(w, r, map[string]interface{}{
			"message": err.Error(),
		})
	}

	var acc service.CreateAccountInput
	if err := render.DecodeJSON(r.Body, &acc); err != nil {
		w.WriteHeader(http.StatusBadRequest)
		render.JSON(w, r, map[string]interface{}{
			"message": err.Error(),
		})
	}

	if err := h.service.AccountService.CreateAccount(acc); err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		render.JSON(w, r, map[string]interface{}{
			"message": err.Error(),
		})
	}

	w.WriteHeader(http.StatusOK)
	render.JSON(w, r, map[string]interface{}{
		"message": "account created successfully",
	})
}

func (h *Handler) updateAccount(w http.ResponseWriter, r *http.Request) {
	var body []byte
	if _, err := r.Body.Read(body); err != nil {
		w.WriteHeader(http.StatusBadRequest)
		render.JSON(w, r, map[string]interface{}{
			"message": err.Error(),
		})
		return
	}

	idParam := chi.URLParam(r, "id")

	id, err := strconv.Atoi(idParam)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	var acc service.UpdateAccountInput
	acc.ID = int64(id)

	if err := render.DecodeJSON(r.Body, &acc); err != nil {
		w.WriteHeader(http.StatusBadRequest)
		render.JSON(w, r, map[string]interface{}{
			"message": err.Error(),
		})
		return
	}

	if err := h.service.AccountService.UpdateAccount(acc); err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		render.JSON(w, r, map[string]interface{}{
			"message": err.Error(),
		})
		return
	}

	w.WriteHeader(http.StatusOK)
	render.JSON(w, r, map[string]interface{}{
		"message": "account updated successfully",
		"id":      acc.ID,
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
