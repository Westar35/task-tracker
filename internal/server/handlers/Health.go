package handlers

import (
	"net/http"
	"fmt"
)

func HealthHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "ok")
}