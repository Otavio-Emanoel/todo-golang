package handlers

import (
	"encoding/json"
	"net/http"
	"todo-golang/models"
)

// ListarTarefas retorna todas as tarefas em formato JSON
func ListarTarefas(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(models.BancoDeDados)
}

// AdicionarTarefa adiciona uma nova tarefa ao banco de dados
func AdicionarTarefa(w http.ResponseWriter, r *http.Request) {
	var novaTarefa models.Tarefa

	err := json.NewDecoder(r.Body).Decode(&novaTarefa)
	if err != nil {
		http.Error(w, "Erro ao decodificar a tarefa", http.StatusBadRequest)
		return
	}

	novaTarefa.ID = len(models.BancoDeDados) + 1
	models.BancoDeDados = append(models.BancoDeDados, novaTarefa)

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(novaTarefa)
}
