# Encurtador de Links Simples

![Go](https://img.shields.io/badge/go-%2300ADD8.svg?style=for-the-badge&logo=go&logoColor=white)
![SQLite](https://img.shields.io/badge/sqlite-%2307405e.svg?style=for-the-badge&logo=sqlite&logoColor=white)
![HTML5](https://img.shields.io/badge/html5-%23E34F26.svg?style=for-the-badge&logo=html5&logoColor=white)
![TailwindCSS](https://img.shields.io/badge/tailwindcss-%2338B2AC.svg?style=for-the-badge&logo=tailwind-css&logoColor=white)

Um projeto simples e funcional de um encurtador de links. Ele foi desenvolvido como um exercício prático para aprender e demonstrar o uso de **Go (Golang)** no back-end, um banco de dados **SQLite** embutido e uma interface de usuário leve com **HTML** e **Tailwind CSS**.

## Funcionalidades

- **Encurtar links:** Converte URLs longas em links curtos e únicos.
- **Redirecionamento:** Redireciona o usuário do link encurtado para a URL original.
- **Armazenamento embutido:** Usa um banco de dados SQLite para salvar os links, sem a necessidade de um servidor de banco de dados externo.
- **Interface simples:** Uma única página para todas as interações.

## Tecnologias Utilizadas

- **Back-end:**
  - **Go (Golang):** Para o servidor web e a lógica do sistema.
  - **SQLite:** Banco de dados relacional embutido para persistir os dados.
- **Front-end:**
  - **HTML:** Estrutura da página.
  - **Tailwind CSS (via CDN):** Para a estilização rápida e responsiva.

## Estrutura do Projeto

O projeto é minimalista e consiste em apenas dois arquivos principais:
/
├── main.go            # Lógica do servidor, banco de dados e rotas.
└── index.html         # Front-end da aplicação.

## Como Executar o Projeto

Siga os passos abaixo para rodar a aplicação em seu ambiente local.

### Pré-requisitos

- **Go** (versão 1.16 ou superior)

### Passos

1.  **Clone o repositório:**
    ```bash
    git clone https://github.com/marquesjosemar/encurtador-ads.git 
    cd seu-repositorio
    ```

2.  **Inicialize o módulo Go e baixe as dependências:**
    A versão mais recente do Go requer a inicialização de um módulo para gerenciar as dependências. Execute os comandos abaixo na pasta do projeto:
    ```bash
    go mod init encurtador-links
    go get [github.com/mattn/go-sqlite3](https://github.com/mattn/go-sqlite3)
    ```

3.  **Execute o servidor:**
    O projeto usa a biblioteca `go-sqlite3`, que requer o **CGO** habilitado. Execute o comando a seguir:
    
    **Windows (PowerShell):**
    ```bash
    $env:CGO_ENABLED=1; go run .
    ```
    
    **Linux/macOS:**
    ```bash
    CGO_ENABLED=1 go run .
    ```

4.  **Acesse a aplicação:**
    Abra seu navegador e navegue até `http://localhost:8080`.

## Contribuição

Contribuições são bem-vindas! Se você tiver alguma ideia para melhorar o projeto, sinta-se à vontade para abrir uma *issue* ou enviar um *pull request*.

---

## Licença

Este projeto está licenciado sob a licença MIT. Veja o arquivo [LICENSE](LICENSE) para mais detalhes.
