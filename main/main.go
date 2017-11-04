package main

import (
  "fmt"
  "coin_currency/managers"
  "log"
  "net/http"
)

func main() {
  	fmt.Println("Retrieving from ", manager.Shared().Url)

  	manager.Shared().GetJson()

  	http.HandleFunc("/", sayhelloName) // set router
    err := http.ListenAndServe(":9090", nil) // set listen port
    if err != nil {
        log.Fatal("ListenAndServe: ", err)
    }

}


func sayhelloName(w http.ResponseWriter, r *http.Request) {
    r.ParseForm()  // parse arguments, you have to call this by yourself

    // Serve JSON
    w.Header().Set("Content-Type", "application/json")
    w.Write(manager.Shared().JsonData())
}