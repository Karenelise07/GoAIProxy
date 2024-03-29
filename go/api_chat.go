/*
 * OpenAI API
 *
 * The OpenAI REST API. Please see https://platform.openai.com/docs/api-reference for more details.
 *
 * API version: 2.0.0
 * Generated by: Swagger Codegen (https://github.com/swagger-api/swagger-codegen.git)
 */
package swagger

import (
	"context"
	"encoding/json"
	"net/http"

	openai "github.com/sashabaranov/go-openai"
)

const DefaultModel = "qwen1.5-chat"

func CreateChatCompletion(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json; charset=UTF-8")
	w.WriteHeader(http.StatusOK)
	// 从查询参数中获取API密钥
	apiKey := r.URL.Query().Get("api_key")
	if apiKey == "" {
		apiKey = "empty"
	}

	// 从查询参数中获取BaseURL，如果未提供，则使用默认值
	baseURL := r.URL.Query().Get("base_url")
	if baseURL == "" {
		baseURL = "http://172.21.44.125:8091/v1" // 这里设置为您的默认BaseURL
	}

	// 使用提供的API密钥创建OpenAI客户端
	config := openai.DefaultConfig(apiKey)
	client := openai.NewClientWithConfig(config)

	var requestBody CreateChatCompletionRequest
	if err := json.NewDecoder(r.Body).Decode(&requestBody); err != nil {
		logError(w, "Error decoding request body: "+err.Error(), http.StatusBadRequest)
		return
	}

	chatRequest := openai.ChatCompletionRequest{
		Model: DefaultModel,
		Messages: []openai.ChatCompletionMessage{
			{Role: openai.ChatMessageRoleSystem, Content: "You are a helpful assistant."},
			{Role: openai.ChatMessageRoleUser, Content: requestBody.Messages[0].ChatCompletionRequestUserMessage.Content.Content},
		},
	}

	setOptionalFields(&chatRequest, requestBody)

	resp, err := client.CreateChatCompletion(context.Background(), chatRequest)
	if err != nil {
		logError(w, "ChatCompletion error: "+err.Error(), http.StatusInternalServerError)
		return
	}

	lastMessage := resp.Choices[len(resp.Choices)-1].Message.Content

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(map[string]string{"response": lastMessage})
}

func setOptionalFields(request *openai.ChatCompletionRequest, body CreateChatCompletionRequest) {
	if body.Model != "" {
		request.Model = body.Model
	}
	if body.Temperature != 0 {
		request.Temperature = float32(body.Temperature)
	}
	if body.MaxTokens != 0 {
		request.MaxTokens = int(body.MaxTokens)
	}
	if body.TopP != 0 {
		request.TopP = float32(body.TopP)
	}
	if body.FrequencyPenalty != 0 {
		request.FrequencyPenalty = float32(body.FrequencyPenalty)
	}
	if body.PresencePenalty != 0 {
		request.PresencePenalty = float32(body.PresencePenalty)
	}
}
