package http

import (
	"context"
	"net/http"
	"strconv"

	"go_docker/user/usecase"

	"github.com/go-chi/chi/v5"
)

type ResponseError struct {
	Message string `json:"message"`
}
type UserHandler struct {
	userUsecase usecase.UserUsecase
}

func NewUserHandler(u usecase.UserUsecase) *UserHandler {
	handler := &UserHandler{
		userUsecase: u,
	}
	return handler
}

func (h *UserHandler) Routes() http.Handler {
	r := chi.NewRouter()
	r.Get("/{id}", h.GetById)
	return r
}

func (h *UserHandler) test(w http.ResponseWriter, r *http.Request) {
	id, err := strconv.ParseInt(chi.URLParam(r, "id"), 10, 64)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte(err.Error()))
		return
	}
	w.Write([]byte(strconv.FormatInt(id, 10)))

}

func (h *UserHandler) GetById(w http.ResponseWriter, r *http.Request) {
	id, err := strconv.ParseInt(chi.URLParam(r, "id"), 10, 64)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte(err.Error()))
		return
	}
	user, err := h.userUsecase.GetById(context.Background(), id)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte(err.Error()))
		return
	}

	w.Write([]byte(user.Name))
}
