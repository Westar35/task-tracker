package handlers

import (
	"net/http"
	"fmt"
)

func TasksHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method == http.MethodGet {
		fmt.Fprintf(w, "Заглушка: GET списка задач")
	}
	if r.Method == http.MethodPost {
		fmt.Fprintf(w, "Заглушка: POST создание задачи")
	}
}