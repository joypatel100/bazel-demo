package main

import (
    "log"
    "net/http"
    "github.com/gorilla/mux"
    "fmt"
    "io/ioutil"
)

func HomeHandler(w http.ResponseWriter, r *http.Request) {
  w.WriteHeader(http.StatusOK)
  fmt.Fprintf(w, "Go Home \n")
  resp, err := http.Get("http://localhost:5000/")
  if err != nil {
	   fmt.Printf("error")
  }
  defer resp.Body.Close()
  bodyBytes, err := ioutil.ReadAll(resp.Body)
  if err != nil {
	   fmt.Printf("error")
  }
  fmt.Printf(string(bodyBytes))
}

// our main function
func main() {
    router := mux.NewRouter()
    router.HandleFunc("/", HomeHandler)
    log.Fatal(http.ListenAndServe(":8000", router))
}
