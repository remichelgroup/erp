package main

import (
	"article-api/internal/core/domain"
	"context"
	"encoding/json"
	"net/http"
	"time"

	"gorm.io/gorm"
)

// CreateArticleRequest each method, we define request and response structs
type CreateArticleRequest struct {
	Title       string `json:"title"`
	Description string `json:"description"`
	Group       string `json:"group"`
	Typ         string `json:"typ"` // product/component
}

// CreateArticleResponse returns  the status
type CreateArticleResponse struct {
	Msg string `json:"msg"`
	Err string `json:"err,omitempty"`
}

type GetArticleRequest struct {
	ArticleID uint32
}

type GetArticleResponse struct {
	ArticleID   uint32
	Title       string
	Description string
	Group       string
	Typ         string
	CreatedAt   time.Time
	UpdatedAt   time.Time
	DeletedAt   gorm.DeletedAt
	Err         string `json:"err,omitempty"`
}

type GetArticleListRequest struct {
	Typ  string
	Page int
}

type GetArticleListResponse struct {
	ArticleList []domain.ArticleListItem
	Err         string `json:"err,omitempty"`
}

type UpdateArticleRequest struct {
	ArticleID   uint32
	Title       string
	Description string
	Group       string
	Typ         string
}

type UpdateArticleResponse struct {
	Msg string
	Err string `json:"err,omitempty"`
}

type DeleteArticleRequest struct {
	ArticleID uint32
}

type DeleteArticleResponse struct {
	Msg string
	Err string `json:"err,omitempty"`
}

func decodeCreateArticleRequest(_ context.Context, r *http.Request) (interface{}, error) {
	var request CreateArticleRequest
	if err := json.NewDecoder(r.Body).Decode(&request); err != nil {
		return nil, err
	}
	return request, nil
}

func decodeGetArticleRequest(_ context.Context, r *http.Request) (interface{}, error) {
	var request GetArticleRequest
	if err := json.NewDecoder(r.Body).Decode(&request); err != nil {
		return nil, err
	}
	return request, nil
}

func decodeGetArticleListRequest(_ context.Context, r *http.Request) (interface{}, error) {
	var request GetArticleListRequest
	if err := json.NewDecoder(r.Body).Decode(&request); err != nil {
		return nil, err
	}
	return request, nil
}

func decodeUpdateArticleRequest(_ context.Context, r *http.Request) (interface{}, error) {
	var request UpdateArticleRequest
	if err := json.NewDecoder(r.Body).Decode(&request); err != nil {
		return nil, err
	}
	return request, nil
}

func decodeDeleteArticleRequest(_ context.Context, r *http.Request) (interface{}, error) {
	var request DeleteArticleRequest
	if err := json.NewDecoder(r.Body).Decode(&request); err != nil {
		return nil, err
	}
	return request, nil
}

func encodeResponse(_ context.Context, w http.ResponseWriter, response interface{}) error {
	return json.NewEncoder(w).Encode(response)
}
