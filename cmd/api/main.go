package main

import (
	"emailn/internal/contract"
	"emailn/internal/domain/campaing"
	"emailn/internal/infrastructure/database"
	internalerros "emailn/internal/internalErros"
	"errors"
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

	service := campaing.Service{
		Repository: &database.CampaingRepository{},
	}
	r.Post("/campaings", func(w http.ResponseWriter, r *http.Request) {
		var request contract.NewCampaing
		err := render.DecodeJSON(r.Body, &request)
		if err != nil {
			http.Error(w, "Invalid request payload", http.StatusBadRequest)
			return
		}

		id, err := service.Create(request)
		if err != nil {
			if errors.Is(err, internalerros.ErrInternal) {
				render.Status(r, http.StatusInternalServerError)
				render.JSON(w, r, map[string]string{"error": err.Error()})
			} else {
				render.Status(r, http.StatusBadRequest)
				render.JSON(w, r, map[string]string{"error": err.Error()})
			}
			return
		}

		render.Status(r, http.StatusCreated)
		render.JSON(w, r, map[string]string{"id": id})
	})

	http.ListenAndServe(":3000", r)
}
