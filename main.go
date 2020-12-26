package main

import (
    "io"
    "os"
    "fmt"
    "log"
    "math/rand"
    "net/http"
    "encoding/json"
    "encoding/csv"

    "github.com/gorilla/mux"
)

type quote struct {
    ID string `json:"ID"`
    Author string `json:"Author"`
    Content string `json:"Content"`
}

type allQuotes []quote

var quotes = loadQuotes()

func loadQuotes() allQuotes {
    quotes := allQuotes{}
    csvFile, _ := os.Open("quotes.csv")
    reader := csv.NewReader(csvFile)
    reader.Comma = ','
    reader.LazyQuotes = true

    for {
        line, err := reader.Read()
        if err == io.EOF {
            break
        } else if err != nil {
            log.Fatal(err)
        }
        quotes = append(quotes, quote{
              ID: line[0],
              Author: line[1],
              Content: line[2],
        })
    }
    return quotes
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

func getRandomQuote(w http.ResponseWriter, r *http.Request) {
    randomQuote := quotes[rand.Intn(len(quotes))]
    err := json.NewEncoder(w).Encode(randomQuote)
    if err != nil {
        fmt.Printf("There was an error. \n")
    }
}

func getAllQuotes(w http.ResponseWriter, r *http.Request) {
    err := json.NewEncoder(w).Encode(quotes)
    if err != nil {
        fmt.Printf("There was an error. \n")
    }
}

func main() {
    loadQuotes()
    addr, err := determineListenAddress()
    if err != nil {
        log.Fatal(err)
    }
    router := mux.NewRouter().StrictSlash(true)
    router.HandleFunc("/", homeLink)
    router.HandleFunc("/quotes", getAllQuotes).Methods("GET")
    router.HandleFunc("/quotes/random", getRandomQuote).Methods("GET")
    router.HandleFunc("/quotes/{id}", getOneQuote).Methods("GET")
    log.Fatal(http.ListenAndServe(addr, router))
}
