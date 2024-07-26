package openai

import (
    "fmt"
    "github.com/pkoukk/tiktoken-go"
)

func NumTokens(text string) int {
	encoding := "o200k_base" // gpt-4o
	// encoding := "cl100k_base" // gpt-4

	tke, err := tiktoken.GetEncoding(encoding)
	if err != nil {
		err = fmt.Errorf("getEncoding: %v", err)
		fmt.Println("ERROR")
		return 0
	}

	// encode
	token := tke.Encode(text, nil, nil)

	// tokens
	fmt.Println((token))
	// num_tokens
	fmt.Println(len(token))
	return len(token)
}
