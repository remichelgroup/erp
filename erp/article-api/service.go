package main

import (
	"article-api/internal/core/domain"
	"context"
	"errors"

	"github.com/go-kit/kit/log"
	"github.com/go-kit/kit/log/level"
)

// Service provides operations on strings.
type Service interface {
	MigrateRepo() (string, error)
	CreateArticle(ctx context.Context, title string, description string, group string, typ string) (string, error)
	GetArticle(ctx context.Context, articleID uint32) (domain.Article, error)
	GetArticleList(ctx context.Context, typ string, page int) ([]domain.ArticleListItem, error)
	UpdateArticle(ctx context.Context, articleID uint32, title string, description string, group string, typ string) (string, error)
	DeleteArticle(ctx context.Context, articleID uint32) (string, error)
}

// service is a concrete implementation of Service
type service struct {
	repository Repository
	logger     log.Logger
}

// NewService create a new instance of Service
func NewService(repository Repository, logger log.Logger) Service {
	return &service{
		repository: repository,
		logger:     logger,
	}
}
func (svc service) MigrateRepo() (string, error) {
	logger := log.With(svc.logger, "method", "MigrateRepo")
	ok, err := svc.repository.Migrate()
	if err != nil {
		level.Error(logger).Log("err", err)
		return "", err
	}
	return ok, nil
}
func (svc service) CreateArticle(ctx context.Context, title string, description string, group string, typ string) (string, error) {
	logger := log.With(svc.logger, "method", "CreateArticle")
	a := domain.Article{Title: title, Description: description, Group: group, Typ: typ}
	res, err := svc.repository.Create(ctx, a)
	if err != nil {
		level.Error(logger).Log("err", err)
		return "", err
	}

	if res == "" {
		return "", ErrEmpty
	}

	return res, nil
}

func (svc service) GetArticle(ctx context.Context, articleID uint32) (domain.Article, error) {
	logger := log.With(svc.logger, "method", "CreateArticle")
	res, err := svc.repository.Get(ctx, articleID)
	if err != nil {
		level.Error(logger).Log("err", err)
		return res, err
	}
	return res, nil
}

func (svc service) GetArticleList(ctx context.Context, typ string, page int) ([]domain.ArticleListItem, error) {
	logger := log.With(svc.logger, "method", "GetArticleList")
	res, err := svc.repository.GetList(ctx, typ, page)
	if err != nil {
		level.Error(logger).Log("err", err)
		return res, err
	}
	return res, nil
}

func (svc service) UpdateArticle(ctx context.Context, articleID uint32, title string, description string, group string, typ string) (string, error) {
	logger := log.With(svc.logger, "method", "UpdateArticle")
	res, err := svc.repository.Update(ctx, articleID, title, description, group, typ)
	if err != nil {
		level.Error(logger).Log("err", err)
		return res, err
	}
	return res, nil
}

func (svc service) DeleteArticle(ctx context.Context, articleID uint32) (string, error) {
	logger := log.With(svc.logger, "method", "DeleteArticle")
	res, err := svc.repository.Delete(ctx, articleID)
	if err != nil {
		level.Error(logger).Log("err", err)
		return res, err
	}
	return res, nil
}

// ErrEmpty is returned when an input string is empty.
var ErrEmpty = errors.New("empty string")
