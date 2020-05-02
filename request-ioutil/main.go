package main

import (
	"encoding/json"

	"io/ioutil"
	"net/http"
)

type Article struct {
	Title string `json:"title"`
	Desc  string `json:"desc"`
}

type Articles []Article

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
		// ========menangani request post dgn "io/ioutil"================
		body, err := ioutil.ReadAll(r.Body)
		if err != nil {
			http.Error(w, "Can't read body", http.StatusInternalServerError)
		}
		w.Write([]byte(string(body)))

	} else {
		http.Error(w, "Invalid request method", http.StatusMethodNotAllowed)
	}
}
