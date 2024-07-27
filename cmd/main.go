package main

import (
	"log"
	"fmt"
	"github.com/abava00/KakinoTane/internal/news"
	"github.com/abava00/KakinoTane/internal/ai"
	"github.com/abava00/KakinoTane/internal/misskey"
)

func main() {
	// news.Test()
	// return
	client, _ := misskey.NewMisskeyClient()
	client.GetStats()
	pre_text := `下記は News API から取得された最新のニュース記事の見出しと要約の一覧です。
以下の Description を読みつつ、2つ以上の記事の内容を上手く絡めることによって、騙されるまではいかない冗談交じりの偽記事を作成してください。
嘘記事には最初に\"タイトル: {TITLE}\"として見出しを設定し、以降に本文を記載するようにしてください。
`
	articles, _ := news.GetNews()
	articles, _ = news.GetRandomExtractNews(articles)
	for i, v := range articles {
		pre_text = pre_text + fmt.Sprintf("%d: %s\n", i, news.StringArticle(v))
	}
  log.Println("Prompt:", pre_text)
	openai.NumTokens(pre_text)
	text := openai.Openai(pre_text)
  text = "**※ この記事は LLM によって生成された虚構記事です。真実の内容が含まれることもありますが、基本的には嘘の内容であることをくれぐれもご理解ください。**\n\n" + text
	openai.NumTokens(text)
	log.Println("Output:", text)
	err := client.CreateNote(text)
	if err != nil {
		log.Fatalf("Failed to create note")
	}
}

