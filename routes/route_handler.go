package routes

import (
	"github.com/gorilla/mux"
	"github.com/prometheus/client_golang/prometheus/promhttp"
	"net/http/pprof"
	"url-shortner/api"
)

const (
	HandlerTimeOutInMs = 1000
)

func UrlShortnerAPIs(router *mux.Router) {
	//Health check and rotation from ELB API's
	router.HandleFunc("/health_check", api.HealthCheck).Methods("GET", "HEAD")
	router.HandleFunc("/oor", api.OutOfRotation).Methods("GET", "HEAD")
	router.HandleFunc("/bir", api.BackInRotation).Methods("GET", "HEAD")
	//URL Shrinker APIs
	router.HandleFunc("/redirect/{url_id}", api.RedirectRequest).Methods("GET")
	router.HandleFunc("/shorten", api.ShortenURL).Methods("POST")
	//Stats API's for URL clicks
	router.HandleFunc("/stats", api.Stats).Methods("GET")

	//PProf apis
	router.HandleFunc("/debug/pprof/", pprof.Index)
	router.HandleFunc("/debug/pprof/cmdline", pprof.Cmdline)
	router.HandleFunc("/debug/pprof/profile", pprof.Profile)
	router.HandleFunc("/debug/pprof/symbol", pprof.Symbol)
	router.HandleFunc("/debug/pprof/trace", pprof.Trace)

	//Metrics apis from prometheus
	router.HandleFunc("/metrics", promhttp.Handler().ServeHTTP)
}
