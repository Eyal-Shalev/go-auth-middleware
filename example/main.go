package main

import (
	"context"
	"net/http"
	"strings"

	authMiddleware "github.com/Eyal-Shalev/go-auth-middleware"
)

const helloWorld httpStringHandler = "Hello, World!"
const helloAdmin httpStringHandler = "Hello, Admin!"

func basicAuthenticate(_ context.Context, userName, password string) (string, bool, error) {
	if password == strings.ToUpper(userName) {
		return userName, true, nil
	}
	return "", false, nil
}

func isAdmin(_ *http.Request, userName string) bool {
	return userName == "admin"
}

func main() {
	authenticate := authMiddleware.BasicAuthFunc[string](basicAuthenticate)
	authorizeAdmin := authMiddleware.AuthorizeFunc[string](isAdmin)
	http.Handle("/", helloWorld)
	http.Handle("/admin", authenticate.Wrap(authorizeAdmin.Wrap(helloAdmin)))
	err := http.ListenAndServe("localhost:9090", nil)
	if err != nil {
		panic(err)
	}
}
