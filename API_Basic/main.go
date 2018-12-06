/*
A basic API to display the content in a particular port given by USER
*/
package main

// Importing the needed packages for API
import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

// Declaring a structure for adding the content
type Article struct {
	Id      int    `json:"Id"`
	Title   string `json:"Title"`
	Desc    string `json:"desc"`
	Content string `json:"content"`
}

type Articles []Article

// Can write as many function dependson URL's and add different content to it for displaying.
// With respect to the URL these function is being called and displays the content
func homePage(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Welcome to the HomePage!")
	fmt.Fprintf(w, "Vipin Paul here!!!")
	fmt.Println("Endpoint Hit: homePage")
}

// This function displays all the content added to the structure
func returnAllArticles(w http.ResponseWriter, r *http.Request) {
	articles := Articles{
		Article{Id: 2, Title: "Hello", Desc: "Article Description", Content: "Article Content"},
		Article{Title: "Hello 2", Desc: "Article Description", Content: "Article Content"},
	}
	fmt.Println("Endpoint Hit: returnAllArticles")

	json.NewEncoder(w).Encode(articles)
}

// It displays the content given in URL in the place of "string" ("/article/string").
func returnSingleArticle(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	key := vars["string"]
	fmt.Fprintf(w, "Key: "+key)
}

func userPage(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "User")
	fmt.Fprintf(w, "Vipin Paul here!!!")
	fmt.Println("Endpoint Hit: userPage")
}

// Instead of main function,I have added all the code here, can be added in main too.
func handleRequests() {
	// mux.NewRouter checks incoming requests
	myRouter := mux.NewRouter().StrictSlash(true)

	// Each function is being called using different URL's.
	myRouter.HandleFunc("/", homePage)
	myRouter.HandleFunc("/user", userPage)
	myRouter.HandleFunc("/all", returnAllArticles)
	myRouter.HandleFunc("/article/{string}", returnSingleArticle)

	// Port number is being addressed here, can be changed. then open http://localhost:8081
	log.Fatal(http.ListenAndServe(":8081", myRouter))
}

func main() {
	// Calling the function
	handleRequests()
}
