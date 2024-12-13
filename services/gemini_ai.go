package services

import (
	"ametory-crud/config"
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"os"

	"github.com/google/generative-ai-go/genai"
	"google.golang.org/api/option"
)

func GeminiPrompt(input string, parts ...genai.Part) (string, error) {
	ctx := context.Background()

	apiKey := config.App.Google.GeminiApiKey
	if apiKey == "" {
		return "", errors.New("Environment variable GEMINI_API_KEY not set")
	}

	client, err := genai.NewClient(ctx, option.WithAPIKey(apiKey))
	if err != nil {
		return "", errors.New("Error creating client: " + err.Error()) //log.Fatalf("Error creating client: %v", err)
	}
	defer client.Close()

	model := client.GenerativeModel("gemini-1.5-flash-8b")

	model.SetTemperature(1)
	model.SetTopK(40)
	model.SetTopP(0.95)
	model.SetMaxOutputTokens(8192)
	model.ResponseMIMEType = "text/plain"
	if config.App.Google.GeminiResponseMIMEType != "" {
		model.ResponseMIMEType = config.App.Google.GeminiResponseMIMEType
	}

	// model.SafetySettings = []*genai.SafetySetting{
	// 	{
	// 		Category:  genai.HarmCategoryHarassment,
	// 		Threshold: genai.HarmBlockNone,
	// 	},
	// 	{
	// 		Category:  genai.HarmCategoryHateSpeech,
	// 		Threshold: genai.HarmBlockNone,
	// 	},
	// }
	model.ResponseMIMEType = "text/plain"
	if config.App.Google.GeminiSystemInstruction != "" {
		model.SystemInstruction = genai.NewUserContent(genai.Text(config.App.Google.GeminiSystemInstruction))
	}

	session := model.StartChat()

	session.History = []*genai.Content{}
	if config.App.Google.GeminiHistoryFile != "" {
		historyFile, err := os.Open(config.App.Google.GeminiHistoryFile)
		if err != nil {
			return "", errors.New("Error opening history file: " + err.Error())
		}
		defer historyFile.Close()

		historyDecoder := json.NewDecoder(historyFile)
		var historyData []GeminiTrainingData
		if err := historyDecoder.Decode(&historyData); err != nil {
			return "", errors.New("Error decoding history: " + err.Error())
		}

		for _, v := range historyData {
			history := genai.Content{
				Role: v.Role,
				Parts: []genai.Part{
					genai.Text(v.Parts[0].Text),
				},
			}
			session.History = append(session.History, &history)
		}
	}

	tokResp, err := model.CountTokens(ctx, genai.Text(input))
	if err != nil {
		return "", errors.New("Error sending message: " + err.Error())
	}

	fmt.Println("total_tokens:", tokResp.TotalTokens)

	geminiInput := append([]genai.Part{genai.Text(input)}, parts...)
	resp, err := session.SendMessage(ctx, geminiInput...)
	if err != nil {
		return "", errors.New("Error sending message: " + err.Error())
	}

	for _, part := range resp.Candidates[0].Content.Parts {
		fmt.Println("total response token :", resp.UsageMetadata.TotalTokenCount)
		return fmt.Sprintf("%v\n", part), nil

	}

	return "", nil
}

type GeminiTrainingData struct {
	Role  string `json:"role"`
	Parts []struct {
		Text string `json:"text"`
	} `json:"parts"`
}

func (g *GeminiTrainingData) FromJson(data []byte) error {
	return json.Unmarshal(data, g)
}

func (g *GeminiTrainingData) ToJson() ([]byte, error) {
	return json.Marshal(g)
}
