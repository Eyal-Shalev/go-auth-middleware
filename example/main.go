package main

import (
	authMiddleware "github.com/Eyal-Shalev/go-auth-middleware"
	"net/http"
	"strings"
)

const helloWorld httpStringHandler = "Hello, World!"
const helloAdmin httpStringHandler = "Hello, Admin!"

func basicAuthenticate(userName, password string) (string, bool, error) {
	if password == strings.ToUpper(userName) {
		return userName, true, nil
	}
	return "", false, nil
}

func main() {
	authenticate := authMiddleware.NewBasicAuthenticateMiddleware(basicAuthenticate)
	authorizeAdmin := authMiddleware.NewAuthorizationMiddleware(func(r *http.Request, value string) bool {
		return value == "admin"
	})
	http.Handle("/", helloWorld)
	http.Handle("/admin", authenticate(authorizeAdmin(helloAdmin)))
	err := http.ListenAndServe("localhost:9090", nil)
	if err != nil {
		panic(err)
	}
}