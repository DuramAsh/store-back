package handler

import (
	"store-back/docs"
	"store-back/internal/config"
	"store-back/internal/handler/http"
	"store-back/internal/service/general"
	"store-back/pkg/server/router"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	httpSwagger "github.com/swaggo/http-swagger/v2"
)

type Dependencies struct {
	Configs        config.Configs
	GeneralService *general.Service
}

// Configuration is an alias for a function that will take in a pointer to a Handler and modify it
type Configuration func(h *Handler) error

// Handler is an implementation of the Handler
type Handler struct {
	dependencies Dependencies

	HTTP *chi.Mux
}

// New takes a variable amount of Configuration functions and returns a new Handler
// Each Configuration will be called in the order they are passed in
func New(d Dependencies, configs ...Configuration) (h *Handler, err error) {
	// Insert the handler
	h = &Handler{
		dependencies: d,
	}

	// Apply all Configurations passed in
	for _, cfg := range configs {
		// Pass the service into the configuration function
		if err = cfg(h); err != nil {
			return
		}
	}

	return
}

// WithHTTPHandler applies a http handler to the Handler
func WithHTTPHandler() Configuration {
	return func(h *Handler) (err error) {
		// Insert the http handler, if we needed parameters, such as connection strings they could be inputted here
		h.HTTP = router.New()

		h.HTTP.Use(middleware.Timeout(h.dependencies.Configs.APP.Timeout))

		// Init swagger handler
		docs.SwaggerInfo.BasePath = h.dependencies.Configs.APP.Path
		h.HTTP.Get("/swagger/*", httpSwagger.WrapHandler)

		// Init service handlers
		generalHandler := http.NewGeneralHandler(h.dependencies.GeneralService)

		if err != nil {
			return
		}

		h.HTTP.Route("/", func(r chi.Router) {
			r.Mount("/api", generalHandler.Routes())
		})

		return
	}
}
