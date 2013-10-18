package main

import (
  "io"
  "net/http"
  "log"
  "fmt"
  "strings"
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
  err := req.ParseMultipartForm(1000)
  if err != nil {
    log.Fatal("HomeHandler: ", err)
  }

  format := " hit %s%s from %s via %s method\n"
  s := fmt.Sprintf(format, 
                   req.Host,
                   req.URL.String(), 
                   req.RemoteAddr, 
                   req.Method)

  io.WriteString(w, "You " + s)
  fmt.Print("They " + s)

  if len(req.Form) != 0 {
    io.WriteString(w, "\nWith the following data:\n")
    fmt.Println("\nWith the following data:")
    for param, val := range req.Form {
      io.WriteString(w, "-" + param + ": " + strings.Join(val, ", ") + "\n")
      fmt.Println("-" + param + ": " + strings.Join(val, ", "))
    }
    fmt.Print("\n\n=====================\n\n")
  }
}