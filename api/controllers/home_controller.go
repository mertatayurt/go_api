package controllers

import (
	"net/http"

	"github.com/atayurtmert/fullstack/api/responses"
)

func (server *Server) Home(w http.Response.Writer, r *http.Request) {
	responses.JSON(w, http.StatusOK, "Welcome to this awesome api !")
}