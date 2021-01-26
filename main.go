package main


import (
	"fmt"
	"log"
	"net/http"
	"encoding/json"
)
type Article struct{
	Title string `json:"Title"`
	Desc string `json:"desc"`
	Content string `json:"content"`
}

// let's declare a global Articles array
// that we can then populate in our main function
// to simulate a database
var Articles []Article


func homePage(w http.ResponseWriter, r *http.Request){
	fmt.Fprintf(w,"welcome to homepage")
	fmt.Println("Welcome to homepage")
}

func handleRequests() {
	http.HandleFunc("/", homePage)
	http.HandleFunc("/articles", returnAllArticles)

    log.Fatal(http.ListenAndServe(":8000", nil))
}
func returnAllArticles(w http.ResponseWriter, r *http.Request){

	fmt.Println("Endpoint Hit: returnAllArticles")
	json.NewEncoder(w).Encode(Articles)
}




func main() {
	Articles = []Article{
        Article{Title: "Article1", Desc: "This is an article ablout games", Content: "Article one is good all about all every an "},
        Article{Title: "Hello 2", Desc: "This is an article about movies", Content: "Article content is very good is an "},
    }
    handleRequests()
}
