package main

import (
	"encoding/json"
	"fmt"
	"net/http"
)

type Post struct {
	UserId int `json:"userId"`
	Id     int `json:"id"`
	Title  string `json:"title"`
	Body   string `json:"body"`
}

func main() {
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