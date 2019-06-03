
package main

import (
    "log"
    "net/http"
    "text/template"
    "path/filepath"
    "sync"
)

type templateHandler struct {
    once     sync.Once
    filename string
    templ    *template.Template
}

func (t *templateHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
    t.once.Do(func () {
        t.templ = template.Must(template.ParseFiles(
            filepath.Join("templates", t.filename)))
    })
    t.templ.Execute(w, nil)
}

func main()  {
    http.Handle("/", &templateHandler{filename: "chat.html"})

    err := http.ListenAndServe(":8080", nil)
    if err != nil {
        log.Fatal("ListenAndServe:", err)
    }
}
