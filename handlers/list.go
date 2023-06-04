package handlers

import (
	"encoding/json"
	"go_learn/models"
	"log"
	"net/http"
)

func List(w http.ResponseWriter, r *http.Request) {
	todo, err := models.GetAll()
	if err != nil {
		log.Println("Erro ao buscar todos: %v", err)
	}

	w.Header().Add("Content-Type", "application/json")
	json.NewEncoder(w).Encode(todo)
}