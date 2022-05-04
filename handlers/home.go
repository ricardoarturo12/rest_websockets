package handlers

import (
	"encoding/json"
	"net/http"

	"github.com/ricardoarturo12/rest_websockets/server"
)

type HomeReponse struct {
	Message string `json:"message"`
	Status  bool   `json:"status"`
}

func HomeHandler(s server.Server) http.HandlerFunc {
	// w da la respuesta, r request lo que envia el cliente
	return func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)
		json.NewEncoder(w).Encode(HomeReponse{
			Message: "welcome",
			Status:  true,
		})
	}
}
