
# **DevBook**

Uma aplicação web para gerenciar suas publicações e seu usuário, com a funcionalidade de criar, visualizar, atualizar e deletar publicações e usuários.

## **Índice**
1. [Sobre o Projeto](#sobre-o-projeto)
2. [Tecnologias Utilizadas](#tecnologias-utilizadas)
3. [Outras Ferramentas](#outras-ferramentas)
4. [Funcionalidades](#funcionalidades)
5. [Instalação](#instalação)
6. [Estrutura de Diretórios](#estrutura-de-diretórios)
7. [API Endpoints](#api-endpoints)

---

## **Sobre o Projeto**
Este é um projeto de gerenciamento de publicações e usuário, desenvolvido para fins de aprendizado e prática com tecnologias web. Ele inclui:
- Interface responsiva para desktop e dispositivos móveis.
- CRUD completo de usuários e publicações.
- Integração com uma API para gerenciamento de dados.
- Autenticação via JWT.
- Documentação da API utilizando Swagger.

---

## **Tecnologias Utilizadas**
### **Frontend**
- **Linguagem:** [Go](https://go.dev/learn/)
- **Estilização:** [Bootstrap](https://getbootstrap.com/)

### **Backend**
- **Linguagem:** [Go](https://go.dev/learn/)
- **Banco de Dados:** [MySQL](https://dev.mysql.com/doc/)

### **Outras Ferramentas**
- **Documentação da API:** Swagger.
- **Autenticação:** JWT.

---

## **Funcionalidades**
- **Usuários:**
  - Cadastro e login para o usuário.
  - Buscar e visualizar perfis.
  - Atualizar e excluir seu usuário.
  - Seguir e deixar de seguir outros usuários.

  - **Publicações:**
  - Criar, visualizar, atualizar e excluir suas publicações.
  - Curtir e descurtir outras publicações.

  - **Senha:**
  - Atualização de senha do usuário logado,

  - **Autenticação:**
  - Login e geração de tokens JWT.

  - **Segurança:**
  - Requisições autenticadas com Bearer Token.

- **Feedbacks Visuais:**
  - Notificações e alertas utilizando SweetAlert.

- **Interface Responsiva:**
  - Adaptada para dispositivos móveis e desktops.

- **Documentação da API:**
  - Swagger disponível em `http://localhost:5000/swagger/index.html`.

---

## **Instalação**
### Pré-requisitos
- **Go** (versão mínima: `1.20`): [Download Go](https://go.dev/)
- **MySQL** (versão mínima: `8.0`): [Download Git](https://dev.mysql.com/downloads/)
- **Git** (opcional, mas recomendado): [Download Git](https://git-scm.com/)

### Passos para rodar o projeto:

1. Clone este repositório:
   ```bash
   git clone https://github.com/Heloisatanquella/DevBook.git
   cd DevBook

1. Configurar o banco de dados:
 - **Crie um banco de dados MySQL e configure as credenciais no backend.**

2. Iniciar o Backend:
 - **Navegue para /Backend/api**
 - **Instale as dependências:** go mod tidy
 - **Inicie o servidor:** go run main.go
 - **O backend estará disponível em http://localhost:5000**

 3. Iniciar o Frontend:
 - **Navegue para /Frontend/webapp**
 - **Instale as dependências:** go mod tidy
 - **Inicie o servidor:** go run main.go
 - **O frontend estará disponível em http://localhost:2000**

### Estrututa de Diretórios:

```bash 
todo-list/
├── ONBOARDING_HELOISA/
│   ├── Backend/                      # Backend da aplicação (API em Go)
│   │   ├── api/                      # Diretório principal da API
│   │   │   ├── controllers/          # Controladores da aplicação 
│   │   │   ├── entities/             # Modelos de dados da aplicação
│   │   │   ├── routes/               # Definição das rotas da API
│   │   │   ├── swagger.json          # Documentação da API em formato Swagger
│   │   ├── Dockerfile                # Configuração do Docker para o backend
│   │   ├── go.mod                    # Dependências do Go
│   │   ├── go.sum                    # Checksum das dependências
│   │   ├── main.go                   # Arquivo principal para execução da API
│   │   ├── todo.db                   # Banco de dados SQLite 
│   └── Frontend/                     # Frontend da aplicação (Next.js)
│   │   ├── .next/                    # Arquivos gerados pelo Next.js (build)
│   │   ├── node_modules/             # Dependências do projeto (yarn)
│   │   ├── public/                   # Imagens
│   │   ├── src/                      # Código-fonte da aplicação
│   │   │   ├── actions/              # Ações para manipulação de estados e dados da API
│   │   │   ├── app/                  # Configurações iniciais do Next.js
│   │   │   ├── components/           # Componentes reutilizáveis como botões, listas e formulários
│   │   │   ├── hooks/                # Hooks personalizados
│   │   │   ├── lib/                  # Biblioteca para Styled Components
│   │   │   ├── providers/            # Contextos e provedores globais
│   │   │   ├── theme/                # Definições de temas e estilos globais
│   │   │   ├── types/                # Tipagens 
│   │   │   │   └── styled.d.ts       # Tipagem para Styled Components
│   │   ├── .eslintrc.json            # Configuração do ESLint
│   │   ├── .gitignore                # Arquivos e pastas ignorados pelo Git
│   │   ├── Dockerfile                # Configuração do Docker para o frontend
│   │   ├── next-env.d.ts             # Tipagens padrão do Next.js
│   │   ├── next.config.ts            # Configurações do Next.js
│   │   ├── package-lock.json         # Lockfile para dependências (yarn)
│   │   ├── package.json              # Lista de dependências e scripts do projeto
│   │   ├── registry.tsx              # Registro de componentes ou rotas
│   │   ├── tsconfig.json             # Configuração do TypeScript
│   │   └── yarn.lock                 # Lockfile para dependências do Yarn
├── docker-compose.yml                # Arquivo para orquestração de contêineres Docker
└──  README.MD                        # README da aplicação

```

## API Endpoints:

**Autenticação**
- POST /login - Autenticar usuário e gerar token JWT.

**Usuários**
- POST /usuarios - Criar um novo usuário.
- GET /usuarios - Listar todos os usuários.
- GET /usuarios/{usuarioId} - Buscar um usuário por ID.
- PUT /usuarios/{usuarioId} - Atualizar um usuário.
- DELETE /usuarios/{usuarioId} - Excluir um usuário.
- POST /usuarios/{usuarioId}/seguir - Seguir um usuário.
- POST /usuarios/{usuarioId}/parar-de-seguir - Parar de seguir um usuário.
- GET /usuarios/{usuarioId}/seguidores - Buscar os seguidores de um usuário.
- GET /usuarios/{usuarioId}/seguindo - Buscar os usuários que um usuário segue.
- POST /usuarios/{usuarioId}/atualizar-senha - Atualizar senha do usuário logado.

**Publicações**

- POST /publicacoes - Criar uma nova publicação.
- GET /publicacoes - Listar publicações do feed.
- GET /publicacoes/{publicacaoId} - Buscar uma publicação por ID.
- PUT /publicacoes/{publicacaoId} - Atualizar uma publicação.
- DELETE /publicacoes/{publicacaoId} - Excluir uma publicação.
- GET /usuarios/{usuarioId}/publicacoes - Buscar publicações de um usuário específico.
- POST /publicacoes/{publicacaoId}/curtir - Curtir uma publicação.
- POST /publicacoes/{publicacaoId}/descurtir - Descurtir uma publicação.