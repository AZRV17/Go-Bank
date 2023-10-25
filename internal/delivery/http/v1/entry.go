package v1

import (
	"github.com/go-chi/chi/v5"
	"github.com/go-chi/render"
	"net/http"
	"strconv"
)

func (h *Handler) initEntryRoutes(r chi.Router) {
	r.Route("/entry", func(r chi.Router) {
		r.Get("/{id}", h.getEntry)
		r.Get("/", h.getEntries)
	})
}

func (h *Handler) getEntry(w http.ResponseWriter, r *http.Request) {
	idParam := chi.URLParam(r, "id")

	id, err := strconv.Atoi(idParam)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		render.JSON(w, r, map[string]interface{}{
			"message": err.Error(),
		})
		return
	}

	entry, err := h.service.EntryService.GetEntry(int64(id))
	if err != nil {
		if err.Error() == "record not found" {
			w.WriteHeader(http.StatusNotFound)
			render.JSON(w, r, map[string]interface{}{
				"message": "entry with id " + idParam + " not found",
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
	render.JSON(w, r, entry)
}

func (h *Handler) getEntries(w http.ResponseWriter, r *http.Request) {
	entries, err := h.service.EntryService.GetAllEntries()
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		render.JSON(w, r, map[string]interface{}{
			"message": err.Error(),
		})
		return
	}

	w.WriteHeader(http.StatusOK)
	render.JSON(w, r, entries)
}
