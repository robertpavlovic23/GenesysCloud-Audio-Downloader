package utility

import (
	"encoding/csv"
	"fmt"
	"io"
	"log"
	"net/http"
	"os"
	"regexp"
)

func CreateAudioFile(url, name, language string) string {
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

func FillCSVFile(name, description, audio, tts string) error {
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
