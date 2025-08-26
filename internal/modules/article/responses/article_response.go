package responses

import (
	"fmt"
	ArticleModels "v2/internal/modules/article/models"
	UserResponse "v2/internal/modules/user/responses"
)

type Article struct {
	ID        uint
	Image     string
	Title     string
	Content   string
	CreatedAt string
	User      UserResponse.User
}

type Articles struct {
	Data []Article
}

func ToArticle(article ArticleModels.Article) Article {
	return Article{
		ID:        article.ID,
		Title:     article.Title,
		Image:     "/assets/images/demopic/10.jpg",
		Content:   article.Content,
		CreatedAt: fmt.Sprintf("%d/%02d/%02d", article.CreatedAt.Year(), article.CreatedAt.Month(), article.CreatedAt.Day()),
		User:      UserResponse.ToUser(article.User),
	}
}

func ToArticles(articles []ArticleModels.Article) Articles {
	var response Articles
	for _, article := range articles {
		response.Data = append(response.Data, ToArticle(article))
	}
	return response
}
