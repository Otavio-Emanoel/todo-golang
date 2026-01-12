package main

import (
	"fmt"
	"net/http"
	"encoding/json"
)

type Tarefa struct {
	ID         int    `json:"id"`
	Descricao  string `json:"descricao"`
	Feita 	   bool   `json:"feita"`
}

var bancoDeDados = []Tarefa{
	{ID: 1, Descricao: "Estudar Go", Feita: false},
	{ID: 2, Descricao: "Instalar Go no Fedora", Feita: true},
}

// listar tarefas retorna todas as tarefas em formato JSON
func listarTarefas(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(bancoDeDados)
}

// adicionar tarefa adiciona uma nova tarefa ao banco de dados
func adicionarTarefa(w http.ResponseWriter, r *http.Request) {
	var novaTarefa Tarefa

	err := json.NewDecoder(r.Body).Decode(&novaTarefa)
	if err != nil {
		http.Error(w, "Erro ao decodificar a tarefa", http.StatusBadRequest)
		return
	}
	
	novaTarefa.ID = len(bancoDeDados) + 1
	bancoDeDados = append(bancoDeDados, novaTarefa)

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(novaTarefa)
}

func main() {
	mux := http.NewServeMux()

	mux.HandleFunc("GET /tarefas", listarTarefas)
	mux.HandleFunc("POST /tarefas", adicionarTarefa)

	fmt.Println("Servidor rodando na porta 8080")
	err := http.ListenAndServe(":8080", mux)
	if err != nil {
		fmt.Println("Erro ao iniciar o servidor:", err)
	}
}