package main

import (
	"context"
	"fmt"
	"log/slog"
	"net/http"

	"github.com/belwals/minikube-k8s-demo-app/service/tiny-url-implementation/config"
	"github.com/belwals/minikube-k8s-demo-app/service/tiny-url-implementation/constants"
	"github.com/belwals/minikube-k8s-demo-app/service/tiny-url-implementation/controller"
	"github.com/belwals/minikube-k8s-demo-app/service/tiny-url-implementation/model"
	"github.com/belwals/minikube-k8s-demo-app/service/tiny-url-implementation/repository"
	"github.com/belwals/minikube-k8s-demo-app/service/tiny-url-implementation/service"
	confgiEnv "github.com/caarlos0/env/v11"
)

func initialize() dependency {
	// injecting a logger to service and will be forarded to inner calls
	getLogger := config.NewLogger()
	var env model.Environment

	err := confgiEnv.Parse(&env)
	if err != nil {
		panic(fmt.Sprintf("unable to load env variable, error came %v", err))
	}

	return dependency{log: getLogger, server: http.NewServeMux(), env: env}
}

type dependency struct {
	log    *slog.Logger
	server *http.ServeMux
	env    model.Environment
}

func main() {
	// bootstrap application dependency
	dep := initialize()
	// initialize api dependencieas and register APIs with server
	registerApi(dep)
	// starting http server on a given port
	http.ListenAndServe(fmt.Sprintf(":%d", dep.env.Port), dep.server)
}

func registerApi(dep dependency) {
	log := dep.log

	log.Info("starting the setup of web server")
	env := dep.env
	// inject required env variable for the service

	// create dependency for initialization
	mongoClient, err := config.NewMongoInput(env.MongoUsername, env.MongoPassword, env.MongoClusterUrl).NewClient(context.TODO())
	if err != nil {
		panic("unable to create repository")
	}
	tinyUrlRepo := repository.Client(*mongoClient)

	// Dependency initialization
	tinyService := service.NewTinyUrlService(env, tinyUrlRepo)

	restController := controller.NewRestController(log, tinyService)

	// Registering APIs
	dep.server.Handle(constants.ApiPathHonePageUrl, controller.ResponseHandler(restController.GettHomePage))
	dep.server.Handle(constants.ApiPathGetFullUrl, controller.ResponseHandler(restController.GetFullUrl))
	dep.log.Info("Registered API path", "url", constants.ApiPathGetFullUrl)

	dep.server.HandleFunc(constants.ApiPathGenerateShortUrl, controller.ResponseHandler(restController.GenerateShortUrl))
	dep.log.Info("Registered API path", "url", constants.ApiPathGenerateShortUrl)

}
