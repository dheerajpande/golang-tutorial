package models

type Article struct {
	Title       string `json:"title"`
	Description string `json:"description"`
	Content     string `jons:"content"`
}

var allArticles []Article

func GetArticles() []Article {
	articles := []Article{
		{Title: "Test title 1", Description: "Some test descriptions 1", Content: "Hello this one is first test article"},
		{Title: "Test title 3", Description: "Some test descriptions 2", Content: "Hello this one is second test article"},
	}
	return articles
}
