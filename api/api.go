package api

import (
	"net/http"

	"github.com/micahli/notable-take-home/api/response"
	"github.com/micahli/notable-take-home/db"

	"github.com/gorilla/mux"
)

// Config represents the API configuration
type Config struct {
	Domain string `yaml:"domain"`
}

// API represents the structure of the API
type API struct {
	Router *mux.Router

	config *Config
	db     *db.DB
}

// New returns the api settings
func New(config *Config, db *db.DB, router *mux.Router) (*API, error) {
	api := &API{
		config: config,
		db:     db,
		Router: router,
	}

	// Endpoint for browser preflight requests
	api.Router.Methods("OPTIONS").HandlerFunc(api.corsMiddleware(api.preflightHandler))

	// Endpoint for healtcheck
	api.Router.HandleFunc("/api/v1/health", api.corsMiddleware(api.logMiddleware(api.healthHandler))).Methods("GET")

	// Doctor related api endpoints
	api.Router.HandleFunc("/api/v1/doctor", api.corsMiddleware(api.logMiddleware(api.getDoctorListHandler))).Methods("GET")
	api.Router.HandleFunc("/api/v1/doctor/{id}", api.corsMiddleware(api.logMiddleware(api.getAppointmentListHandler))).Methods("GET")
	api.Router.HandleFunc("/api/v1/doctor/appointment", api.corsMiddleware(api.logMiddleware(api.addAppointmentHandler))).Methods("POST")
	api.Router.HandleFunc("/api/v1/doctor/appointment", api.corsMiddleware(api.logMiddleware(api.cancelAppointmentHandler))).Methods("DELETE")

	return api, nil
}

func (a *API) healthHandler(w http.ResponseWriter, r *http.Request) {
	response.Write(w, r, struct {
		Status string `json:"status"`
	}{
		"ok",
	})

	return
}

func (a *API) preflightHandler(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusOK)
	return
}
