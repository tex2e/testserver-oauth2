
package main

import (
    "log"
    "net/http"
    "text/template"
    "path/filepath"
    "sync"
    "flag"

    "github.com/stretchr/gomniauth"
    "github.com/stretchr/gomniauth/providers/google"
    "github.com/stretchr/gomniauth/providers/github"
    "github.com/stretchr/objx"
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

    data := map[string]interface{}{}
    data["Host"] = r.Host

    authCookie, err := r.Cookie("auth")
    if err == nil {
        data["UserData"] = objx.MustFromBase64(authCookie.Value)
    }
    t.templ.Execute(w, data)
}

func main() {
    var addr = flag.String("addr", ":8080", "Application Address")
    flag.Parse()

    gomniauth.SetSecurityKey(securityKey)
    gomniauth.WithProviders(
        google.New(googleClientID, googleSecret,
            "http://localhost:8080/auth/callback/google"),
        github.New(githubClientID, githubSecret,
            "http://localhost:8080/auth/callback/github"),
    )

    r := newRoom()
    http.Handle("/chat", MustAuth(&templateHandler{filename: "chat.html"}))
    http.Handle("/login", &templateHandler{filename: "login.html"})
    http.HandleFunc("/auth/", loginHandler)
    http.Handle("/room", r)

    go r.run()

    log.Println("Start WebServer with port", *addr)

    err := http.ListenAndServe(*addr, nil)
    if err != nil {
        log.Fatal("ListenAndServe:", err)
    }
}
