package main

import (
	"context"
	"net/http"

	httptransport "github.com/go-kit/kit/transport/http"
	"github.com/gorilla/mux"
	"github.com/prometheus/client_golang/prometheus/promhttp"
)

// NewHTTPServer create new HttpServer
func NewHTTPServer(ctx context.Context, endpoints Endpoints) http.Handler {
	r := mux.NewRouter()
	r.Use(commonMiddleware)
	r.Handle("/metrics", promhttp.Handler())
	r.Methods("POST").Path("/create").Handler(httptransport.NewServer(
		endpoints.CreateArticle,
		decodeCreateArticleRequest,
		encodeResponse,
	))
	r.Methods("POST").Path("/get").Handler(httptransport.NewServer(
		endpoints.GetArticle,
		decodeGetArticleRequest,
		encodeResponse,
	))
	r.Methods("POST").Path("/list").Handler(httptransport.NewServer(
		endpoints.GetArticleList,
		decodeGetArticleListRequest,
		encodeResponse,
	))
	r.Methods("POST").Path("/update").Handler(httptransport.NewServer(
		endpoints.UpdateArticle,
		decodeUpdateArticleRequest,
		encodeResponse,
	))
	r.Methods("POST").Path("/delete").Handler(httptransport.NewServer(
		endpoints.DeleteArticle,
		decodeDeleteArticleRequest,
		encodeResponse,
	))
	return r
}

func commonMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Add("Content-Type", "application/json")
		if r.Method == "OPTIONS" {
			return
		}
		next.ServeHTTP(w, r)
	})
}
