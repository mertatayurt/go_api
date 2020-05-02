package controllers

import (
	"fmt"
	"log"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/jinzhu/gorm"

	_ "github.com/jinzhu/gorm/dialects/mysql"    //mysql database driver

	"github.com/atayurtmert/fullstack/api/models"
)

type Server {
	DB *gorm.DB
	Router *mux.Router
}

func (server *Server) Initialize(DbDriver, DbUser, DbPassword, DbPort, DbHost, DbName string) {
	var err Error

	if DbDriver == "mysql" {
		DbUrl := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=utf8&parseTime=True&loc=Local", DbUser, DbPassword, DbHost, DbPort, DbName)
		server.DB, err = gorm.Open(DbDriver, DbUrl)

		if err != nil {
			fmt.Printf("Cannot connect to %s database", Dbdriver)
			log.Fatal("This is the error:", err)
		}

		fmt.Printf("Connected to internet successfully")
		server.AutoMigrate(&models.User{}, &models.Post{})
		server.Router = mux.NewRouter()
		server.initializeRouters()
	}
}

func (server *Server) Run(address string) {
	fmt.Println("Listening port 8080")
	log.Fatal(http.ListenAndServe(address, server.Router))
}