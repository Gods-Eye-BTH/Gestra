package main

import (
    "fmt"
    "log"
    "strconv"
    "net/http"
    "encoding/json"
    "github.com/gorilla/mux"
)

//port to run on
var port = 8080

//Article is a json object
type Article struct {
    ID int `json:"id"`
    Title string `json:"Title"`
    Desc string `json:"desc"`
    Content string `json:"content"`
}

//Articles is a array of the type Article
type Articles []Article

//Json objects go here
var articles = Articles {
    Article{
        ID: 0,
        Title:"Test Title",
        Desc:"a test article",
        Content:"Hello json world",
    },
    Article{
        ID: 1,
        Title: "Hello 2",
        Desc: "Another Article Description",
        Content: "Article Content",
    },
    Article{
        ID: 2,
        Title: "Part 2",
        Desc: "More Data",
        Content: "and even more comments",
    },
}

//index route
func index(w http.ResponseWriter, r *http.Request) {
    fmt.Fprintf(w,
        `<h1>GESTRA API</h1>
        <p>
            Welcome to the <b>Gestra api</b>. use the route <code>/all</code>
            for all the items or <code>/id/{id}</code> for a specific item
        </p>`)
}

//display all objects route
func allArticles(w http.ResponseWriter, r *http.Request) {
    w.Header().Set("Content-Type", "application/json")
    json.NewEncoder(w).Encode(articles)
}

//return a spesific object route
func returnArticleByID(w http.ResponseWriter, r *http.Request) {
    vars := mux.Vars(r)
    key := vars["id"]
    convd, err := strconv.Atoi(key)
    if (err == nil) {
        if (convd > len(articles)-1) {
            w.WriteHeader(http.StatusBadRequest)
            fmt.Fprint(w, "Out of range")
        } else {
            w.Header().Set("Content-Type", "application/json")
            w.WriteHeader(http.StatusOK)
            json.NewEncoder(w).Encode(articles[convd])
        }
    } else {
        fmt.Fprintf(w, "parse error on key: " + key)
    }
}

func requestHandler()  {
    muxRouter := mux.NewRouter().StrictSlash(true)
    muxRouter.HandleFunc("/", index)
    muxRouter.HandleFunc("/all", allArticles)
    muxRouter.HandleFunc("/id/{id}", returnArticleByID)
    log.Fatal(http.ListenAndServe(":" + strconv.Itoa(port), muxRouter))
}

func main() {
    fmt.Println("GESTRA is now running\nport:" + strconv.Itoa(port) + "\n---")
    requestHandler()
}
