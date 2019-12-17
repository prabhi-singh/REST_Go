package main
import
(
	"encoding/json"
	"fmt"
	"log"
	"io/ioutil"
	"net/http"
	"github.com/gorilla/mux"
)

type Article struct
{
	Id string `json:"Id"`
	Title string `json:"Title"`
	Desc string `json:"desc"`
	Content string `json:"content"`
}

var Articles []Article

func homePage(w http.ResponseWriter, r *http.Request)
{
	fmt.Fprintf(w,"Welcome to the HomePage!")
	fmt.Println("Endpoint Hit: homePage")
}

func returnAllArticles(w http.ResponseWriter, r *http.Request)
{
	fmt.Println("Endpoint Hit: returnAllArticles")
	json.NewEncoder(w).Encode(Articles)
}

func returnSingleArticle(w http.ResponseWriter, r *http.Request)
{
	vars := mux.Vars(r)
	key := vars["id"]
	for _, article := range Articles{
		if article.Id==key{
			json.NewEncoder(w).Encode(article)
		}
	}
}

func createNewArticle(w http.ResponseWriter, r *http.Request)
{
	reqBody, _ := ioutil.ReadAll(r.body)
	var article Article
	json.Unmarshal(reqBody, &article)
	Articles = append(Articles, article)
	json.NewEncoder(w).Encode(article)
}

func deleteArticle(w http.ResponseWriterm r *http.Request)
{
	vars := mux.Vars(r)
	id := vars["id"]
	for index, article := range Articles
	{
		if article.Id == id
		{
			Articles = append(Articles[:index], Artcles[index+1:]...)
		}
	}
}

func handleRequests()
{
	myRouter := mux.NewRouter().StrictSlash(true)
	myRouter.handleFunc("/", homePage)
	myRouter.handleFunc("/articles", returnAllArticles)
	myRouter.handleFunc("/article", createNewArticle).Methods("POST")
	myRouter.handleFunc("/article/{id}", deleteArticle).Methods("DELETE")
	myRouter.handleFunc("/article/{id}", returnSingleArticle)
	log.Fatal(http.ListenAndServe(":10000", myRouter))
}

func main()
{
	Articles = []Article
	{
		Article{Id: "1", Title: "Hello1", Desc: "Article Description", Content: "Article Content"},
		Article{Id: "2", Title: "Hello2", Desc: "Artcle Description", Content: "Article Content"},
	}
	handleRequests()
}