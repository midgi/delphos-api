package handlers

import (
	"encoding/json"
	"net/http"

	"github.com/pivotal-golang/lager"
)

type InfoHandler struct {
	logger lager.Logger
}

func NewInfoHandler(logger lager.Logger) *InfoHandler {
	return &InfoHandler{logger: logger}
}

func (h *InfoHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	logger := h.logger.Session("info-handler")

	infoResponse := map[string]interface{}{
		"version": "0.0.1",
	}

	encoder := json.NewEncoder(w)
	err := encoder.Encode(infoResponse)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		logger.Error("failed-marshaling-info", err)
		return
	}
}
