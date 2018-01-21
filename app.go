package main

import (
    "net/http"
    "strconv"
    "log"
    "math/rand"
    "time"
    "os"
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

    source := rand.NewSource(time.Now().UnixNano())
    root := rand.New(source)
    dollars := root.Intn(10000)
    cents := root.Intn(99)
    
    price := strconv.Itoa(dollars) + "." + strconv.Itoa(cents)
    stock := mux.Vars(r)["stock"]
    crypto := generateString(10)

    message := price + "," + stock + "," + crypto
    w.Write([]byte(message))
}

func main() {
    router :=  mux.NewRouter()
    port, err := strconv.Atoi(os.Getenv("QUOTE_SERVER_PORT"))
    
    if err != nil{
        log.Fatal(err)
    }

    log.Println("Running quote server on port: " + strconv.Itoa(port) )
    router.HandleFunc("/api/getQuote/{username}/{stock}", getQuote)
    http.Handle("/", router)

    if err := http.ListenAndServe(":" + strconv.Itoa(port), nil); err != nil {
        log.Fatal(err)
        panic(err)
    }
}
