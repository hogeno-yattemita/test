package http

import(
	"net/http"
)

type HealthHandler struct {
}

func NewHealthHandler() *HealthHandler {
	handler := &HealthHandler{}
	return handler
}

func (h *HealthHandler) Routes() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "text/plain")
		_, err := w.Write([]byte("OK"))
		if err != nil {
			panic(err)
		}
	}
}