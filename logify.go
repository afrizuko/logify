package main

import (
	"log"
	"net/http"
	"os"

	"github.com/afrizuko/logify/handler"
	"github.com/afrizuko/logify/model"
	"github.com/xujiajun/gorouter"
)

func init() {

}

func main() {
	mux := gorouter.New()
	group := mux.Group("/api/v1")

	userService := model.NewUserMockImpl(5)
	handler.NewUserHandler(userService).Routes(group)

	port := os.Getenv("PORT")
	if port == "" {
		port = "4000"
	}
	log.Fatal(http.ListenAndServe(":"+port, mux))
}
