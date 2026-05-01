package shared

import (
	"net/http"
	"strings"

	"github.com/go-chi/chi/v5"

	"ds2api/internal/config"
)

type ModelsHandler struct {
	Store ConfigReader
}

func (h *ModelsHandler) ListModels(w http.ResponseWriter, _ *http.Request) {
	var aliases map[string]string
	if h.Store != nil {
		aliases = h.Store.ConfigOnlyModelAliases()
	}
	WriteJSON(w, http.StatusOK, config.OpenAIModelsResponse(aliases))
}

func (h *ModelsHandler) GetModel(w http.ResponseWriter, r *http.Request) {
	modelID := strings.TrimSpace(chi.URLParam(r, "model_id"))
	model, ok := config.OpenAIModelByID(h.Store, modelID)
	if !ok {
		WriteOpenAIError(w, http.StatusNotFound, "Model not found.")
		return
	}
	WriteJSON(w, http.StatusOK, model)
}
