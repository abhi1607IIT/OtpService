package main

import (
	"OtpService/controller"
	repo2 "OtpService/dal/repo"
	service2 "OtpService/service"
	"github.com/gorilla/mux"
	"github.com/jinzhu/gorm"
	_ "github.com/lib/pq"
	"log"
	"net/http"
	"strconv"
)

// docker run --name postgres-db -e POSTGRES_USER=sqe -e POSTGRES_PASSWORD=sqesecretpassword -e POSTGRES_DB=sqedb -p 5432:5432 -d postgres
const (
	host     = "localhost"
	port     = 5432
	user     = "sqe"
	password = "sqesecretpassword"
	dbname   = "sqedb"
)

func main() {
	connectionString :=
		"host=" + host +
			" port=" + strconv.Itoa(port) +
			" user=" + user +
			" password=" + password +
			" dbname=" + dbname +
			" sslmode=disable"

	var err error
	var db *gorm.DB
	db, err = gorm.Open("postgres", connectionString)
	if err != nil {
		log.Fatal(err)
	}

	defer db.Close()

	router := mux.NewRouter()

	repo := repo2.NewOtpRepo(db)
	service := service2.NewOtpService(repo)

	controller.NewUserController(service, router)

	log.Fatal(http.ListenAndServe(":8080", router))

}
