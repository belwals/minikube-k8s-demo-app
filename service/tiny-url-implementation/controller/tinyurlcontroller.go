package controller

import (
	"context"
	"encoding/json"
	"io"
	"log/slog"
	"net/http"

	"github.com/belwals/minikube-k8s-demo-app/service/tiny-url-implementation/model"
	"github.com/belwals/minikube-k8s-demo-app/service/tiny-url-implementation/service"
)

type RestUrlController struct {
	log     *slog.Logger
	service service.IRestService
}

func NewRestController(log *slog.Logger, service service.IRestService) RestUrlController {
	return RestUrlController{
		log:     log,
		service: service,
	}
}

func (cntrler *RestUrlController) GettHomePage(ctx context.Context, request ApiRequest) (ApiResponse, error) {

	cntrler.log.Info("accessing home route of tiny-url-service")
	resp := `
	<h1>Welcome to the webpage of Tiny URL service </h1>
	<p>
		Thanks for reaching out, this process is under developed.
		Try below route for doing the url related queries
		1. GET /v1/get-url?key=<shortend url key>
		2. POST /v1/generate-url with body {"url": "<sample url>"}
	</p>
	`
	return buildApiRespose(resp, http.StatusOK), nil
}

func (cntrler *RestUrlController) GetFullUrl(ctx context.Context, request ApiRequest) (ApiResponse, error) {

	cntrler.log.Info("received request with", "request", request.String())

	tinyIdentifiers := request.QueryParam["key"]
	if len(tinyIdentifiers) == 0 {
		return buildApiRespose("bad request", http.StatusBadRequest), nil
	}
	resp, err := cntrler.service.GetFullURl(ctx, tinyIdentifiers[0])
	if err != nil {
		return buildApiRespose(resp, http.StatusInternalServerError), nil
	}

	return buildApiRespose(resp, http.StatusOK), nil
}

func (cntrler *RestUrlController) GenerateShortUrl(ctx context.Context, request ApiRequest) (ApiResponse, error) {

	cntrler.log.Info("received request with", "body", request.Body)
	bodyBytes, err := io.ReadAll(request.Body)
	if err != nil || bodyBytes == nil || len(bodyBytes) == 0 {
		errorBody := struct {
			Msg string
		}{Msg: "Invalid request body"}
		return buildApiRespose(errorBody, http.StatusBadRequest), nil
	}
	var requestPayload model.RequestForUrlShortening
	err = json.Unmarshal(bodyBytes, &requestPayload)
	if err != nil {
		errorBody := struct {
			Msg string
		}{Msg: "Unable to parse incoming request payload"}
		return buildApiRespose(errorBody, http.StatusBadRequest), nil
	}

	resp, err := cntrler.service.GenerateShortUrl(ctx, requestPayload.Url)
	if err != nil {
		return buildApiRespose(resp, http.StatusInternalServerError), nil
	}

	return buildApiRespose(resp, http.StatusOK), nil
}

func buildApiRespose(body interface{}, code int) ApiResponse {
	return ApiResponse{
		ResponseBody: body,
		StatusCode:   code,
	}
}
