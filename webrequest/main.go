package main

import (
	"fmt"
	"io"
	"log"
	"net/http"
)

func main() {
	// url := "https://jsonplaceholder.typicode.com/posts/1"

	// response, err := http.Get(url)
	// if err != nil {
	// 	log.Fatalf("Error occurred while making the GET request: %v", err)
	// }
	// defer response.Body.Close() 

	// body, err := io.ReadAll(response.Body)
	// if err != nil {
	// 	log.Fatalf("Error reading the response body: %v", err)
	// }

	// fmt.Println("Response from the server:")
	// fmt.Println(string(body))


	// Create a new HTTP client
	client := &http.Client{}

	req, err := http.NewRequest("GET", "https://jsonplaceholder.typicode.com/posts/1", nil)
	if err != nil {
		log.Fatalf("Error creating the request: %v", err)
	}

	req.Header.Add("Authorization", "Bearer your-token-here")
	req.Header.Add("Content-Type", "application/json")

	resp, err := client.Do(req)
	if err != nil {
		log.Fatalf("Error making the request: %v", err)
	}
	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		log.Fatalf("Error reading the response body: %v", err)
	}

	fmt.Println("Response Status:", resp.Status)
	fmt.Println("Response Body:", string(body))
}
