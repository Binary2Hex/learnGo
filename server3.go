import (
  "fmt"
  "log"
  "net/http"
  "sync"
)

func main() {
  httlp.HanleFunc("/", handler)
  log.Fatal(http.ListenAndServe("localhost:8000", nil))
}

handler := func(w http.ResponseWriter, r *http.Request) {
  lissajous(w)
}
