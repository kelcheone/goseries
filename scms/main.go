package main

import (
	"encoding/json"
	"fmt"
	"log"
	"math/rand"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
)

type Post struct {
	ID        string `json:"ID"`
	Title     string `json:"Title"`
	Content   string `json:"Content"`
	Published bool   `json:"Published"`
	Author    *User  `json:"Author"`
}

type User struct {
	Firstname string `json:"Firstname"`
	Lastname  string `json:"Lastname"`
}

var posts []Post

func root(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "It is on")
}
func getPosts(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(posts)
}

func deletePost(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	params := mux.Vars(r)

	for index, item := range posts {
		if item.ID == params["id"] {
			posts = append(posts[:index], posts[index+1:]...)
			json.NewEncoder(w).Encode(posts)
			break
		}
	}
}
func getPost(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	params := mux.Vars(r)

	for _, item := range posts {
		if item.ID == params["id"] {
			json.NewEncoder(w).Encode(item)
			return
		}
	}
}

func createPost(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Conten-Type", "application/json")
	var post Post
	_ = json.NewDecoder(r.Body).Decode(&post)
	post.ID = strconv.Itoa((rand.Intn(10000000)))

	posts = append(posts, post)

	json.NewEncoder(w).Encode(post)

}
func updatePost(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	var post Post
	_ = json.NewDecoder(r.Body).Decode(&post)
	params := mux.Vars(r)

	for index, item := range posts {
		if item.ID == params["id"] {
			posts = append(posts[:index], posts[index+1:]...)
			post.ID = params["id"]
			posts = append(posts, post)

			json.NewEncoder(w).Encode(post)
			return
		}
	}
}

func main() {

	r := mux.NewRouter()

	posts = append(posts, Post{ID: "1", Title: "My life", Content: "A day in the life of a digital marketer", Author: &User{Firstname: "Kevin", Lastname: "Miano"}})
	posts = append(posts, Post{ID: "2", Title: "My life", Content: "A day in the life of a digital marketer", Published: false, Author: &User{Firstname: "Kevin", Lastname: "Miano"}})
	posts = append(posts, Post{ID: "3", Title: "My life", Content: "A day in the life of a digital marketer", Published: false, Author: &User{Firstname: "Kevin", Lastname: "Miano"}})

	r.HandleFunc("/", root).Methods("GET")
	r.HandleFunc("/posts/{id}", deletePost).Methods("DELETE")
	r.HandleFunc("/posts", getPosts).Methods("GET")
	r.HandleFunc("/posts/{id}", getPost).Methods("GET")
	r.HandleFunc("/posts", createPost).Methods("POST")
	r.HandleFunc("/posts/{id}", updatePost).Methods("PUT")

	fmt.Println("Listening on port 8080")
	log.Fatal(http.ListenAndServe(":8080", r))
}
