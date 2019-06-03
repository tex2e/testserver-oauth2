
package main

import (
    "log"
    "net/http"
    "text/template"
    "path/filepath"
    "sync"
    "flag"
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
    t.templ.Execute(w, r)
}

func main() {
    var addr = flag.String("addr", ":8080", "Application Address")
    flag.Parse()

    r := newRoom()
    http.Handle("/", &templateHandler{filename: "chat.html"})
    http.Handle("/room", r)

    go r.run()

    log.Println("Start WebServer with port", *addr)

    err := http.ListenAndServe(*addr, nil)
    if err != nil {
        log.Fatal("ListenAndServe:", err)
    }
}
