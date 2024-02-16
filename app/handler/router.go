package handler

import (
	"database/sql"
	"log"
	"net/http"
	"time"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	"github.com/go-chi/cors"

	_userHttpDelivery "go_docker/user/delivery/http"
	_userRepository "go_docker/user/repository/mysql"
	_userUsecase "go_docker/user/usecase"

	_healthHttpDelivery "go_docker/health/delivery/http"

)

func NewRouter() http.Handler {
	r := chi.NewRouter()

	// A good base middleware stack
	r.Use(middleware.RequestID)
	r.Use(middleware.RealIP)
	r.Use(middleware.Logger)
	r.Use(middleware.Recoverer)
	r.Use(newCORS().Handler)

	// Set a timeout value on the request context (ctx), that will signal
	// through ctx.Done() that the request has timed out and further
	// processing should be stopped.
	r.Use(middleware.Timeout(60 * time.Second))

	dbConn, err := sql.Open("mysql", "root:root@tcp(mysql:3306)/userdb1")
	if err != nil {
		log.Fatal(err)
	}

	err = dbConn.Ping()
	if err != nil {
		log.Fatal(err)
	}

	defer func() {
		err := dbConn.Close()
		if err != nil {
			log.Fatal(err)
		}
	}()

	userRepo := _userRepository.NewMysqlUserRepository(dbConn)
	userUsecase := _userUsecase.NewUserUsecase(userRepo)
	userHandler := _userHttpDelivery.NewUserHandler(*userUsecase)
	r.Mount("/user", userHandler.Routes())

	// health
	r.Mount("/", _healthHttpDelivery.NewHealthHandler().Routes())

	return r
}

func newCORS() *cors.Cors {
	return cors.New(cors.Options{
		AllowedOrigins: []string{"*"},
		AllowedHeaders: []string{"*"},
		AllowedMethods: []string{
			http.MethodGet,
			http.MethodHead,
			http.MethodPut,
			http.MethodPatch,
			http.MethodPost,
			http.MethodDelete,
			http.MethodOptions,
		},
	})
}
