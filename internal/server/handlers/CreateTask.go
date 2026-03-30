package handlers

import (
	"net/http"
	"fmt"
)

func CreateTaskHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		http.Error(w, "Method Not Allowed", http.StatusMethodNotAllowed)
	}
	fmt.Fprintf(w, "Заглушка: POST создание задачи")
}