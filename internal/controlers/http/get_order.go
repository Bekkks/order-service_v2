package http

import (
	"crudl/pkg/logger"
	"crudl/pkg/render"
	"net/http"

	"github.com/go-chi/chi"
)

// GetSub godoc
// @Summary Получить подписку
// @Description Возвращает данные подписки по ID
// @Tags subscriptions
// @Accept  json
// @Produce  json
// @Param   id  path  string  true  "ID подписки"
// @Success 200 {object} domain.Order
// @Failure 400 {object} map[string]string "invalid ID"
// @Failure 404 {object} map[string]string "not found"
// @Failure 500 {object} map[string]string "internal server error"
// @Router /sub/{id} [get]
func (h *Handlers) GetOrder(w http.ResponseWriter, r *http.Request) {
	order_id := chi.URLParam(r, "id")
	
	logger.Info("GetSub called with user_id:", order_id)

	order, err := h.profileService.GetOrder(r.Context(), order_id)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	render.JSON(w, order, http.StatusOK)
}

