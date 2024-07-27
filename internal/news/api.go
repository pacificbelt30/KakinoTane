package news

import (
	"os"
	"fmt"
	"net/http"
	"time"
	"math/rand"

	"github.com/robtec/newsapi/api"
)

func Test() {
	rand.Seed(time.Now().UnixNano())
	httpClient := http.Client{}
	key := os.Getenv("NEWS_API_KEY")
	url := "https://newsapi.org"
	// Create a client, passing in the above
	client, err := api.New(&httpClient, key, url)
	if err != nil {
		return
	}

	// Create options for Ireland and Business
	opts := api.Options{Country: "jp", Category: "technology", PageSize: 50}

	// Get Top Headlines with options from above
	topHeadlines, err := client.TopHeadlines(opts)
	fmt.Println(topHeadlines)
	fmt.Println(topHeadlines.Articles)
	fmt.Println(topHeadlines.Articles[0])
	for i, v := range topHeadlines.Articles {
		fmt.Println(i, v)
	}
	var shuffle_idx [50]int
	for i := 0; i < len(shuffle_idx); i = i+1 {
		shuffle_idx[i] = i
	}
	new_shuffle_idx := shuffle_idx[:len(topHeadlines.Articles)]
	rand.Shuffle(len(new_shuffle_idx), func(i, j int){
		new_shuffle_idx[i], new_shuffle_idx[j] = new_shuffle_idx[j], new_shuffle_idx[i]
	})
	for i, v := range new_shuffle_idx[:10] {
		fmt.Println(i, ": ", v, topHeadlines.Articles[v].Title, "\n内容: ", topHeadlines.Articles[v].Description)
	}
}

// func GetNews() (*api.Response, error){
func GetNews() ([]api.Article, error){
	httpClient := http.Client{}
	key := os.Getenv("NEWS_API_KEY")
	url := "https://newsapi.org"
	// query := "Security"
	// Create a client, passing in the above
	client, err := api.New(&httpClient, key, url)
	if err != nil {
		return nil, err
	}

	// Create options for Ireland and Business
	opts := api.Options{Country: "jp", Category: "technology", PageSize: 50}

	// Get Top Headlines with options from above
	topHeadlines, err := client.TopHeadlines(opts)
	return topHeadlines.Articles, err
}

func GetRandomExtractNews(articles []api.Article) ([]api.Article, error){
  extract_news := 10
	rand.Shuffle(len(articles), func(i, j int){
		articles[i], articles[j] = articles[j], articles[i]
	})
	
	return articles[:extract_news], nil
}

func StringArticle(art api.Article) string{
  return fmt.Sprintf("%s\n内容: %s", art.Title, art.Description)
}
