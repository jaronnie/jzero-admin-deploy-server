package main

import (
	"net/http"

	"github.com/jzero-io/jzero-admin/server/api"
)

func main() {
	err := http.ListenAndServe(":8001", http.HandlerFunc(api.Index))
	if err != nil {
		panic(err)
	}
}
