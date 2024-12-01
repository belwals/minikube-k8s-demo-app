package main

import (
	"log/slog"
	"net/http"

	"github.com/belwals/minikube-k8s-demo-app/service/tiny-url-implementation/config"
)

func initialize() dependency {
	// injecting a logger to service and will be forarded to inner calls
	getLogger := config.NewLogger()

	return dependency{
		log:    getLogger,
		server: http.NewServeMux(),
	}
}

type dependency struct {
	log    *slog.Logger
	server *http.ServeMux
}

func main() {

}
