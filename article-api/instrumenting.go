package main

import (
	"article-api/internal/core/domain"
	"context"
	"fmt"
	"time"

	"github.com/go-kit/kit/metrics"
)

type instrumentingMiddleware struct {
	requestCount   metrics.Counter
	requestLatency metrics.Histogram
	countResult    metrics.Histogram
	next           Service
}

func (mw instrumentingMiddleware) MigrateRepo() (output string, err error) {
	defer func(begin time.Time) {
		lvs := []string{"method", "migrateRepo", "error", fmt.Sprint(err != nil)}
		mw.requestCount.With(lvs...).Add(1)
		mw.requestLatency.With(lvs...).Observe(time.Since(begin).Seconds())
	}(time.Now())

	output, err = mw.next.MigrateRepo()
	return
}

func (mw instrumentingMiddleware) CreateArticle(ctx context.Context, title string, description string, group string, typ string) (output string, err error) {
	defer func(begin time.Time) {
		lvs := []string{"method", "createArticle", "error", fmt.Sprint(err != nil)}
		mw.requestCount.With(lvs...).Add(1)
		mw.requestLatency.With(lvs...).Observe(time.Since(begin).Seconds())
	}(time.Now())

	output, err = mw.next.CreateArticle(ctx, title, description, group, typ)
	return
}

func (mw instrumentingMiddleware) GetArticle(ctx context.Context, articleID uint32) (output domain.Article, err error) {
	defer func(begin time.Time) {
		lvs := []string{"method", "getArticle", "error", fmt.Sprint(err != nil)}
		mw.requestCount.With(lvs...).Add(1)
		mw.requestLatency.With(lvs...).Observe(time.Since(begin).Seconds())
	}(time.Now())

	output, err = mw.next.GetArticle(ctx, articleID)
	return
}

func (mw instrumentingMiddleware) GetArticleList(ctx context.Context, typ string, page int) (output []domain.ArticleListItem, err error) {
	defer func(begin time.Time) {
		lvs := []string{"method", "getArticleList", "error", fmt.Sprint(err != nil)}
		mw.requestCount.With(lvs...).Add(1)
		mw.requestLatency.With(lvs...).Observe(time.Since(begin).Seconds())
	}(time.Now())

	output, err = mw.next.GetArticleList(ctx, typ, page)
	return
}

func (mw instrumentingMiddleware) UpdateArticle(ctx context.Context, articleID uint32, title string, description string, group string, typ string) (output string, err error) {
	defer func(begin time.Time) {
		lvs := []string{"method", "updateArticle", "error", fmt.Sprint(err != nil)}
		mw.requestCount.With(lvs...).Add(1)
		mw.requestLatency.With(lvs...).Observe(time.Since(begin).Seconds())
	}(time.Now())

	output, err = mw.next.UpdateArticle(ctx, articleID, title, description, group, typ)
	return
}

func (mw instrumentingMiddleware) DeleteArticle(ctx context.Context, articleID uint32) (output string, err error) {
	defer func(begin time.Time) {
		lvs := []string{"method", "deleteArticle", "error", fmt.Sprint(err != nil)}
		mw.requestCount.With(lvs...).Add(1)
		mw.requestLatency.With(lvs...).Observe(time.Since(begin).Seconds())
	}(time.Now())

	output, err = mw.next.DeleteArticle(ctx, articleID)
	return
}
