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
	pre_text := `あなたはニュース記事の作成する記事執筆者です。（決して，フィクション小説家ではありません。）
下記に News API から取得された最新のニュース記事の見出しと要約の一覧を示します。
以下の Description を読みつつ、2つ以上の記事の内容を上手く絡めることによって、騙されるまではいかないが実際にありそうな冗談交じりの偽記事を作成してください。ただし、記事中の固有名詞はすべて名前の似た架空のものに変換して、絶対に存在する企業名、大学名、個人名等をそのまま使わないようにしてください。（見出しにも固有名詞をそのまま使うことは許されません。）
嘘記事には最初に\"タイトル: {TITLE} - {SITE}\"として見出しと架空の Web サイト名を設定し、以降に本文を記載するようにしてください。そして、最後に架空の記事執筆者と所属を記載してください。
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

