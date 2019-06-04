
# TestServer

Chat App using WebSocket, and Authorization with OAuth 2.0.

WebSocketを使ったチャットアプリと、OAuth 2.0によるユーザ認証

TODO: Migrate from Oauth 2.0 to OpenID Connect

### Requirements

- Golang
- go get github.com/gorilla/websocket
- go get github.com/stretchr/gomniauth
- go get github.com/stretchr/gomniauth/providers/google
- go get github.com/stretchr/gomniauth/providers/github
- go get github.com/stretchr/objx

### Execution

Type following:

```
cd chat
go build -o chat
./chat
```

then, access to [http://localhost:8080/chat](http://localhost:8080/chat).

Supported authorization provider is:

- Google (https://console.developers.google.com/apis/credentials)
- GitHub (Settings > Developer settings > OAuth Apps)
