package router

import (
	"backend-habitus-app/internal/handler"
	"backend-habitus-app/internal/repository"
	"backend-habitus-app/internal/service"
	"database/sql"

	"github.com/go-chi/chi/v5"
)

func RegisterHahitRoutes(r chi.Router, db *sql.DB) {

	rp := repository.NewHabitRepository(db)

	sv := service.NewHabitService(rp)

	hd := handler.NewHabitHandler(sv)

	r.Route("/habits", func(r chi.Router) {
		r.Get("/", hd.GetAll)
		r.Post("/", hd.Create)
		r.Get("/{id}", hd.GetByID)
		r.Put("/{id}", hd.Update)
		r.Delete("/{id}", hd.Delete)
	})

}