package main

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"strings"
)

type Post struct {
	UserId int `json:"userId"`
	Id     int `json:"id"`
	Title  string `json:"title"`
	Body   string `json:"body"`
}

func getRequest() {
	res, err := http.Get("https://jsonplaceholder.typicode.com/posts/2")
	if err != nil {
		fmt.Println(err)
		return

	}
	defer res.Body.Close()

	if res.StatusCode != http.StatusOK {
		fmt.Println("Error: ", res.Status)
		return
	}

	// data, err := io.ReadAll(res.Body) 
	// if err != nil {
	// 	fmt.Println(err)
	// 	return
	// }
	// fmt.Println(string(data))

	var post Post
	err = json.NewDecoder(res.Body).Decode(&post)
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println(post)
}


func postRequest() {
	 post := Post{1, 1, "Title", "Body"}
	 jsonData, err := json.Marshal(post)
	 if err != nil {
		 fmt.Println(err)
		 return
	 }
	 stringData := string(jsonData)
	 jsonReader:=strings.NewReader(stringData)
	 newUrl := "https://jsonplaceholder.typicode.com/posts"

	 http.Post(newUrl, "application/json", jsonReader)
	 res, err := http.Post(newUrl, "application/json", jsonReader)
	 if err != nil {
		 fmt.Println(err)
		 return
	 }
	 defer res.Body.Close()

	 data,_:=io.ReadAll(res.Body)
	 fmt.Println(string(data))
}

func main() {
	// getRequest()
	postRequest()
}