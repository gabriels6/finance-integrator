package main

import (
	"context"
	"io"
	"os"

	"google.golang.org/genai"
)

var client *genai.Client

func EnableClient() error {

	currClient, err := genai.NewClient(context.Background(), &genai.ClientConfig{
		APIKey:  os.Getenv("GEMINI_API_KEY"),
		Backend: genai.BackendGeminiAPI,
	})
	if err != nil {
		return err
	}
	client = currClient
	return nil
}

func AskForResponse(contents []*genai.Content) (string, error) {
	response, err := client.Models.GenerateContent(context.Background(), "gemini-2.0-flash-001", contents, &genai.GenerateContentConfig{})
	if err != nil {
		return "", err
	}
	return response.Text(), nil
}

func UploadFile(reader io.Reader, mimeType string) (*genai.File, error) {
	aiFile, err := client.Files.Upload(context.Background(), reader, &genai.UploadFileConfig{
		MIMEType: mimeType,
	})
	if err != nil {
		return nil, err
	}
	return aiFile, nil
}
