package main

import (
	"log"
	"net/http"

	"github.com/afrizuko/logly/handler"
	"github.com/afrizuko/logly/model"
	"github.com/xujiajun/gorouter"
)

func main() {

	mux := gorouter.New()
	group := mux.Group("/api/v1")

	userService := model.NewUserMockImpl(5)
	handler.NewUserHandler(userService).Routes(group)

	log.Fatal(http.ListenAndServe(":8181", mux))
}
