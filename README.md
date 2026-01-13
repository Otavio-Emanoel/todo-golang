# Todo API em Go

API simples para gerenciamento de tarefas (ToDo) construída em Go. No momento, expõe endpoints para listar e criar tarefas, com armazenamento em memória (não persistente).

## Requisitos
- Go 1.22+ (necessário para o roteamento com `http.NewServeMux` usando padrões como `"GET /tarefas"`).
- `curl` ou qualquer cliente HTTP para testar.

## Como executar
```bash
go run main.go
```
O servidor inicia em `http://localhost:8080`.

## Modelo de dados
```json
{
  "id": 1,
  "descricao": "Estudar Go",
  "feita": false
}
```

## Endpoints

### GET /tarefas
- **Descrição:** Lista todas as tarefas.
- **Resposta 200:** Array de tarefas.
- **Exemplo:**
```bash
curl -s http://localhost:8080/tarefas | jq
```

### POST /tarefas
- **Descrição:** Cria uma nova tarefa.
- **Body (JSON):**
```json
{
  "descricao": "Ler documentação do Go",
  "feita": false
}
```
- **Resposta 201:** Tarefa criada com `id` atribuído.
- **Erros 400:** JSON inválido.
- **Exemplo:**
```bash
curl -s -X POST http://localhost:8080/tarefas \
  -H "Content-Type: application/json" \
  -d '{"descricao":"Ler documentação do Go","feita":false}' | jq
```

## Comportamento atual
- Armazenamento **em memória**: ao reiniciar o servidor, os dados voltam ao estado inicial definido em `main.go`.
- Sem paginação, filtros, autenticação ou persistência.

## Roadmap e futuras adições
- **CRUD completo:**
  - `GET /tarefas/{id}`: buscar tarefa por id
  - `PUT /tarefas/{id}`: atualizar `descricao` e `feita`
  - `PATCH /tarefas/{id}`: atualizações parciais
  - `DELETE /tarefas/{id}`: remover tarefa
- **Filtros e paginação:**
  - `GET /tarefas?feita=true|false`
  - Limites (`limit`) e deslocamento (`offset`)
- **Persistência:**
  - Banco de dados (SQLite/PostgreSQL) com camada de repositório
  - Migrações e configuração via variáveis de ambiente
- **Validação e erros:**
  - Validação de `descricao` obrigatória
  - Respostas de erro padronizadas (JSON)
- **Observabilidade e manutenção:**
  - Logs estruturados
  - Métricas (Prometheus) e perfis (pprof)
- **Infra:**
  - Dockerfile e docker-compose
  - CI para testes e lint
- **Docs:**
  - Especificação OpenAPI/Swagger
  - Exemplos mais extensos de uso
- **Segurança:**
  - CORS configurável
  - Rate limiting básico

## Estrutura do projeto
```
todo-golang/
├── main.go                    # Entry point da aplicação
├── go.mod                     # Dependências e módulo
├── models/
│   └── tarefa.go             # Definição da struct Tarefa e banco em memória
├── handlers/
│   └── tarefa_handler.go     # Handlers HTTP para operações com tarefas
└── routes/
    └── routes.go             # Configuração centralizada de rotas
```

### Descrição dos pacotes
- **`main.go`**: Inicializa o servidor HTTP e chama a configuração de rotas.
- **`models/`**: Define as estruturas de dados (struct `Tarefa`) e armazenamento temporário.
- **`handlers/`**: Contém a lógica de negócio dos handlers HTTP (`ListarTarefas`, `AdicionarTarefa`).
- **`routes/`**: Centraliza a configuração de todas as rotas da aplicação.

## Dicas
- Se estiver usando uma versão de Go anterior a 1.22, adapte o roteamento em [routes/routes.go](routes/routes.go) para `mux.HandleFunc("/tarefas", handler)` e faça o switch por método dentro do handler.
- A estrutura modular facilita a adição de novos recursos: basta criar novos handlers em `handlers/` e registrá-los em `routes/routes.go`.
- Para adicionar persistência com banco de dados, crie um pacote `repository/` ou `database/` e adapte os handlers para usar o repositório ao invés de `models.BancoDeDados`.

## Licença
Projeto educacional/demonstrativo. Use e adapte livremente.
