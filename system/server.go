package system

import (
  "fmt"
  "net/http"
)

func Server(port string) {
  fmt.Printf("Server start running on\n")
  fmt.Printf("[+] success! \n")
  err := http.ListenAndServe(port, nil)
  if err != nil {
    fmt.Println(err.Error())
    fmt.Printf("[-] warning! \n")
  }
}