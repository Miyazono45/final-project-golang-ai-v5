package service

import (
	"a21hc3NpZ25tZW50/model"
	"bytes"
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"net/http"
	// "strings"
)

type HTTPClient interface {
	Do(req *http.Request) (*http.Response, error)
}

type AIService struct {
	Client HTTPClient
}

func (s *AIService) AnalyzeData(table map[string][]string, query, token string) (string, error) {
	if len(table) == 0 {
		return "", errors.New("data cannot empty")
	}

	// HELP ME AAAA ====
	inputReq := &model.Inputs{
		Table: table,
		Query: query,
	}

	var jsonData, err = json.Marshal(inputReq)
	if err != nil {
		return "", err
	}
	// ALAMAK =========

	// prepare to connect model and create request
	url := "https://api-inference.huggingface.co/models/google/tapas-base-finetuned-wtq"
	req, err := http.NewRequest("POST", url, bytes.NewBuffer(jsonData))
	if err != nil {
		return "", err
	}

	// set header in request (content type and auth)
	req.Header.Add("Content-Type", "application/json")
	req.Header.Add("Authorization", "Bearer "+token)

	responseReq, err := s.Client.Do(req)
	if err != nil {
		return "", err
	}
	defer responseReq.Body.Close()

	if responseReq.StatusCode != http.StatusOK {
		return "", fmt.Errorf("error in model in status code : %s", responseReq.StatusCode)
	}

	strRespReq, _ := io.ReadAll(responseReq.Body)
	res := string(strRespReq)
	fmt.Println(res)

	return res, nil
}

func (s *AIService) ChatWithAI(context, query, token string) (model.ChatResponse, error) {
	inputReq := map[string]any{
		"model": "microsoft/Phi-3.5-mini-instruct",
		"messages": []map[string]string{
			{
				"role":    "assistant",
				"content": query,
			},
		},
		"max_tokens": 500,
		"stream":     false,
	}

	var jsonData, err = json.Marshal(inputReq)
	if err != nil {
		panic(err)
	}

	url := "https://api-inference.huggingface.co/models/microsoft/Phi-3.5-mini-instruct/v1/chat/completions"
	req, err := http.NewRequest("POST", url, bytes.NewBuffer(jsonData))
	if err != nil {
		return model.ChatResponse{}, err
	}

	req.Header.Add("Content-Type", "application/json")
	req.Header.Add("Authorization", "Bearer "+token)

	responseReq, err := s.Client.Do(req)
	if err != nil {
		return model.ChatResponse{}, fmt.Errorf("failed to send request: %w", err)
	}
	// defer responseReq.Body.Close()

	fmt.Println("INI CHECK")

	strRespReq, err := io.ReadAll(responseReq.Body)
	if err != nil {
		panic(err)
	}

	// PLEASE DONT ASK ME WHAT IS THIS

	if len(strRespReq) == 0 {
		return model.ChatResponse{}, errors.New("body is null")
	}

	// var chatResponse []model.ChatResponse
	// if err := json.Unmarshal(strRespReq, &chatResponse); err == nil {
	// 	if len(chatResponse) == 0 || strings.TrimSpace(chatResponse[0].GeneratedText) == "" {
	// 		return model.ChatResponse{}, errors.New("received empty or invalid response from Chat API")
	// 	}
	// 	chatResponse[0].GeneratedText = strings.TrimSpace(chatResponse[0].GeneratedText)
	// 	return chatResponse[0], nil
	// }

	// var chatResponses []model.ChatResponse
	// if err := json.Unmarshal(strRespReq, &chatResponses); err == nil {
	// 	if len(chatResponses) == 0 || strings.TrimSpace(chatResponses[0].GeneratedText) == "" {
	// 		return model.ChatResponse{}, errors.New("received empty or invalid response from Chat API")
	// 	}
	// 	chatResponses[0].GeneratedText = strings.TrimSpace(chatResponses[0].GeneratedText)
	// 	return chatResponses[0], nil
	// }

	// var chatResponse model.ChatResponse
	// if err := json.Unmarshal(strRespReq, &chatResponse); err != nil {
	// 	return model.ChatResponse{}, fmt.Errorf("failed to unmarshal response: %w", err)
	// }

	// if strings.TrimSpace(chatResponse.GeneratedText) == "" {
	// 	return model.ChatResponse{}, errors.New("received empty or invalid response from Chat API")
	// }

	// IM SORRY TO USING GPT FOR GENERATE THIS CODE :(
	type Message struct {
		Role    string `json:"role"`
		Content string `json:"content"`
	}

	type Choice struct {
		Index        int         `json:"index"`
		Message      Message     `json:"message"`
		Logprobs     interface{} `json:"logprobs"` // Use interface{} if the type is unknown or nil
		FinishReason string      `json:"finish_reason"`
	}

	type Usage struct {
		PromptTokens     int `json:"prompt_tokens"`
		CompletionTokens int `json:"completion_tokens"`
		TotalTokens      int `json:"total_tokens"`
	}

	type Response struct {
		Object            string   `json:"object"`
		ID                string   `json:"id"`
		Created           int64    `json:"created"`
		Model             string   `json:"model"`
		SystemFingerprint string   `json:"system_fingerprint"`
		Choices           []Choice `json:"choices"`
		Usage             Usage    `json:"usage"`
	}

	// Parse the response
	var result model.ChatResponse
	var response Response
	err = json.Unmarshal(strRespReq, &response)
	if err != nil {
		return model.ChatResponse{}, fmt.Errorf("failed to unmarshal response: %w", err)
	}

	// var result model.ChatResponse
	fmt.Println(response.Choices[0].Message.Content)
	result.GeneratedText = response.Choices[0].Message.Content
	// =======================================================================

	return result, nil
}
