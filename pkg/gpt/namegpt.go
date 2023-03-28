package gpt

import (
	"context"
	"fmt"
	"log"
	"os"
	"path"
	"time"

	"github.com/briandowns/spinner"
	"github.com/sashabaranov/go-openai"
	gpt3 "github.com/sashabaranov/go-openai"
)

type Consultant interface {
	GetNamingSuggestions(filePath string)
}

type gptConsultant struct {
	client *gpt3.Client
}

func NewConsultant(apiKey string) Consultant {
	return &gptConsultant{client: gpt3.NewClient(apiKey)}
}

func (c *gptConsultant) GetNamingSuggestions(filePath string) {
	content, _ := os.ReadFile(filePath)
	lang := getFileLanguage(path.Ext(filePath))

	resChan := make(chan openai.ChatCompletionResponse)
	errChan := make(chan error)

	go func() {
		res, err := c.client.CreateChatCompletion(
			context.Background(), openai.ChatCompletionRequest{
				Model: openai.GPT3Dot5Turbo0301,
				Messages: []openai.ChatCompletionMessage{
					{
						Role:    openai.ChatMessageRoleUser,
						Content: getReviewPrompt(lang) + string(content),
					},
				},
			})
		if err != nil {
			errChan <- err
			return
		}
		resChan <- res
	}()

	waitingSpinner := spinner.New(spinner.CharSets[11], 100*time.Millisecond) // create new spinner
	waitingSpinner.Suffix = " Generating naming suggestions..."
	waitingSpinner.Start()

	select {
	case res := <-resChan:
		waitingSpinner.Stop() // stop spinner when results are received
		for _, choice := range res.Choices {
			printSuggestionsTable(choice.Message.Content)
		}

	case err := <-errChan:
		waitingSpinner.Stop() // stop spinner when results are received
		log.Fatalln(err)
	}
}

func getFileLanguage(fileExt string) string {
	switch fileExt {
	case ".go":
		return "golang"
	case ".py":
		return "python"
	case ".js":
		return "javascript"
	case ".ts":
		return "typescript"
	case ".java":
		return "java"
	case ".c":
		return "c"
	case ".cpp":
		return "cpp"
	case ".rs":
		return "rust"
	default:
		return "open source "
	}
}

func getReviewPrompt(lang string) string {
	promptTemplate := `Code Review Feedback on Readability and Maintainability

Please review the following code and provide feedback on the readability and maintainability using the common naming convention used in %s projects. 
After reviewing, please list all variable and function names that can be improved in the following format: 
\{old name | new name | one-sentence summary of the reason\}
Please focus only on naming conventions and do not provide any additional feedback.
Please do not output any content outside of the specified format \{old name | new name | one-sentence summary of the reason\}, as my command line tool needs to read that format and parse it.

example:
package main

import "fmt"

func f(a, b int) int { return a + b }

func main() {
	fmt.Println(f(1, 2))
}
output:
{f|add|"add" more accurately describes the function's purpose}
{...}

Code:
`
	return fmt.Sprintf(promptTemplate, lang)
}
