package service

import (
	"context"
	"go-clean-architecture/model"
	"time"
)

// IArticleService represent the article's usecases
type IArticleService interface {
	Fetch(ctx context.Context, createdDate time.Time, num int) ([]model.Article, error)
}
