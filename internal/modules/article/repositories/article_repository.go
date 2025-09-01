package repositories

import (
	articleModels "v2/internal/modules/article/models"
	log "v2/logger"
	"v2/pkg/database"

	"gorm.io/gorm"
)

type ArticleRepository struct {
	DB *gorm.DB
}

func New() *ArticleRepository {
	return &ArticleRepository{
		DB: database.Connection(),
	}
}

func (articleRepository *ArticleRepository) List(limit int) []articleModels.Article {
	var articles []articleModels.Article
	result := articleRepository.DB.Limit(limit).Joins("User").Order("rand()").Find(&articles)
	if result.Error != nil {
		log.Log.Error("Error fetching articles:", result.Error)
		return nil
	}
	return articles
}
func (articleRepository *ArticleRepository) Find(id int) (articleModels.Article, error) {
	var article articleModels.Article
	result := articleRepository.DB.First(&article, id)
	if result.Error != nil {
		log.Log.Error("Error fetching article:", result.Error)
		return articleModels.Article{}, result.Error
	}
	return article, nil
}
