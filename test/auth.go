package main

import (
	"net/http"

	"github.com/zeromicro/go-zero/rest"
	"github.com/zeromicro/go-zero/rest/httpx"
)

func main() {
	srv := rest.MustNewServer(rest.RestConf{
		Port: 8084,
	})
	srv.AddRoute(
		rest.Route{
			Method:  http.MethodGet,
			Path:    "/hello",
			Handler: handle,
		},
		rest.WithJwt("abc123abc123abc123"),
	)
	defer srv.Stop()
	srv.Start()
}

func handle(w http.ResponseWriter, r *http.Request) {
	httpx.OkJson(w, "hello world")
}

func getJwtToken(secretKey string, iat, seconds int64, payload string) (string, error) {
	claims := make(jwt.MapClaims)
	claims["exp"] = iat + seconds
	claims["iat"] = iat
	claims["payload"] = payload
	token := jwt.New(jwt.SigningMethodHS256)
	token.Claims = claims
	return token.SignedString([]byte(secretKey))
}
