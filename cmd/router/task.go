package router

import (
	"backend-habitus-app/internal/handler"
	"backend-habitus-app/internal/repository"
	"backend-habitus-app/internal/service"
	"database/sql"

	"github.com/go-chi/chi/v5"
)

func RegisterTaskRoutes(r chi.Router, db *sql.DB) {

	rp := repository.NewTasktRepository(db)

	sv := service.NewTaskService(rp)

	hd := handler.NewTaskHandler(sv)

	r.Route("/tasks", func(r chi.Router) {
		r.Get("/", hd.GetAll)
		r.Post("/", hd.Create)
		r.Get("/{id}", hd.GetByID)
		r.Put("/{id}", hd.Update)
		r.Delete("/{id}", hd.Delete)
	})

}