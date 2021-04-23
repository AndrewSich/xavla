package main

import (
  "net/http"
  "xavla/system"
)

func main() {
  // http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
  //   w.Write([]byte("Welcome to Xaavla Web Framework"))
  // })

  // This a routing function
  http.HandleFunc("/", system.HandlerIndex)

  // Static routing system
  http.Handle("/static/",
    http.StripPrefix("/static/",
      http.FileServer(http.Dir("public"))))

  var address = ":0"
  system.Server(address)
}go 