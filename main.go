package main

import (
	"database/sql"
	"log"
	"net/http"

	"users/handlers"
	"users/repositories"
	"users/services"

	_ "github.com/lib/pq"
)

func main() {
	db, err := sql.Open("postgres", "user=myuser dbname=mydb sslmode=disable")
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()

	repository := repositories.NewUserRepository(db)
	service := services.NewUserService(repository)
	handler := handlers.NewUserHandler(service)

	http.HandleFunc("/users", handler.CreateUser)
	http.HandleFunc("/users/all", handler.GetUsers)
	http.HandleFunc("/users/by_date_and_age", handler.GetUsersByDateAndAge)
	http.HandleFunc("/users/count", handler.GetUsersCount)

	log.Fatal(http.ListenAndServe(":8080", nil))
}
