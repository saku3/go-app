package main

import(
  "log"
  "net/http"
  "path/filepath"
  "text/template"
)

type templateHandler struct {
  filename string
  templ *template.Template
}

func (t *templateHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
  t.templ = template.Must(template.ParseFiles(filepath.Join("templates", t.filename)))
  t.templ.Execute(w, nil)
}

func main() {
  http.Handle("/", &templateHandler{filename: "index.html"})

  if err := http.ListenAndServe(":8080", nil); err != nil {
    log.Fatal("ListenAndServe", err)
  }
}
