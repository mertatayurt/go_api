package controllers

import (
	"fmt"
	"net/http"

	"github.com/mertatayurt/go_api/api/responses"
)

func (server *Server) Home(w http.ResponseWriter, r *http.Request) {
	responses.JSON(w, http.StatusOK, "Welcome to this awesome api !")
}

func (server *Server) Deneme(w http.ResponseWriter, r *http.Request) {
	fmt.Println("deneme")
	responses.JSON(w, http.StatusOK, "Deneme")
}
