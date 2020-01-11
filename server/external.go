package server

import (
	"fmt"
	"github.com/gorilla/mux"
	"github.com/tylerb/graceful"
	"log"
	"net/http"
	"runtime/debug"
	"time"
	"url-shortner/routes"
)

var (
	srv          *graceful.Server
	httpShutdown = make(chan struct{})
)

// Start starts the internal server
func StartServer() {
	defer doAPIPanicRecovery("url-shortner")

	r := mux.NewRouter()

	routes.UrlShortnerAPIs(r)

	log.Printf(logTag+" Starting server on %s", "0.0.0.0:8080")

	srv = &graceful.Server{
		Server: &http.Server{
			Addr:           "0.0.0.0:8080",
			Handler:        r,
			ReadTimeout:    10 * time.Second,
			WriteTimeout:   50 * time.Second,
			MaxHeaderBytes: 1 << 20,
		},
		ShutdownInitiated: func() {
			// TODO:
		},
		Timeout: time.Second * 3,
	}

	err := srv.ListenAndServe()
	if err != nil {
		log.Fatalf(logTag+" %s", err)
	}

	close(httpShutdown)
	log.Printf(logTag+" %v", "url-shortner HTTP server exits")
}

func doAPIPanicRecovery(serviceTag string) {
	if r := recover(); r != nil {
		logMessage := fmt.Sprintf("%s%s service got exception, failed with error %s %s", "[url-shortner Alert]", serviceTag, r, string(debug.Stack()))
		log.Printf(logTag+" %v", logMessage)
	}
}

func stopHTTPServer() <-chan struct{} {
	return httpShutdown
}
