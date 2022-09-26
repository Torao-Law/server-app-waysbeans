package routes

import (
	"waybeens-app/handlers"
	"waybeens-app/pkg/mysql"
	"waybeens-app/repositories"

	"github.com/gorilla/mux"
)

func UserRoutes(r *mux.Router) {
	// create connection to db in repo/func connect
	userRepository := repositories.RepositoryUser(mysql.DB)

	// call handler user
	h := handlers.HandlerUser(userRepository)

	r.HandleFunc("/users", h.FindUsers).Methods("GET")
	r.HandleFunc("/user/{id}", h.GetUser).Methods("GET")
	r.HandleFunc("/user/{id}", h.UpdateUser).Methods("PATCH")
	r.HandleFunc("/user/{id}", h.DeleteUser).Methods("DELETE")
}
