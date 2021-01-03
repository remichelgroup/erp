package main

import (
	"context"

	"github.com/go-kit/kit/endpoint"
)

// Endpoints contain all service endpoints
type Endpoints struct {
	CreateArticle  endpoint.Endpoint
	GetArticle     endpoint.Endpoint
	GetArticleList endpoint.Endpoint
	UpdateArticle  endpoint.Endpoint
	DeleteArticle  endpoint.Endpoint
}

// MakeEndpoints create Endpoints
func MakeEndpoints(svc Service) Endpoints {
	return Endpoints{
		CreateArticle:  makeCreateArticleEndpoint(svc),
		GetArticle:     makeGetArticleEndpoint(svc),
		GetArticleList: makeGetArticleListEndpoint(svc),
		UpdateArticle:  makeUpdateArticleEndpoint(svc),
		DeleteArticle:  makeDeleteArticleEndpoint(svc),
	}
}

func makeCreateArticleEndpoint(svc Service) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (interface{}, error) {
		req := request.(CreateArticleRequest)
		res, err := svc.CreateArticle(ctx, req.Title, req.Description, req.Group, req.Typ)
		if err != nil {
			return CreateArticleResponse{res, err.Error()}, nil
		}
		return CreateArticleResponse{res, ""}, nil
	}
}

func makeGetArticleEndpoint(svc Service) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (interface{}, error) {
		req := request.(GetArticleRequest)
		res, err := svc.GetArticle(ctx, req.ArticleID)
		if err != nil {
			return GetArticleResponse{res.ArticleID, res.Title, res.Description, res.Group, res.Typ, res.CreatedAt, res.UpdatedAt, res.DeletedAt, err.Error()}, nil
		}
		return GetArticleResponse{res.ArticleID, res.Title, res.Description, res.Group, res.Typ, res.CreatedAt, res.UpdatedAt, res.DeletedAt, ""}, nil
	}
}

func makeGetArticleListEndpoint(svc Service) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (interface{}, error) {
		req := request.(GetArticleListRequest)
		res, err := svc.GetArticleList(ctx, req.Typ, req.Page)
		if err != nil {
			return GetArticleListResponse{res, err.Error()}, nil
		}
		return GetArticleListResponse{res, ""}, nil
	}
}

func makeUpdateArticleEndpoint(svc Service) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (interface{}, error) {
		req := request.(UpdateArticleRequest)
		res, err := svc.UpdateArticle(ctx, req.ArticleID, req.Title, req.Description, req.Group, req.Typ)
		if err != nil {
			return UpdateArticleResponse{res, err.Error()}, nil
		}
		return UpdateArticleResponse{res, ""}, nil
	}
}

func makeDeleteArticleEndpoint(svc Service) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (interface{}, error) {
		req := request.(DeleteArticleRequest)
		res, err := svc.DeleteArticle(ctx, req.ArticleID)
		if err != nil {
			return DeleteArticleResponse{res, err.Error()}, nil
		}
		return DeleteArticleResponse{res, ""}, nil
	}
}
