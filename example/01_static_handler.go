package main

import "net/http"

type httpStringHandler string

func (h httpStringHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	_, _ = w.Write([]byte(h))
}
