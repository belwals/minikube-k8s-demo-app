package controller

import (
	"context"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
)

type ApiRequest struct {
	Headers    map[string][]string
	Body       io.Reader
	PathParam  map[string]any
	QueryParam map[string][]string
	Method     string
}

func (r ApiRequest) String() string {
	return fmt.Sprintf(`ApiRequest{"Headers" : %v, "Body": %v, "PathParam": %v, "QueryParam": %v, "Method: %v}`,
		r.Headers, r.Body, r.PathParam, r.QueryParam, r.Method,
	)
}

type ApiResponse struct {
	ResponseBody interface{}
	StatusCode   int
}

type CustomHandler func(ctx context.Context, request ApiRequest) (ApiResponse, error)

func ResponseHandler(handler CustomHandler) http.HandlerFunc {

	return func(w http.ResponseWriter, r *http.Request) {
		Ctx := r.Context()
		request := ApiRequest{
			Headers:    r.Header,
			Body:       r.Body,
			Method:     r.Method,
			QueryParam: r.URL.Query(),
		}

		resp, err := handler(Ctx, request)
		if err != nil {
			errorBody := fmt.Sprintf(`{"error: : %s}`, err.Error())
			err = WriteJson(w, http.StatusInternalServerError, errorBody)
			if err != nil {
				panic("error occurred while writing response to Json")
			}
			return
		}
		err = WriteJson(w, resp.StatusCode, resp.ResponseBody)
		if err != nil {
			panic("error occurred while writing response to Json")
		}
	}
}

func WriteJson(w http.ResponseWriter, statusCode int, v any) error {
	w.WriteHeader(statusCode)
	return json.NewEncoder(w).Encode(v)
}
