package main

import (
  "log"
  "net/http"
)

func main() {
  //your http.Handle calls here
  log.Fatal(http.ListerAndServe("localhost:4000", nil))
}

