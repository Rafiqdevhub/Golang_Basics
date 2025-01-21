package main

import (
	"fmt"
	"net/url"
)

func main() {
	
	// rawURL := "https://example.com:8080/path?query=golang&lang=en#fragment"

	// parsedURL, err := url.Parse(rawURL)
	// if err != nil {
	// 	log.Fatalf("Error parsing the URL: %v", err)
	// }

	// fmt.Println("Scheme:", parsedURL.Scheme)
	// fmt.Println("Host:", parsedURL.Host)
	// fmt.Println("Path:", parsedURL.Path)
	// fmt.Println("Query:", parsedURL.RawQuery)
	// fmt.Println("Fragment:", parsedURL.Fragment)

	// queryParams := parsedURL.Query()
	// fmt.Println("Query Parameters:")
	// for key, values := range queryParams {
	// 	fmt.Printf("  %s: %v\n", key, values)
	// }


	
	baseURL := &url.URL{
		Scheme: "https",
		Host:   "example.com",
		Path:   "search",
	}

	params := url.Values{}
	params.Add("query", "golang")
	params.Add("page", "1")

	baseURL.RawQuery = params.Encode()

	fmt.Println("Constructed URL:", baseURL.String())
}
