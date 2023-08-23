package main

import (
	"braingain/pdf"
	"context"
	"encoding/json"
	"flag"
	"fmt"
	"github.com/sashabaranov/go-openai"
	"log"
	"os"
)

const baseDir = "/Users/patrick/patrick.zierahn@gmail.com - Google Drive/My Drive/KIT/2023-SS/DeSys/Lecture Slides"

var (
	pdfFile = baseDir + "/DeSys_01_Intro.pdf"
	//pdfFile     = baseDir + "/DeSys_11_Payment_Channel_Networks.pdf"
	//pdfFile     = baseDir + "/DeSys_12_Smart_Contract_Platforms_Ethereum.pdf"
	//pdfFile     = baseDir + "/DeSys_13_Decentralized_File_Storage_IPFS.pdf"
	//pdfFile     = "/Users/patrick/patrick.zierahn@gmail.com - Google Drive/My Drive/KIT/2023-SS/DeSys/Further Readings/176429260X.pdf"
	first       = 2
	last        = 17
	temperature = 0.25
	model       = openai.GPT3Dot5Turbo16K
	prompt      = "What is a tragedy of the commons in decentralized systems?"
)

func main() {
	log.SetFlags(log.Lshortfile | log.LstdFlags)
	flag.Parse()

	log.Printf("Filename: %s\n", pdfFile)

	pages, err := pdf.ReadPages(context.Background(), pdfFile)
	if err != nil {
		log.Fatal(err)
	}
	log.Printf("Page count: %d\n", len(pages))

	var messages []openai.ChatCompletionMessage

	if first > last {
		log.Fatal("First page must be less than or equal to last page")
	}

	for _, page := range pages[first-1 : last] {
		messages = append(messages, openai.ChatCompletionMessage{
			Role:    openai.ChatMessageRoleUser,
			Content: page,
		})
	}

	messages = append(messages, openai.ChatCompletionMessage{
		Role:    openai.ChatMessageRoleUser,
		Content: prompt,
	})

	token := os.Getenv("OPENAI_API_KEY")
	client := openai.NewClient(token)

	resp, err := client.CreateChatCompletion(
		context.Background(),
		openai.ChatCompletionRequest{
			Model:       model,
			Temperature: float32(temperature),
			Messages:    messages,
			N:           1,
		},
	)

	if err != nil {
		fmt.Printf("ChatCompletion error: %v\n", err)
		return
	}

	usage, _ := json.MarshalIndent(resp.Usage, "", "  ")
	log.Printf("Usage: %s\n", usage)

	log.Println(resp.Choices[0].Message.Content)
	_ = os.WriteFile("output.txt", []byte(resp.Choices[0].Message.Content), 0644)
}
