package main

import (
  "io"
  "net/http"
  "log"
  "fmt"
  "github.com/gorilla/mux"
)

func main() {
  r := mux.NewRouter();
  r.HandleFunc("/", HomeHandler)

  http.Handle("/", r);

  err := http.ListenAndServe(":12345", nil)
  if err != nil {
    log.Fatal("ListenAndServe: ", err)
  }
}

func HomeHandler(w http.ResponseWriter, req *http.Request) {
  s := "You hit %s from %s via %s method\n"
  io.WriteString(w, fmt.Sprintf(s, 
                                req.URL.String(), 
                                req.Host, 
                                req.Method))
}