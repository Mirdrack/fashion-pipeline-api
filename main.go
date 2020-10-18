package main

import (
    "fmt"
    "log"
    "net/http"
    "encoding/json"

    "github.com/gorilla/mux"
)

type quote struct {
    ID string `json:"ID"`
    Author string `json:"Author"`
    Content string `json:"Content"`
}

type allQuotes []quote

var quotes = allQuotes{
    {
        ID: "1",
        Author: "Karl Lagerfeld",
        Content: "One is never over-dressed or under-dressed with a Little Black Dress.",
    },
    {
        ID: "2",
        Author: "Bette Midler",
        Content: "I firmly believe that with the right footwear one can rule the world.",
    },
}


func homeLink(w http.ResponseWriter, r *http.Request) {
    fmt.Fprintf(w, "Welcome to the Fashion Pipeline API!\n")
}

func getOneQuote(w http.ResponseWriter, r *http.Request) {
    quoteID := mux.Vars(r)["id"]
    for _, singleQuote := range quotes {
        if singleQuote.ID == quoteID {
            err := json.NewEncoder(w).Encode(singleQuote)
            if err != nil {
                fmt.Printf("There was an error. \n")
            }
        }
    }
}

func getAllQuotes(w http.ResponseWriter, r *http.Request) {
    err := json.NewEncoder(w).Encode(quotes)
    if err != nil {
        fmt.Printf("There was an error. \n")
    }
}

func main() {
    router := mux.NewRouter().StrictSlash(true)
    router.HandleFunc("/", homeLink)
    router.HandleFunc("/quotes", getAllQuotes).Methods("GET")
    router.HandleFunc("/quotes/{id}", getOneQuote).Methods("GET")
    log.Fatal(http.ListenAndServe(":8000", router))
}
