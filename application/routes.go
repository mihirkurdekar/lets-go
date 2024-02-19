package application

import (
	"net/http"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	"github.com/mihirkurdekar/lets-go/handler"
)

func (a *App) loadRoutes() {
	router := chi.NewRouter()

	router.Use(middleware.Logger)

	router.Get("/", func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusOK)
	})

	router.Route("/orders", a.loadOrderRoutes)

	a.router = router
}

func (a *App) loadOrderRoutes(router chi.Router) {
	orderHander := &handler.Order{}

	router.Post("/", orderHander.Create)
	router.Get("/", orderHander.List)
	router.Get("/{id}", orderHander.GetByID)
	router.Put("/{id}", orderHander.UpdateByID)
	router.Delete("/{id}", orderHander.DeleteByID)
}
