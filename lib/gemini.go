package gemini

import (
	"context"
	"fmt"
	gen "github.com/google/generative-ai-go/genai"
	"google.golang.org/api/iterator"
	"google.golang.org/api/option"
	"strings"
)

func GenerateContent(apiKey string, modelName string, query string) (string, error) {
	ctx := context.Background()
	client, err := gen.NewClient(ctx, option.WithAPIKey(apiKey))
	if err != nil {
		return "", err
	}
	defer client.Close()
	model := client.GenerativeModel(modelName)
	model.Temperature = gen.Ptr[float32](0)

	resp, err := model.GenerateContent(ctx, gen.Text(query))
	if err != nil {
		return "", err
	}
	return responseString(resp), nil
}

// GenerateContentStream 流式生成内容
func GenerateContentStream(apiKey string, modelName string, query string) ([]string, error) {
	ctx := context.Background()
	client, err := gen.NewClient(ctx, option.WithAPIKey(apiKey))
	if err != nil {
		return nil, err
	}
	defer client.Close()

	model := client.GenerativeModel(modelName)
	model.Temperature = gen.Ptr[float32](0)

	iter := model.GenerateContentStream(ctx, gen.Text(query))

	var responses []string
	for {
		resp, err := iter.Next()
		if err == iterator.Done {
			break
		}
		if err != nil {
			return nil, err
		}
		responses = append(responses, responseString(resp))
	}
	return responses, nil
}

func responseString(resp *gen.GenerateContentResponse) string {
	var b strings.Builder
	for i, cand := range resp.Candidates {
		if len(resp.Candidates) > 1 {
			fmt.Fprintf(&b, "%d:", i+1)
		}
		b.WriteString(contentString(cand.Content))
	}
	return b.String()
}

func contentString(c *gen.Content) string {
	var b strings.Builder
	if c == nil || c.Parts == nil {
		return ""
	}
	for i, part := range c.Parts {
		if i > 0 {
			fmt.Fprintf(&b, ";")
		}
		fmt.Fprintf(&b, "%v", part)
	}
	return b.String()
}
