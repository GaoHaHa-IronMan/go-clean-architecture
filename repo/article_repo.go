package repo

import (
	"context"
	"go-clean-architecture/model"
	"gorm.io/gorm"
	"time"
)

type mysqlArticleRepository struct {
	DB *gorm.DB
}

// NewMysqlArticleRepository will create an object that represent the article.Repository interface
func NewMysqlArticleRepository(DB *gorm.DB) IArticleRepo {
	return &mysqlArticleRepository{DB}
}

func (m *mysqlArticleRepository) Fetch(ctx context.Context, createdDate time.Time,
	num int) (res []model.Article, err error) {

	err = m.DB.WithContext(ctx).Model(&model.Article{}).
		Select("id,title,content, updated_at, created_at").
		Where("created_at > ?", createdDate).Limit(num).Find(&res).Error
	return
}
