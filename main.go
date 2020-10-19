package main

import (
    "os"
    "fmt"
    "log"
    "math/rand"
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
        Author: "Edith Head",
        Content: "You can have anything you want in life if you dress for it.",
    },
    {
        ID: "6",
        Author: "Carrie Bradshaw",
        Content: "I like my money right where I can see it…hanging in my closet.",
    },
    {
        ID: "7",
        Author: "Alexander McQueen",
        Content: "I think there is beauty in everything. What 'normal' people perceive as ugly, I can usually see something of beauty in it.",
    },
    {
        ID: "8",
        Author: "Diane von Furstenberg",
        Content: "Style is something each of us already has, all we need to do is find it.",
    },
    {
        ID: "9",
        Author: "Bill Cunningham",
        Content: "Fashion is the armor to survive the reality of everyday life.",
    },
    {
        ID: "10",
        Author: "Ralph Lauren",
        Content: "I don't design clothes. I design dreams.",
    },
    {
        ID: "11",
        Author: "Yves Saint Laurent",
        Content: "Fashions fade, style is eternal.",
    },
    {
        ID: "12",
        Author: "Alexander Wang",
        Content: "Anyone can get dressed up and glamorous, but it is how people dress in their days off that are the most intriguing.",
    },
    {
        ID: "13",
        Author: "Sonia Rykiel",
        Content: "How can you live the high life if you do not wear the high heels?",
    },
    {
        ID: "14",
        Author: "Elsa Schiaparelli",
        Content: "In difficult times, fashion is always outrageous.",
    },
    {
        ID: "15",
        Author: "Marc Jacobs",
        Content: "Clothes mean nothing until someone lives in them",
    },
    {
        ID: "16",
        Author: "Diana Vreeland",
        Content: "You gotta have style. It helps you get down the stairs. It helps you get up in the morning. It’s a way of life. Without it, you’re nobody. I’m not talking about lots of clothes.",
    },
    {
        ID: "17",
        Author: "Vivienne Westwood",
        Content: "Fashion is very important. It is life-enhancing and, like everything that gives pleasure, it is worth doing well.",
    },
    {
        ID: "18",
        Author: "Christian Dior",
        Content: "You can never take too much care over the choice of your shoes. Too many women think that they are unimportant, but the real proof of an elegant woman is what is on her feet.",
    },
    {
        ID: "19",
        Author: "Lauren Hutton",
        Content: "Fashion is what you're offered four times a year by designers. And style is what you choose.",
    },
    {
        ID: "20",
        Author: "Hubert de Givenchy",
        Content: "The dress must follow the body of a woman, not the body following the shape of the dress.",
    },
    {
        ID: "21",
        Author: "Marc Jacobs",
        Content: "I always find beauty in things that are odd and imperfect, they are much more interesting.",
    },
    {
        ID: "22",
        Author: "Iris Apfel",
        Content: "Fashion you can buy, but style you possess. The key to style is learning who you are, which takes years. There's no how-to road map to style. It's about self expression and, above all, attitude.",
    },
    {
        ID: "23",
        Author: "Rachel Zoe",
        Content: "Style is a way to say who you are without having to speak.",
    },
    {
        ID: "24",
        Author: "Karl Lagerfeld",
        Content: "Trendy is the last stage before tacky.",
    },
    {
        ID: "25",
        Author: "Harry Winston",
        Content: "People will stare. Make it worth their while.",
    },
    {
        ID: "26",
        Author: "Cristóbal Balenciaga",
        Content: "Elegance is elimination.",
    },
    {
        ID: "27",
        Author: "Christian Louboutin",
        Content: "Shoes transform your body language and attitude. They lift you physically and emotionally.",
    },
    {
        ID: "28",
        Author: "Alber Elbaz",
        Content: "Style is the only thing you can’t buy. It’s not in a shopping bag, a label, or a price tag. It’s something reflected from our soul to the outside world—an emotion.",
    },
    {
        ID: "29",
        Author: "Coco Chanel",
        Content: "In order to be irreplaceable one must always be different",
    },
    {
        ID: "30",
        Author: "Yves Saint Laurent",
        Content: "We must never confuse elegance with snobbery.",
    },
    {
        ID: "31",
        Author: "Kenzo Takada",
        Content: "Fashion is like eating, you shouldn't stick to the same menu.",
    },
    {
        ID: "32",
        Author: "Kate Spade",
        Content: "Playing dress-up begins at age five and never truly ends.",
    },
    {
        ID: "33",
        Author: "Giorgio Armani",
        Content: "Elegance is not standing out, but being remembered.",
    },
    {
        ID: "34",
        Author: "Giambattista Valli",
        Content: "The hardest thing in fashion is not to be known for a logo, but to be known for a silhouette.",
    },
    {
        ID: "35",
        Author: "Vera Wang",
        Content: "I want people to see the dress, but focus on the woman.",
    },
    {
        ID: "36",
        Author: "Linda Evangelista",
        Content: "We have this saying, Christy and I. We don't wake up for less than $10,000 a day.",
    },
    {
        ID: "37",
        Author: "Azzedine Alaïa",
        Content: "I make clothes, women make fashion.",
    },
    {
        ID: "38",
        Author: "Iris Apfel",
        Content: "What's my style is not your style, and I don’t see how you can define it. It’s something that expresses who you are in your own way.",
    },
    {
        ID: "39",
        Author: "Azzedine Alaïa",
        Content: "I make clothes, women make fashion.",
    },
    {
        ID: "40",
        Author: "Yves Saint Laurent",
        Content: "Over the years I have learned that what is important in a dress is the woman who's wearing it.",
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
