package handlers

import (
	"net/http"
	"fmt"
)

func RootHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Server is running")
}