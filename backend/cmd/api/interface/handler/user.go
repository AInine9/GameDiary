package handler

import (
	"backend/cmd/api/usecase"
	"encoding/json"
	"github.com/julienschmidt/httprouter"
	"net/http"
)

type UserHandler interface {
	Index(http.ResponseWriter, *http.Request, httprouter.Params)
}

type userHandler struct {
	userUseCase usecase.UserUseCase
}

func NewUserHandler(uu usecase.UserUseCase) UserHandler {
	return &userHandler{userUseCase: uu}
}

func (uh userHandler) Index(w http.ResponseWriter, r *http.Request, pr httprouter.Params) {
	name := r.FormValue("name")

	user, err := uh.userUseCase.Search(name)
	if err != nil {
		http.Error(w, err.Error(), 500)
		return
	}

	if err = json.NewEncoder(w).Encode(user); err != nil {
		http.Error(w, err.Error(), 500)
		return
	}
}
