package handler

import (
	"encoding/json"
	"log"
	"net/http"
	"strconv"

	"github.com/afrizuko/logify/model"
	"github.com/xujiajun/gorouter"
)

type UserHandler struct {
	http.Handler
	repository model.UserRepository
}

func NewUserHandler(repo model.UserRepository) *UserHandler {
	handler := new(UserHandler)
	handler.repository = repo
	mux := gorouter.New()
	handler.Routes(mux)
	handler.Handler = mux
	return handler
}

func (h *UserHandler) Routes(mux *gorouter.Router) {
	mux.GET("/users", h.GetUsers)
	mux.GET("/users/:id", h.GetUser)
	mux.POST("/users", h.CreateUser)
	mux.PUT("/users/:id", h.ModifyUser)
	mux.DELETE("/users/:id", h.DeleteUser)
	h.Handler = mux
}

func (h *UserHandler) DeleteUser(w http.ResponseWriter, r *http.Request) {
	id, _ := strconv.ParseUint(gorouter.GetParam(r, "id"), 10, 64)

	if err := h.repository.DeleteUser(uint(id)); err != nil {
		renderJSON(404, w, err)
		return
	}
	renderJSON(202, w, nil)
}

func (h *UserHandler) ModifyUser(w http.ResponseWriter, r *http.Request) {

	id, _ := strconv.ParseUint(gorouter.GetParam(r, "id"), 10, 64)
	var user model.User
	json.NewDecoder(r.Body).Decode(&user)
	user.ID = uint(id)

	h.repository.UpdateUser(&user)
	renderJSON(202, w, user)
}

func (h *UserHandler) CreateUser(w http.ResponseWriter, r *http.Request) {

	var user model.User
	json.NewDecoder(r.Body).Decode(&user)
	h.repository.SaveUser(&user)
	renderJSON(200, w, user)
}

func (h *UserHandler) GetUsers(w http.ResponseWriter, r *http.Request) {

	users, _ := h.repository.GetUsers(0, 0)
	renderJSON(200, w, users)
}

func (h *UserHandler) GetUser(w http.ResponseWriter, r *http.Request) {

	id, _ := strconv.ParseUint(gorouter.GetParam(r, "id"), 10, 64)
	user, err := h.repository.GetUser(uint(id))
	if err != nil {
		renderJSON(404, w, map[string]string{"error": err.Error()})
		return
	}
	renderJSON(200, w, user)
}

func renderJSON(statusCode int, w http.ResponseWriter, data interface{}) {
	w.WriteHeader(statusCode)
	if data == nil {
		return
	}
	if err := json.NewEncoder(w).Encode(data); err != nil {
		log.Println(err)
	}
}
