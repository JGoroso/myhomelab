package main
import (
  "fmt"
  "log"
  "net/http"
)
func main() {
  http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
    fmt.Fprintln(w, "Hello from my Homelab via Helm + CI/CD 🚀")
  })
  log.Fatal(http.ListenAndServe(":8080", nil))
}
