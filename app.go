package main

import (
	"log"
	"math/rand"
	"net/http"
	"strconv"
	"time"

	"github.com/gorilla/mux"
)



func generateString(n int) string {

	var letters = []rune("abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ")
	b := make([]rune, n)
	for i := range b {
		b[i] = letters[rand.Intn(len(letters))]
	}
	return string(b)
}

func getQuote(w http.ResponseWriter, r *http.Request) {
    stock := mux.Vars(r)["stock"]

    source := rand.NewSource(time.Now().UnixNano())
    root := rand.New(source)

    dollars := root.Intn(10000)
    cents := root.Intn(99)
    
    price := strconv.Itoa(dollars) + "." + strconv.Itoa(cents)

    if stock == "TEST" {
        price = "200.00"
    }

    crypto := generateString(10)

    message := price + "," + stock + "," + crypto
    w.Write([]byte(message))
}

func main() {
	router := mux.NewRouter()
	port := 8000

	log.Println("Running quote server on port: " + strconv.Itoa(port))
	router.HandleFunc("/api/getQuote/{username}/{stock}", getQuote)
	http.Handle("/", router)

	if err := http.ListenAndServe(":"+strconv.Itoa(port), nil); err != nil {
		log.Fatal(err)
		panic(err)
	}
}
