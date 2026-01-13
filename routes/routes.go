package routes

import (
	"net/http"
	"todo-golang/handlers"
)

// SetupRoutes configura todas as rotas da aplicação
func SetupRoutes() *http.ServeMux {
	mux := http.NewServeMux()

	// Rotas de tarefas
	mux.HandleFunc("GET /tarefas", handlers.ListarTarefas)
	mux.HandleFunc("POST /tarefas", handlers.AdicionarTarefa)

	return mux
}
