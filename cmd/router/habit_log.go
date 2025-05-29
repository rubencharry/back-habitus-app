package router

import (
	"backend-habitus-app/internal/handler"
	"backend-habitus-app/internal/repository"
	"backend-habitus-app/internal/service"
	"database/sql"

	"github.com/go-chi/chi/v5"
)

func RegisterHabitLogRoutes(r chi.Router, db *sql.DB) {

	rp := repository.NewHabitLogtRepository(db)

	sv := service.NewHabitLogService(rp)

	hd := handler.NewHabitLogHandler(sv)

	r.Route("/habits_logs", func(r chi.Router) {
		r.Get("/", hd.GetAll)
		r.Post("/", hd.Create)
	})

}