package main

import (
	"database/sql"
	"encoding/json"
	"fmt"
	"log"
	"math/rand"
	"net/http"
	"strings"
	"time"

	_ "github.com/glebarez/go-sqlite"
)

const tamanhoCodigo = 6
const letrasNumeros = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789"

var db *sql.DB

// Estrutura para os dados da requisição
type Requisicao struct {
	Link string `json:"link"`
}

// Estrutura para a resposta
type Resposta struct {
	LinkCurto string `json:"link_curto,omitempty"`
	Erro      string `json:"erro,omitempty"`
}

func init() {
	rand.Seed(time.Now().UnixNano())
}

func main() {
	// Inicializa o banco de dados
	var err error
	db, err = sql.Open("sqlite", "./links.db")
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()

	// Cria a tabela se ela não existir
	sqlStmt := `
	CREATE TABLE IF NOT EXISTS links (
		codigo TEXT PRIMARY KEY,
		link_longo TEXT NOT NULL
	);`
	_, err = db.Exec(sqlStmt)
	if err != nil {
		log.Fatalf("Erro ao criar a tabela: %v", err)
	}

	// Define os manipuladores de rotas
	http.HandleFunc("/", rotaPrincipal)
	http.HandleFunc("/encurtar", rotaEncurtar)

	fmt.Println("Servidor iniciado em http://localhost:8080")
	log.Fatal(http.ListenAndServe(":8080", nil))
}

// Manipulador da rota principal
func rotaPrincipal(w http.ResponseWriter, r *http.Request) {
	if r.URL.Path == "/" {
		// Serve o arquivo index.html na raiz
		http.ServeFile(w, r, "index.html")
		return
	}

	// Tenta obter o código curto da URL
	codigo := strings.TrimPrefix(r.URL.Path, "/")
	if len(codigo) == 0 {
		http.Error(w, "Código não encontrado", http.StatusNotFound)
		return
	}

	// Busca o link longo no banco de dados
	var linkLongo string
	err := db.QueryRow("SELECT link_longo FROM links WHERE codigo = ?", codigo).Scan(&linkLongo)
	if err != nil {
		if err == sql.ErrNoRows {
			http.Error(w, "Link encurtado não encontrado", http.StatusNotFound)
			return
		}
		http.Error(w, "Erro interno do servidor", http.StatusInternalServerError)
		log.Printf("Erro ao buscar no banco de dados: %v", err)
		return
	}

	// Redireciona o usuário
	http.Redirect(w, r, linkLongo, http.StatusFound)
}

// Manipulador da rota de encurtar
func rotaEncurtar(w http.ResponseWriter, r *http.Request) {
	if r.Method != "POST" {
		http.Error(w, "Método não permitido", http.StatusMethodNotAllowed)
		return
	}

	var req Requisicao
	decoder := json.NewDecoder(r.Body)
	err := decoder.Decode(&req)
	if err != nil {
		http.Error(w, "Requisição JSON inválida", http.StatusBadRequest)
		return
	}

	if req.Link == "" {
		http.Error(w, "Link não fornecido", http.StatusBadRequest)
		return
	}

	// Tenta gerar um código único
	var codigo string
	var tentativas int
	for {
		codigo = gerarCodigo()
		var existe int
		err := db.QueryRow("SELECT COUNT(*) FROM links WHERE codigo = ?", codigo).Scan(&existe)
		if err != nil {
			http.Error(w, "Erro ao verificar código", http.StatusInternalServerError)
			log.Printf("Erro ao verificar código: %v", err)
			return
		}
		if existe == 0 {
			break
		}
		tentativas++
		if tentativas > 10 { // Evita loop infinito em caso de falha na geração
			http.Error(w, "Não foi possível gerar um código único", http.StatusInternalServerError)
			return
		}
	}

	// Insere o par no banco de dados
	_, err = db.Exec("INSERT INTO links (codigo, link_longo) VALUES (?, ?)", codigo, req.Link)
	if err != nil {
		http.Error(w, "Erro ao salvar no banco de dados", http.StatusInternalServerError)
		log.Printf("Erro ao inserir link: %v", err)
		return
	}

	// Constrói o link curto completo
	linkCurto := fmt.Sprintf("http://%s/%s", r.Host, codigo)

	// Retorna a resposta JSON
	resp := Resposta{LinkCurto: linkCurto}
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(resp)
}

// Função para gerar um código aleatório
func gerarCodigo() string {
	b := make([]byte, tamanhoCodigo)
	for i := range b {
		b[i] = letrasNumeros[rand.Intn(len(letrasNumeros))]
	}
	return string(b)
}
