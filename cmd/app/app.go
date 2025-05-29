package app

import (
	"backend-habitus-app/cmd/router"
	"database/sql"
	"net/http"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	"github.com/go-sql-driver/mysql"
)

type ConfigServerMySQLChi struct {
	Db            *mysql.Config
	ServerAddress string
}

type ServerMySQLChi struct {
	cfgDb   *mysql.Config
	cfgAddr string
	db      *sql.DB
	router  chi.Router
}

func NewConfigServerMySQLChi(cfg *ConfigServerMySQLChi) *ServerMySQLChi {
	defaultCfg := &ConfigServerMySQLChi{
		Db:            nil,
		ServerAddress: ":8080",
	}

	if cfg != nil {
		if cfg.Db != nil {
			defaultCfg.Db = cfg.Db
		}
		if cfg.ServerAddress != "" {
			defaultCfg.ServerAddress = cfg.ServerAddress
		}
	}

	return &ServerMySQLChi{
		cfgDb:   defaultCfg.Db,
		cfgAddr: defaultCfg.ServerAddress,
	}

}

func (s *ServerMySQLChi) Setup() (err error) {

	s.db, err = sql.Open("mysql", s.cfgDb.FormatDSN())
	if err != nil {
		return err
	}

	err = s.db.Ping()
	if err != nil {
		return err
	}

	rt := chi.NewRouter()

	rt.Use(middleware.Logger)
	rt.Use(middleware.Recoverer)

	router.RegisterHabitRoutes(rt, s.db)
	router.RegisterTaskRoutes(rt, s.db)
	router.RegisterHabitLogRoutes(rt, s.db)

	s.router = rt

	return
}

func (s *ServerMySQLChi) Start() (err error) {
	err = http.ListenAndServe(s.cfgAddr, s.router)
	return
}