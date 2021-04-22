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

  // This a routing static files
  http.Handle("/static/",
    http.StripPrefix("/static/",
      http.FileServer(http.Dir("public"))))

  var address = ":0"
  fmt.Printf("Server on start running")
  err := http.ListenAndServe(address, nil)
  if err != nil {
    fmt.Println(err.Error())
  }
}