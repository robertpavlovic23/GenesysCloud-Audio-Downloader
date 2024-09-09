package main

import (
	"fmt"
	"io"
	"net/http"
	"os"
	"regexp"
)

func main() {
	// The URL of the audio file
	url := "https://prod-usw2-audio-prompts.s3.us-west-2.amazonaws.com/organizations/ece448d4-7800-4ed0-a6bd-b3061f0db22c/00963898-e532-42dc-940a-5937b7cb352c/en-us/6f81c31f-4f77-43be-ae39-770d947e1f6a.wav?X-Amz-Security-Token=IQoJb3JpZ2luX2VjEGwaCXVzLXdlc3QtMiJHMEUCID%2Fv2x9dVCarkNXKOX7iUBXKM%2F9oNoNKCqohNQkZlRN2AiEAn4xfDVt4j0mZkWyvIA7t%2B%2BPxBO29XnU5SZ8BdDb%2BtkgqswUIdRAAGgw3NjU2Mjg5ODU0NzEiDHWuxPUeULIT8o8WPyqQBcfFUVgA6%2BHimbFs8edXyIMD6%2BXVUqwA%2BXutThkG7AHIF7GFbZ8PpC%2FnnavnAoPpmxQazKkpA1XHMoGq0hn3khhnB4u%2BSXpJ7nihfONCCBz%2BEGU5TDTO87NW%2BbnJ6SndrfcLTGwYaog9j1N38EF%2BoAJqWzYs8PgUEFEsebP91NELgSwHoWthgGXrrxWF6CYKj626JyGNlIBh46GZJL54DTjX01gmOf6r3%2BZf8iTF0KmaLi5tSRgcy64mgH1JFIgy1Ij%2BZEiob5whly7Sfj6%2FXRAsAoQ7sYHt2G8eod3ih2ufGp2utG6cTz8ycVyTJW3xwjlE%2FIzxE118WNzDQuxxzml9oBvdarFetw9lzPFheENF4LXfMh5xswj044UurY0hFbahu51VQt2R5R%2Fxg9k0qgMV%2F1aqBN56BjK06HYQDj9r4CJD89Hxc2dqBRIfvERsqIPqe4iY18Z0nF5xrKwFXN6l3bHRHRGRDq0XDWEQuXVrdw%2BnFchmUw1DlDquXRV82CYV7n9Fx77WN3ROOqlsT1b7CUzyde1h2huulOVc8GrVd378bQpagcalXSD9C35B9ahR0V5xANXSEfmd7xSnkJgvCss80b6WyigbGXKy6XYX6mpuWwMQERdNZwOd2y9Ztnnbt5IF%2FL%2Feo1djHPJpjkBtzL%2BgYI%2F7J4RThov0Csmgzqdu%2FdqAUiS%2FFkHo1PjkBDe4F4fsCvN%2F%2Bg1eYeab6SOxdAsYOYlvdOggrYoDHRWRnPQ0LliB79vy07PFnWJkT8GFOfgARnH6Hd%2BxCppGC2EWBk8PF5Wm5QXH%2FqBOePO8yhW29x%2B57hPrZT3hz9AjNzyhkZE%2FBiFyyAjMgOFEfCoqiQuATPT292NBy%2F1dHXCGMMifl7YGOrEBMrgclmlqYcgYzxHm8NPfgaEUJBtlYFHMQ2cPndkOZbXsGfr7h1BAoRUHn89oo8s1EPmlegQrKUN6Fq0cadpcUwsPdbdw3SI4tqvIC9zogK6NN22zfm6XycYtGcOYSjpmQ7skQXof3sdbaLC766wtQW79ScbK9spViXjq5p3kdez3ONeSdPhVRuB4cI4WF0AMMfNxnMGE76VuW1HMb42uGT8ZXmn%2BxTWdkPmuP1C%2FF2Sp&X-Amz-Algorithm=AWS4-HMAC-SHA256&X-Amz-Date=20240821T135341Z&X-Amz-SignedHeaders=host&X-Amz-Credential=ASIA3EQYLGB7ZRXZTZJJ%2F20240821%2Fus-west-2%2Fs3%2Faws4_request&X-Amz-Expires=3600&X-Amz-Signature=1e01a3f2b61e2dfefbd0931763140ac496740e9bea8a3f841b64011d612b2a20"

	re := regexp.MustCompile(`\/([^\/]+\.wav)`)
	match := re.FindStringSubmatch(url)

	var fileName string

	if len(match) > 1 {
		fileName = match[1]
		fmt.Println("Extracted filename:", fileName)
	} else {
		fmt.Println("Filename not found in the URL")
		return
	}

	fmt.Println("CHECK:", fileName)

	// Send the GET request
	resp, err := http.Get(url)
	if err != nil {
		fmt.Println("Error while downloading:", err)
		return
	}
	defer resp.Body.Close()

	// Create the file
	out, err := os.Create(fileName)
	if err != nil {
		fmt.Println("Error while creating file:", err)
		return
	}
	defer out.Close()

	// Copy the response body to the file
	_, err = io.Copy(out, resp.Body)
	if err != nil {
		fmt.Println("Error while saving file:", err)
		return
	}

	fmt.Println("File downloaded successfully")
}
