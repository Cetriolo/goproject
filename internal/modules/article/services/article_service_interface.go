package services

import (
	ArticleResponse "v2/internal/modules/article/responses"
)

type ArticleServiceInterface interface {
	GetFeaturedArticles() ArticleResponse.Articles
	GetStoredArticles() ArticleResponse.Articles
}
