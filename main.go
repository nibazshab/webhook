package main

import (
  "fmt"
  "net/http"
)

func main() {
  http.HandleFunc("/0gxFztNrJBeBsCcwuTBdQbxchPDcVEyL", handleWebhook)
  _ = http.ListenAndServe(":10001", nil)
}

func handleWebhook(w http.ResponseWriter, r *http.Request) {
  if r.Method == "POST" {
    fmt.Fprint(w, "OK")
  }
}
