package ggpt

import (
	"bytes"
	"encoding/json"
	"io"
	"net/http"
)

// GGPTAgent implements GGPT agent.
type GGPTAgent struct {
	Config  Config
	History []message
}

type message struct {
	Role    string `json:"role"`
	Content string `json:"content"`
}

// GPTRequest implements the structure of a GPT request.
type GPTRequest struct {
	Model     string    `json:"model"`
	Messages  []message `json:"messages"`
	MaxTokens int       `json:"max_tokens"`
}

type choice struct {
	Index        int     `json:"index"`
	Message      message `json:"message"`
	FinishReason string  `json:"finish_reason"`
}

type usage struct {
	PromptTokens     int `json:"prompt_tokens"`
	CompletionTokens int `json:"completion_tokens"`
	TotalTokens      int `json:"total_tokens"`
}

// GPTResponse implements the structure of a GPT response.
type GPTResponse struct {
	ID      string   `json:"id"`
	Object  string   `json:"object"`
	Created int64    `json:"created"`
	Model   string   `json:"model"`
	Choices []choice `json:"choices"`
	Usage   usage    `json:"usage"`
}

// NewAgent creates a new GGPT agent.
func NewAgent(role string) (*GGPTAgent, error) {
	config, err := LoadConfig()
	if err != nil {
		return nil, err
	}

	return &GGPTAgent{
		Config: *config,
		History: []message{
			{
				Role:    "assistant",
				Content: config.Roles[role],
			},
		},
	}, nil
}

// SendGPTRequest sends GPT requests.
func (ggpt *GGPTAgent) SendGPTRequest(prompt string) (string, error) {
	userPrompt := message{
		Role:    ggpt.Config.GptRole,
		Content: prompt,
	}

	// Append the user question (context)
	ggpt.History = append(ggpt.History, userPrompt)

	requestData := GPTRequest{
		Model:     ggpt.Config.Model,
		Messages:  ggpt.History,
		MaxTokens: ggpt.Config.MaxTokens,
	}

	jsonData, err := json.Marshal(requestData)
	if err != nil {
		return "", err
	}

	req, err := http.NewRequest("POST", ggpt.Config.ApiURL, bytes.NewBuffer(jsonData))
	if err != nil {
		return "", err
	}

	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("Authorization", "Bearer "+ggpt.Config.ApiKey)

	client := http.DefaultClient

	resp, err := client.Do(req)
	if err != nil {
		return "", err
	}
	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return "", err
	}

	var gptResponse GPTResponse
	err = json.Unmarshal(body, &gptResponse)
	if err != nil {
		return "", err
	}

	response := gptResponse.Choices[0].Message.Content
	assistantResponse := message{
		Role:    "assistant",
		Content: response,
	}

	// Append the assistant response (context)
	ggpt.History = append(ggpt.History, assistantResponse)

	return gptResponse.Choices[0].Message.Content, nil
}
