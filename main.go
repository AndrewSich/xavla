package main

import (
  "fmt"
  "net/http"
  "xavla/route"
)

func main() {
  // http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
  //   w.Write([]byte("Welcome to Xaavla Web Framework"))
  // })

  // This a routing function
  http.HandleFunc("/", route.HandlerIndex)

  var address = ":0"
  fmt.Printf("Server on start running")
  err := http.ListenAndServe(address, nil)
  if err != nil {
    fmt.Println(err.Error())
  }
}