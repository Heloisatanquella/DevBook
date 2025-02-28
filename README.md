
# 🚀 **DevBook**

Uma aplicação web para gerenciar suas publicações e seu usuário, com a funcionalidade de criar, visualizar, atualizar e deletar publicações e usuários.

## 📝 **Índice**
1. [📋 Sobre o Projeto](#sobre-o-projeto)
2. [🛠️ Tecnologias Utilizadas](#tecnologias-utilizadas)
3. [⚙️ Funcionalidades](#funcionalidades)
4. [🚀 Instalação](#instalação)
5. [📌 API Endpoints](#api-endpoints)

---

## **📋 Sobre o Projeto** <a name="sobre-o-projeto"></a>
Este é um projeto de gerenciamento de publicações e usuário, desenvolvido para fins de aprendizado e prática com tecnologias web. Ele inclui:
- Interface responsiva para desktop e dispositivos móveis.
- CRUD completo de usuários e publicações.
- Integração com uma API para gerenciamento de dados.
- Autenticação via JWT.
- Documentação da API utilizando Swagger.

---

## **🛠️ Tecnologias Utilizadas** <a name="tecnologias-utilizadas"></a>
### **Frontend**
- **Linguagem:** [Go](https://go.dev/learn/)
- **Comunicação com a API:** [jQuery - Ajax](https://go.dev/learn/](https://api.jquery.com/jQuery.ajax/ )
- **Estilização:** [Bootstrap](https://getbootstrap.com/)

### **Backend**
- **Linguagem:** [Go](https://go.dev/learn/)
- **Banco de Dados:** [MySQL](https://dev.mysql.com/doc/)

### **Outras Ferramentas**
- **Documentação da API:** Swagger.
- **Autenticação:** JWT.

---

## **⚙️ Funcionalidades** <a name="funcionalidades"></a>
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

## **🚀 Instalação** <a name="instalação"></a>
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

 4. Configure as variáveis de ambiente:
  - Crie um arquivo `.env` com as credenciais do banco de dados MySQL e a chave secreta JWT. Um exemplo de arquivo `.env` seria:

  Backend
  ```bash
    DB_USUARIO=usuario_db
    DB_SENHA=senha_db
    DB_NOME=db_name
    
    API_PORT=porta_que_roda_sua_api
    
    SECRET_KEY=sua_secret_key
  ```

 Frontend
  ```bash
    API_URL=url_sua_api
    APP_PORT=porta_que_roda_seu_front

    HASH_KEY=sua_hash_key
    BLOCK_KEY=sua_block_key
  ```

## 📌 API Endpoints: <a name="api-endpoints"></a>

| **Método** | **Endpoint**                                    | **Descrição**                                |
|------------|-------------------------------------------------|----------------------------------------------|
| **Autenticação** |                                                 |                                        |
| POST       | /login                                          | Autenticar usuário e gerar token JWT.        |
| **Usuários**       |                                                 |                                      | 
| POST       | /usuarios                                       | Criar um novo usuário.                       |
| GET        | /usuarios                                       | Listar todos os usuários.                    |
| GET        | /usuarios/{usuarioId}                           | Buscar um usuário por ID.                    |
| PUT        | /usuarios/{usuarioId}                           | Atualizar um usuário.                        |
| DELETE     | /usuarios/{usuarioId}                           | Excluir um usuário.                          |
| POST       | /usuarios/{usuarioId}/seguir                    | Seguir um usuário.                           |
| POST       | /usuarios/{usuarioId}/parar-de-seguir           | Parar de seguir um usuário.                  |
| GET        | /usuarios/{usuarioId}/seguidores                | Buscar os seguidores de um usuário.          |
| GET        | /usuarios/{usuarioId}/seguindo                  | Buscar os usuários que um usuário segue.     |
| POST       | /usuarios/{usuarioId}/atualizar-senha           | Atualizar senha do usuário logado.           |
| **Publicações**    |                                                 |                                      |
| POST       | /publicacoes                                    | Criar uma nova publicação.                   |
| GET        | /publicacoes                                    | Listar publicações do feed.                  |
| GET        | /publicacoes/{publicacaoId}                     | Buscar uma publicação por ID.                |
| PUT        | /publicacoes/{publicacaoId}                     | Atualizar uma publicação.                    |
| DELETE     | /publicacoes/{publicacaoId}                     | Excluir uma publicação.                      |
| GET        | /usuarios/{usuarioId}/publicacoes               | Buscar publicações de um usuário específico. |
| POST       | /publicacoes/{publicacaoId}/curtir              | Curtir uma publicação.                       |
| POST       | /publicacoes/{publicacaoId}/descurtir           | Descurtir uma publicação.                    |
