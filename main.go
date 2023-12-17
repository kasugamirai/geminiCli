package main

import (
	"flag"
	"fmt"
	"strings"
	genai "xy.com/gemini/lib" // 替换为你的实际包名
)

func main() {
	// 定义命令行参数
	apiKey := flag.String("apikey", "", "API key for the AI service")
	modelName := flag.String("model", "", "Name of the model to use")
	query := flag.String("query", "", "The query or text to process")
	stream := flag.Bool("stream", false, "Set true to use streaming method")

	flag.Parse()

	// 验证输入
	if *apiKey == "" || *modelName == "" || *query == "" {
		fmt.Println("API key, model name, and query are required.")
		return
	}

	// 根据用户选择调用相应的函数
	if *stream {
		// 使用流式方法
		responses, err := genai.GenerateContentStream(*apiKey, *modelName, *query)
		if err != nil {
			fmt.Printf("Error: %v\n", err)
			return
		}
		fmt.Println(strings.Join(responses, "\n"))
	} else {
		// 使用普通方法
		response, err := genai.GenerateContent(*apiKey, *modelName, *query)
		if err != nil {
			fmt.Printf("Error: %v\n", err)
			return
		}
		fmt.Println(response)
	}
}
