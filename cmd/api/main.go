package main

import (
	"emailn/internal/contract"
	"emailn/internal/domain/campaing"
	"net/http"

	"github.com/go-chi/chi/middleware"
	"github.com/go-chi/chi/v5"
	"github.com/go-chi/render"
)

func main() {
	r := chi.NewRouter()

	r.Use(middleware.RequestID)
	r.Use(middleware.RealIP)
	r.Use(middleware.Logger)
	r.Use(middleware.Recoverer)

	service := campaing.Service{}
	r.Post("/campaings", func(w http.ResponseWriter, r *http.Request) {
		var campaing contract.NewCampaing
		render.DecodeJSON(r.Body, &campaing)
		id, err := service.Create(campaing)
		if err != nil {
			http.Error(w, err.Error(), 400)
			return
		}
		render.Status(r, 201)
		render.JSON(w, r, map[string]string{"id": id})
	})

	http.ListenAndServe(":3000", r)
}
