package services

import (
	//ArticleModel "v2/internal/modules/article/models"

	ArticleRepository "v2/internal/modules/article/repositories"
	ArticleResponse "v2/internal/modules/article/responses"
)

type ArticleService struct {
	articleRepository ArticleRepository.ArticleRepositoryInterface
}

func New() *ArticleService {
	return &ArticleService{
		articleRepository: ArticleRepository.New(),
	}
}

func (ArticleService *ArticleService) GetFeaturedArticles() ArticleResponse.Articles {
	articles := ArticleService.articleRepository.List(4)
	return ArticleResponse.ToArticles(articles)
}
func (ArticleService *ArticleService) GetStoredArticles() ArticleResponse.Articles {
	articles := ArticleService.articleRepository.List(8)
	return ArticleResponse.ToArticles(articles)
}
