package main

import (
    "os"
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
        Author: "Diana Vreeland",
        Content: "Fashion is part of the daily air and it changes all the time, with all the events. You can even see the approaching of a revolution in clothes. You can see and feel everything in clothes.",
    },
    {
        ID: "2",
        Author: "Karl Lagerfeld",
        Content: "One is never over-dressed or under-dressed with a Little Black Dress.",
    },
    {
        ID: "3",
        Author: "Miuccia Prada",
        Content: "What you wear is how you present yourself to the world, especially today, when human contacts are so quick. Fashion is instant language.",
    },
    {
        ID: "4",
        Author: "Bette Midler",
        Content: "I firmly believe that with the right footwear one can rule the world.",
    },
    {
        ID: "5",
        Author: "Gianni Versace",
        Content: "Don't be into trends. Don't make fashion own you, but you decide what you are, what you want to express by the way you dress and the way to live.",
    },
}

func determineListenAddress() (string, error) {
  port := os.Getenv("PORT")
  if port == "" {
    return "", fmt.Errorf("$PORT not set")
  }
  return ":" + port, nil
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
    addr, err := determineListenAddress()
    if err != nil {
        log.Fatal(err)
    }
    router := mux.NewRouter().StrictSlash(true)
    router.HandleFunc("/", homeLink)
    router.HandleFunc("/quotes", getAllQuotes).Methods("GET")
    router.HandleFunc("/quotes/{id}", getOneQuote).Methods("GET")
    log.Fatal(http.ListenAndServe(addr, router))
}
