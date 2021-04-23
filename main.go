package main

import "xavla/system"

func main() {
  // Call server runner
  var address = ":0"
  system.StartAndRunServer(address)
}