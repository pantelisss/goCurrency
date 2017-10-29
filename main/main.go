package main

import (
  "fmt"
  "coin_currency/managers"
)

func main() {
  	fmt.Println("Retrieving from ", manager.Shared().Url)

  	manager.Shared().GetJson()
}
