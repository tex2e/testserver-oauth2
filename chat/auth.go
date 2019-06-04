
package main

import (
    "fmt"
    "log"
    "strings"
    "net/http"

    "github.com/stretchr/gomniauth"
    // "github.com/stretchr/objx"
)

type authHandler struct {
    next http.Handler
}

func (h *authHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
    _, err := r.Cookie("auth")
    if err == http.ErrNoCookie {
        // 未認証
        w.Header().Set("Location", "/login")
        w.WriteHeader(http.StatusTemporaryRedirect)
    } else if err != nil {
        // エラー
        panic(err.Error())
    } else {
        // 成功
        h.next.ServeHTTP(w, r)
    }
}

func MustAuth(handler http.Handler) http.Handler {
    return &authHandler{ next: handler }
}

// /auth/{action}/{provider}
func loginHandler(w http.ResponseWriter, r *http.Request) {
    segs := strings.Split(r.URL.Path, "/")
    action := segs[2]
    provider := segs[3]
    switch action {
    case "login":
        provider, err := gomniauth.Provider(provider)
        if err != nil {
            log.Fatalln("認証プロバーダーの取得に失敗しました:", provider, "-", err)
        }
        loginUrl, err := provider.GetBeginAuthURL(nil, nil)
        if err != nil {
            log.Fatalln("GetBeginAuthURLの呼び出し中にエラーが発生しました:",
                provider, "-", err)
        }
        w.Header().Set("Location", loginUrl)
        w.WriteHeader(http.StatusTemporaryRedirect)
    default:
        w.WriteHeader(http.StatusNotFound)
        fmt.Fprintf(w, "Action %s is not supported!", action)
    }
}
