package models

// Tarefa representa uma tarefa no sistema
type Tarefa struct {
	ID        int    `json:"id"`
	Descricao string `json:"descricao"`
	Feita     bool   `json:"feita"`
}

// BancoDeDados é um slice temporário para armazenar as tarefas em memória
var BancoDeDados = []Tarefa{
	{ID: 1, Descricao: "Estudar Go", Feita: false},
	{ID: 2, Descricao: "Instalar Go no Fedora", Feita: true},
}
