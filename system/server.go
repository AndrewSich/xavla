package system

import (
  "fmt"
  "net/http"
)

// This a Core Server Xavla
func StartAndRunServer(port string) {
  // add routing base
  routingBase()

  // running server
  fmt.Printf("Server start running on\n")
  fmt.Printf("[+] success! \n")
  err := http.ListenAndServe(port, nil)
  if err != nil {
    fmt.Println(err.Error())
    fmt.Printf("[-] warning! \n")
  }
}