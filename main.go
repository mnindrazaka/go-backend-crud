package main

import (
	"encoding/json"
	"net/http"
	"strconv"
)

type Article struct {
	Title   string `json:"title"`
	Content string `json:"content"`
}

var Articles = []Article{
	{Title: "Hello", Content: "Content of hello"},
	{Title: "Hello 2", Content: "Content of hello 2"},
}

func handleHome(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("welcome !"))
}

func handleArticles(w http.ResponseWriter, r *http.Request) {
	var queryIndex = r.URL.Query().Get("index")

	if queryIndex != "" {
		var index, err = strconv.ParseInt(queryIndex, 10, 32)
		if err != nil {
			w.Write([]byte("something went wrong"))
		} else {
			json.NewEncoder(w).Encode(Articles[index])
		}
	} else {
		json.NewEncoder(w).Encode(Articles)
	}
}

func main() {
	http.HandleFunc("/", handleHome)
	http.HandleFunc("/articles", handleArticles)
	http.ListenAndServe(":3000", nil)
}
