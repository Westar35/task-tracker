package handlers

import (
	"net/http"
	"fmt"
)

func GetTasksHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodGet {
		http.Error(w, "Method Not Allowed", http.StatusMethodNotAllowed)
	}
	fmt.Fprintf(w, "Заглушка: GET списка задач")
}