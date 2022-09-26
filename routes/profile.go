package routes

import (
	"waybeens-app/handlers"
	"waybeens-app/pkg/mysql"
	"waybeens-app/repositories"

	"github.com/gorilla/mux"
)

func ProfileRoutes(r *mux.Router) {
	profileRepository := repositories.RepositoryProfile(mysql.DB)
	h := handlers.HandlerProfile(profileRepository)

	r.HandleFunc("/profile/{id}", h.GetProfile).Methods("GET")
}
