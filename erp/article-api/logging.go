package main

import (
	"article-api/internal/core/domain"
	"context"
	"time"

	"github.com/go-kit/kit/log"
)

type loggingMiddleware struct {
	logger log.Logger
	next   Service
}

func (mw loggingMiddleware) MigrateRepo() (output string, err error) {
	defer func(begin time.Time) {
		_ = mw.logger.Log(
			"method", "migrateRepo",
			"output", output,
			"err", err,
			"took", time.Since(begin),
		)
	}(time.Now())

	output, err = mw.next.MigrateRepo()
	return
}

func (mw loggingMiddleware) CreateArticle(ctx context.Context, title string, description string, group string, typ string) (output string, err error) {
	defer func(begin time.Time) {
		_ = mw.logger.Log(
			"method", "createArticle",
			"articleTitle", title,
			"articleDescription", description,
			"articleGroup", group,
			"articleTyp", typ,
			"output", output,
			"err", err,
			"took", time.Since(begin),
		)
	}(time.Now())

	output, err = mw.next.CreateArticle(ctx, title, description, group, typ)
	return
}

func (mw loggingMiddleware) GetArticle(ctx context.Context, articleID uint32) (output domain.Article, err error) {
	defer func(begin time.Time) {
		_ = mw.logger.Log(
			"method", "getArticle",
			"articleID", articleID,
			"output", output,
			"err", err,
			"took", time.Since(begin),
		)
	}(time.Now())

	output, err = mw.next.GetArticle(ctx, articleID)
	return
}

func (mw loggingMiddleware) GetArticleList(ctx context.Context, typ string, page int) (output []domain.ArticleListItem, err error) {
	defer func(begin time.Time) {
		_ = mw.logger.Log(
			"method", "getArticleList",
			"typ", typ,
			"page", page,
			"output", output,
			"err", err,
			"took", time.Since(begin),
		)
	}(time.Now())

	output, err = mw.next.GetArticleList(ctx, typ, page)
	return
}

func (mw loggingMiddleware) UpdateArticle(ctx context.Context, articleID uint32, title string, description string, group string, typ string) (output string, err error) {
	defer func(begin time.Time) {
		_ = mw.logger.Log(
			"method", "updateArticle",
			"articleID", articleID,
			"articleTitle", title,
			"articleDescription", description,
			"articleGroup", group,
			"articleTyp", typ,
			"output", output,
			"err", err,
			"took", time.Since(begin),
		)
	}(time.Now())

	output, err = mw.next.UpdateArticle(ctx, articleID, title, description, group, typ)
	return
}

func (mw loggingMiddleware) DeleteArticle(ctx context.Context, articleID uint32) (output string, err error) {
	defer func(begin time.Time) {
		_ = mw.logger.Log(
			"method", "deleteArticle",
			"articleID", articleID,
			"output", output,
			"err", err,
			"took", time.Since(begin),
		)
	}(time.Now())

	output, err = mw.next.DeleteArticle(ctx, articleID)
	return
}
