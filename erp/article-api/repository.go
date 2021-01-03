package main

import (
	"article-api/internal/core/domain"
	"context"
	"errors"

	"github.com/go-kit/kit/log"
	"gorm.io/gorm"
)

// Repository provides operations on database
type Repository interface {
	Migrate() (string, error)
	Create(ctx context.Context, a domain.Article) (string, error)
	Get(ctx context.Context, id uint32) (domain.Article, error)
	GetList(ctx context.Context, typ string, page int) ([]domain.ArticleListItem, error)
	Update(ctx context.Context, id uint32, title string, description string, group string, typ string) (string, error)
	Delete(ctx context.Context, id uint32) (string, error)
}

// repository is a concrete implementation of Repository
type repository struct {
	logger log.Logger
	db     *gorm.DB
}

// NewRepository creates a instance of repository
func NewRepository(logger log.Logger, db *gorm.DB) Repository {
	return &repository{
		logger: log.With(logger, "repo", "sqlite3"),
		db:     db,
	}
}

func (r *repository) Migrate() (string, error) {

	r.db.AutoMigrate(&domain.Article{})
	return "Migration sucessfully", nil
}

func (r *repository) Create(ctx context.Context, a domain.Article) (string, error) {

	r.db.Create(&a)
	return "s", nil
}

func (r *repository) Get(ctx context.Context, id uint32) (domain.Article, error) {
	var article domain.Article
	r.db.Find(&article, id)

	return article, nil
}

//GetList of Article
func (r *repository) GetList(ctx context.Context, typ string, page int) ([]domain.ArticleListItem, error) {
	var articles []domain.Article
	r.db.Where("typ = ?", typ).Find(&articles)
	var listItems []domain.ArticleListItem
	for _, a := range articles {
		listItems = append(listItems, domain.ArticleListItem{a.ID, a.ArticleID, a.Title, a.Description, a.Group, a.Typ, "", ""})
	}
	return listItems, nil
}

// Update article
func (r *repository) Update(ctx context.Context, id uint32, title string, description string, group string, typ string) (string, error) {
	var article domain.Article
	r.db.First(&article, id)
	article.Title = title
	article.Description = description
	article.Group = group
	article.Typ = typ
	r.db.Save(&article)
	return "update successfully", nil
}

// Delete article
func (r *repository) Delete(ctx context.Context, id uint32) (string, error) {
	r.db.Delete(&domain.Article{}, id)

	return "delete successfully", nil
}

var RepoErr = errors.New("Unable to handle Repo Request")
