package utils

import (
	"bytes"
	"context"
	"encoding/json"
	"errors"
	"go-cms-gql/graph/model"
	"net/http"
	"strconv"
)

type OpenAIRequest struct {
	Model    string    `json:"model"`
	Messages []Message `json:"messages"`
}

type OpenAIResponse struct {
	ID                string   `json:"id"`
	Object            string   `json:"object"`
	Created           int64    `json:"created"`
	Model             string   `json:"model"`
	Choices           []Choice `json:"choices"`
	Usage             Usage    `json:"usage"`
	SystemFingerprint string   `json:"system_fingerprint"`
}

type Choice struct {
	Index        int64       `json:"index"`
	Message      Message     `json:"message"`
	Logprobs     interface{} `json:"logprobs"`
	FinishReason string      `json:"finish_reason"`
}

type Message struct {
	Role    string      `json:"role"`
	Content string      `json:"content"`
	Refusal interface{} `json:"refusal"`
}

type Usage struct {
	PromptTokens            int64                   `json:"prompt_tokens"`
	CompletionTokens        int64                   `json:"completion_tokens"`
	TotalTokens             int64                   `json:"total_tokens"`
	PromptTokensDetails     PromptTokensDetails     `json:"prompt_tokens_details"`
	CompletionTokensDetails CompletionTokensDetails `json:"completion_tokens_details"`
}

type CompletionTokensDetails struct {
	ReasoningTokens int64 `json:"reasoning_tokens"`
}

type PromptTokensDetails struct {
	CachedTokens int64 `json:"cached_tokens"`
}

func GenerateContent(ctx context.Context, generateInput model.GenerateContent) (string, error) {
	// request
	readDuration := strconv.Itoa(generateInput.Duration)

	requestBody := OpenAIRequest{
		Model: GetValue("OPENAI_MODEL"),
		Messages: []Message{
			{
				Role:    "system",
				Content: "You are a content specialist in " + generateInput.Topic + " topic",
			},
			{
				Role:    "user",
				Content: "Write an article that covers " + generateInput.Title + " with a read duration of " + readDuration + " minutes.",
			},
		},
	}

	response, err := sendRequest(ctx, requestBody)

	if err != nil {
		return "", err
	}

	responseContent := response.Choices[0].Message.Content

	return responseContent, nil
}

func GetTags(ctx context.Context, input model.GetTag) ([]string, error) {
	// request
	requestBody := OpenAIRequest{
		Model: GetValue("OPENAI_MODEL"),
		Messages: []Message{
			{
				Role:    "system",
				Content: "you are a helpful assistant",
			},
			{
				Role:    "user",
				Content: "create a tags based on the given article in array of string format. return string of tags without any explanation. this is the article: " + input.Content,
			},
		},
	}

	response, err := sendRequest(ctx, requestBody)

	if err != nil {
		return nil, err
	}

	responseContent := response.Choices[0].Message.Content
	var tags []string

	if err := json.Unmarshal([]byte(responseContent), &tags); err != nil {
		return nil, errors.New("error parsing response")
	}

	return tags, nil
}

func sendRequest(ctx context.Context, requestBody OpenAIRequest) (OpenAIResponse, error) {
	jsonData, err := json.Marshal(requestBody)
	if err != nil {
		return OpenAIResponse{}, errors.New("error marshalling request body")
	}

	// send the request
	apiKey := GetValue("OPENAI_API_KEY")
	endpoint := "https://api.openai.com/v1/chat/completions"
	body := bytes.NewBuffer(jsonData)

	// Create a new HTTP POST request
	req, err := http.NewRequestWithContext(ctx, http.MethodPost, endpoint, body)
	if err != nil {
		return OpenAIResponse{}, errors.New("error creating request")
	}

	// Set the headers
	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("Authorization", "Bearer "+apiKey)

	// Send the request
	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		return OpenAIResponse{}, errors.New("error sending request")
	}
	defer resp.Body.Close()

	var data OpenAIResponse

	err = json.NewDecoder(resp.Body).Decode(&data)
	if err != nil {
		return OpenAIResponse{}, errors.New("error parsing response")
	}

	return data, nil
}
