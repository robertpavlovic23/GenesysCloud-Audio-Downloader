package main

import (
	"GCAudioDownloader/auth"
	"GCAudioDownloader/handlers"
	"encoding/csv"
	"fmt"
	"io"
	"log"
	"net/http"
	"os"
	"regexp"

	"github.com/joho/godotenv"
)

func main() {
	if err := godotenv.Load(); err != nil {
		fmt.Println("Error loading environment variables:", err)
		//return err
	}

	tokenResponse, err := auth.GetAccessToken(os.Getenv("CLIENT_ID"), os.Getenv("CLIENT_SECRET"))
	if err != nil {
		fmt.Println("Error getting access token:", err)
		return
	}

	fmt.Println("Access Token:", tokenResponse.AccessToken)
	fmt.Println("Token Type:", tokenResponse.TokenType)
	fmt.Println("Expires In:", tokenResponse.ExpiresIn)

	getPromptResponse, err := handlers.GetPrompt(tokenResponse.AccessToken)
	if err != nil {
		fmt.Println("Error getting prompt response:", err)
		return
	}

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

			createAudioFile(resources[resource].MediaURI, getPromptResponse[json].Name, resources[resource].Language)

			//filename := createAudioFile(resources[resource].MediaURI, getPromptResponse[json].Name, resources[resource].Language)
			//err := fillCSVFile(getPromptResponse[json].Name, getPromptResponse[json].Description, filename, resources[resource].TTSString)
			// if err != nil {
			// 	fmt.Println("Error occurred while filling CSV: |", err, " |")
			// 	continue
			// }
		}
		fmt.Println("------------------")
	}
}

func createAudioFile(url, name, language string) string {
	directory := "audio/"
	re := regexp.MustCompile(`\/([^\/]+\.wav)`)
	match := re.FindStringSubmatch(url)

	var fileName string

	if len(match) > 1 {
		fileName = directory + name + "_" + language + ".wav"
		fmt.Println("Extracted filename:", fileName)
	} else {
		fmt.Println("Filename not found in the URL")
		return ""
	}

	fmt.Println("CHECK:", fileName)

	// Send the GET request
	resp, err := http.Get(url)
	if err != nil {
		fmt.Println("Error while downloading:", err)
		return ""
	}
	defer resp.Body.Close()

	// Create the file
	out, err := os.Create(fileName)
	if err != nil {
		fmt.Println("Error while creating file:", err)
		return ""
	}
	defer out.Close()

	// Copy the response body to the file
	_, err = io.Copy(out, resp.Body)
	if err != nil {
		fmt.Println("Error while saving file:", err)
		return ""
	}

	fmt.Println("File downloaded successfully")
	return fileName
}

func fillCSVFile(name, description, audio, tts string) error {
	// inputFile, err := os.Open("prompts.csv")
	file, err := os.OpenFile("prompts.csv", os.O_CREATE|os.O_WRONLY|os.O_APPEND, 0644)
	if err != nil {
		log.Fatal("Failed to open CSV file: ", err)
	}
	defer file.Close()

	writer := csv.NewWriter(file)
	defer writer.Flush()

	writer.Comma = ';'

	record := []string{name, description, audio, tts}
	if err := writer.Write(record); err != nil {
		return fmt.Errorf("failed to write record: %v", err)
	}

	return nil
}
