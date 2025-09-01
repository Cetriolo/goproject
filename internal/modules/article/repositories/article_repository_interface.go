package repositories

import (
	articleModels "v2/internal/modules/article/models"
)

type ArticleRepositoryInterface interface {
	List(limit int) []articleModels.Article
	Find(id int) (articleModels.Article, error)
}
