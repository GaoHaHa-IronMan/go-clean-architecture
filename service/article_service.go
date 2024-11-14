package service

import (
	"context"
	"go-clean-architecture/model"
	"go-clean-architecture/repo"
	"time"
)

type articleService struct {
	articleRepo repo.IArticleRepo
}

// NewArticleService will create new an articleUsecase object representation of domain.ArticleUsecase interface
func NewArticleService(a repo.IArticleRepo) IArticleService {
	return &articleService{
		articleRepo: a,
	}
}

// Fetch
func (a *articleService) Fetch(ctx context.Context, createdDate time.Time, num int) (res []model.Article, err error) {
	if num == 0 {
		num = 10
	}
	res, err = a.articleRepo.Fetch(ctx, createdDate, num)
	if err != nil {
		return nil, err
	}
	return
}
