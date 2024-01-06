# go-auth-middleware

`go-auth-middleware` is a small Go library that provides middleware functionalities for authentication and authorization in HTTP handlers.

## Features

- **Authentication Middleware**: Allows authenticating incoming requests using custom authentication functions.
- **Authorization Middleware**: Enables authorization by validating request context against provided authorization functions.
- **Basic Authentication Middleware**: Offers basic authentication support using HTTP Basic Auth headers.

## Installation

You can install the package using `go get`:

```bash
go get github.com/Eyal-Shalev/go-auth-middleware
```

## Usage

### Authentication Middleware

```go
// Example authentication function
func authenticate(r *http.Request) (value T, ok bool, err error) {
// Your authentication logic here
}

// Create authentication middleware
authMiddleware := authMiddleware.NewAuthenticateMiddleware(authenticate)

// Use `authMiddleware` in your HTTP handlers
```

### Authorization Middleware

```go
// Example authorization function
func authorize(r *http.Request, value T) bool {
// Your authorization logic here
}

// Create authorization middleware
authzMiddleware := authMiddleware.NewAuthorizationMiddleware(authorize)

// Use `authzMiddleware` in your HTTP handlers after authentication
```

### Basic Authentication Middleware

```go
// Example basic authentication function
func basicAuthenticator(ctx context.Context, username, password string) (T, bool, error) {
// Your basic authentication logic here
}

// Create basic authentication middleware
basicAuthMiddleware := authMiddleware.NewBasicAuthenticateMiddleware(basicAuthenticator)

// Use `basicAuthMiddleware` in your HTTP handlers
```

## Example

Here's a simple example demonstrating how to use the library:

```go
package main

import (
	authMiddleware "github.com/Eyal-Shalev/go-auth-middleware"
	"net/http"
	"strings"
)

type httpStringHandler string

func (h httpStringHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	_, _ = w.Write([]byte(h))
}


const helloWorld httpStringHandler = "Hello, World!"
const helloAdmin httpStringHandler = "Hello, Admin!"

func basicAuthenticate(ctx context.Context, userName, password string) (string, bool, error) {
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
```

## Contributing

Feel free to contribute by submitting issues or pull requests!

## License

This library is licensed under the MIT License. See the [LICENSE](LICENSE) file for details.
