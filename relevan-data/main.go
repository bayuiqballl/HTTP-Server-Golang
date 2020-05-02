package main

import (
	"encoding/json"
	"fmt"

	// "io/ioutil"
	"net/http"
)

type Article struct {
	Title string `json:"title"`
	Desc  string `json:"desc"`
}

type Articles []Article

var articles = Articles{
	Article{Title: "judul pertama", Desc: "deskripsi pertama"},
	Article{Title: "judul kedua", Desc: "deskripsi kedua"},
}

func main() {
	http.HandleFunc("/", getHome)
	http.HandleFunc("/articles", getArticles)
	http.HandleFunc("/post-articles", postArticle)
	http.ListenAndServe(":3000", nil)
}

func getHome(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Testing Home Page"))
}

func getArticles(w http.ResponseWriter, r *http.Request) {
	article := Articles{
		Article{Title: "judul pertama", Desc: "deskripsi pertama"},
		Article{Title: "judul kedua", Desc: "deskripsi kedua"},
	}

	json.NewEncoder(w).Encode(article)
}

func postArticle(w http.ResponseWriter, r *http.Request) {
	// panic(r.Method)
	if r.Method == "POST" {
		// // ========menangani request post dgn decoder json================
		var newArticle Article
		err := json.NewDecoder(r.Body).Decode(&newArticle)

		if err != nil {
			fmt.Println("Ada Error", err)
		}
		articles = append(articles, newArticle)
		json.NewEncoder(w).Encode(articles)

	} else {
		http.Error(w, "Invalid request method", http.StatusMethodNotAllowed)
	}
}
