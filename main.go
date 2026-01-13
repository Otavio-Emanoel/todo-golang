package main

import (
	"fmt"
	"net/http"
	"todo-golang/routes"
)

func main() {
	// Configura todas as rotas
	mux := routes.SetupRoutes()

	fmt.Println("Servidor rodando na porta 8080")
	err := http.ListenAndServe(":8080", mux)
	if err != nil {
		fmt.Println("Erro ao iniciar o servidor:", err)
	}
}
