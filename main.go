package main

import (
	"GCAudioDownloader/auth"
	"GCAudioDownloader/handlers"
	"GCAudioDownloader/utility"
	"fmt"
	"os"

	"github.com/joho/godotenv"
)

func main() {
	if err := godotenv.Load(); err != nil {
		fmt.Println("Error loading environment variables:", err)
		//return err
	}

	// Create token
	tokenResponse, err := auth.GetAccessToken(os.Getenv("CLIENT_ID"), os.Getenv("CLIENT_SECRET"))
	if err != nil {
		fmt.Println("Error getting access token:", err)
		return
	}

	// Print auth info
	fmt.Println("Access Token:", tokenResponse.AccessToken)
	fmt.Println("Token Type:", tokenResponse.TokenType)
	fmt.Println("Expires In:", tokenResponse.ExpiresIn)

	// Get API data
	getPromptResponse, err := handlers.GetPrompt(tokenResponse.AccessToken)
	if err != nil {
		fmt.Println("Error getting prompt response:", err)
		return
	}

	// Going through each entry and performing download/fillCSV
	for json := range getPromptResponse {
		resources, err := handlers.ExtractFields(getPromptResponse[json].Resources)
		if err != nil {
			fmt.Printf("\nFor json number: [%v] , an error occurred: %v\n", json, err)
		}

		fmt.Printf("\n[Entity]:[%v]\n[name: %v]", json, getPromptResponse[json].Name)
		fmt.Println("\n[description:]", getPromptResponse[json].Description)
		for resource := range resources {
			fmt.Println("\n[language:]\n", resources[resource].Language)
			fmt.Println("\n[mediaURI:]\n", resources[resource].MediaURI)
			fmt.Println("\n[ttsString:]\n", resources[resource].TTSString)

			// utility.CreateAudioFile(resources[resource].MediaURI, getPromptResponse[json].Name, resources[resource].Language)

			filename := utility.CreateAudioFile(resources[resource].MediaURI, getPromptResponse[json].Name, resources[resource].Language)
			err := utility.FillCSVFile(getPromptResponse[json].Name, getPromptResponse[json].Description, filename, resources[resource].TTSString)
			if err != nil {
				fmt.Println("Error occurred while filling CSV: |", err, " |")
				continue
			}
		}
		fmt.Println("------------------")
	}
}
