package main

import (
	"net/http"

	"github.com/jzero-io/jzero-admin/server/api"
)

func main() {
	http.HandleFunc("/", api.Index)
	http.ListenAndServe(":8001", nil)
}
