package repo

import (
	"context"
	"go-clean-architecture/model"
	"time"
)

// IArticleRepo represent the article's repository contract
type IArticleRepo interface {
	Fetch(ctx context.Context, createdDate time.Time, num int) (res []model.Article, err error)
}
