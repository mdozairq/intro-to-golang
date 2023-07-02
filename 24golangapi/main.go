package main

import (
	"encoding/json"
	"fmt"
	"log"
	"math/rand"
	"net/http"
	"strconv"
	"time"

	"github.com/gorilla/mux"
)

//model for post - file

type Post struct {
	PostId string   `json:"postId"`
	User   *User    `json:"user"`
	Title  string   `json:"title"`
	Body   string   `json:"body"`
	Tags   []string `json:"tags,omitempty"`
}

// model for user
type User struct {
	UserId   string `json:"userId"`
	FullName string `json:"fullname"`
	Phone    string `json:"phone"`
	Email    string `json:"email"`
	Password string `json:"password"`
}

// fake DB
var posts []Post
var users []User

// middleware, helper
func (p *Post) IsEmpty() bool {
	return p.Title == "" && p.User.UserId == ""
}

func main() {

	r := mux.NewRouter()

	//seeding
	users = append(users, User{UserId: "1230987", FullName: "Ozair", Phone: "9876554321", Email: "mdozairq@gmail.com", Password: "Ozair123"})
	users = append(users, User{UserId: "10345348", FullName: "Yogesh", Phone: "9876543210", Email: "yogi@gmail.com", Password: "Yogi123"})

	posts = append(posts, Post{PostId: "7857353", User: &users[0], Title: "Blog site using GoLang", Body: "Lorem Ipsum is simply dummy text of the printing and typesetting industry. Lorem Ipsum has been the industry's standard dummy text ever since the 1500s, when an unknown printer took a galley of type and scrambled it to make a type specimen book. It has survived not only five centuries, but also the leap into electronic typesetting, remaining essentially unchanged. It was popularised in the 1960s with the release of Letraset sheets containing Lorem Ipsum passages, and more recently with desktop publishing software like Aldus PageMaker including versions of Lorem Ipsum.", Tags: []string{"GoLang", "Stuff", "Fun"}})

	posts = append(posts, Post{PostId: "9090053", User: &users[0], Title: "Learn GoLang", Body: "Lorem Ipsum is simply dummy text of the printing and typesetting industry. Lorem Ipsum has been the industry's standard dummy text ever since the 1500s, when an unknown printer took a galley of type and scrambled it to make a type specimen book. It has survived not only five centuries, but also the leap into electronic typesetting, remaining essentially unchanged. It was popularised in the 1960s with the release of Letraset sheets containing Lorem Ipsum passages, and more recently with desktop publishing software like Aldus PageMaker including versions of Lorem Ipsum.", Tags: []string{"GoLang", "New", "Start"}})

	//router
	r.HandleFunc("/", serveHome).Methods("GET")
	r.HandleFunc("/posts", getAllPosts).Methods("GET")
	r.HandleFunc("/post/{id}", getPostById).Methods("GET")
	r.HandleFunc("/post", addNewPost).Methods("POST")
	r.HandleFunc("/post/{id}", updatePostById).Methods("PUT")
	r.HandleFunc("/post/{id}", deletePostById).Methods("DELETE")

	//listen to a port
	log.Fatal(http.ListenAndServe(":4000", r))


}

//controller -file

// serve home route
func serveHome(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("<h1>Welcome to Golang BlogSite Backend</h1>"))
}

func getAllPosts(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Get All Posts")
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(posts)
}

func getPostById(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Get A Post by PostId")
	w.Header().Set("Content-Type", "application/json")

	//getId from request
	params := mux.Vars(r)

	//filter the post by PostId from all posts
	for _, post := range posts {
		if post.PostId == params["id"] {
			json.NewEncoder(w).Encode(post)
			return
		}
	}

	json.NewEncoder(w).Encode("No Post found with given Id")
	return
}

func addNewPost(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Create a User Post to DB")
	w.Header().Set("Content-Type", "application/json")

	//check Empty
	if r.Body == nil {
		json.NewEncoder(w).Encode("No Post Data recieved in Request Body")
	}

	var post Post
	_ = json.NewDecoder(r.Body).Decode(&post)
	if post.IsEmpty() {
		json.NewEncoder(w).Encode("Post Title or UserId is required")
		return
	}

	//generate Unique PostId and add the Post in DB
	rand.Seed(time.Now().UnixNano())
	post.PostId = strconv.Itoa(rand.Intn(100))
	posts = append(posts, post)
	json.NewEncoder(w).Encode(post)
	return
}

// update a post
func updatePostById(w http.ResponseWriter, r *http.Request) {
	fmt.Println("UPDATE A Post by PostId")
	w.Header().Set("Content-Type", "application/json")

	//getId from request
	params := mux.Vars(r)

	//filter the post by PostId from all posts first removove it and update it
	for index, post := range posts {
		if post.PostId == params["id"] {
			posts = append(posts[:index], posts[index+1:]...)
			var updatedPost Post
			_ = json.NewDecoder(r.Body).Decode(&updatedPost)
			if updatedPost.IsEmpty() {
				json.NewEncoder(w).Encode("Post Title or UserId is required")
				return
			}

			updatedPost.PostId = params["id"]
			fmt.Println(updatedPost)
			posts = append(posts, post)
			json.NewEncoder(w).Encode(post)
			return
		}
	}

	json.NewEncoder(w).Encode("No Post found with given Id")
	return
}

// Delete a post by Id
func deletePostById(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Get A Post by PostId")
	w.Header().Set("Content-Type", "application/json")

	//getId from request
	params := mux.Vars(r)

	//filter the post by PostId from all posts first removove it and update it
	for index, post := range posts {
		if post.PostId == params["id"] {
			posts = append(posts[:index], posts[index+1:]...)
			json.NewEncoder(w).Encode("Post deleted by Id Successfully")
			return
		}
	}

	json.NewEncoder(w).Encode("No Post found with given Id")
	return
}
