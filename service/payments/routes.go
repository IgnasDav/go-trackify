package payments

import (
	"net/http"

	"github.com/gorilla/mux"
)

type Handler struct {
}

func NewHandler() *Handler {
	return &Handler{}
}

func (h *Handler) RegisterRoutes(r *mux.Router) {
	r.HandleFunc("/payments", h.getPayments).Methods(http.MethodGet)
}

func (h *Handler) getPayments(w http.ResponseWriter, r *http.Request) {

}
