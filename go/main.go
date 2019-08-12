package main

import (
    "log"
    "net/http"
    "time"
    "github.com/khatribharat/quickstart-web-stacks/go/views"
)

func main() {
    http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
        data := struct {
            CurrentTime string
        }{
            time.Now().Format(time.RFC822),
        }
        w.Write(views.GenerateHTMLForHome(data))
    })
    log.Fatal(http.ListenAndServe(":8080", nil))
}
