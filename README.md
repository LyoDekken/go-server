# Go Server API

Essa é uma API escrita em Go para criar e gerenciar usuários em um banco de dados MySQL.

## Configuração

Antes de executar a API, é necessário configurar o banco de dados. Para isso, é necessário criar um banco de dados MySQL e configurar as variáveis de ambiente com as informações de conexão:

## Infos
 - Go 1.16 ou superior
 - MySQL 8.0 ou superior
 - Instalação

Para instalar a aplicação, execute os seguintes comandos:

shell
Copy code
$ git clone https://github.com/seu-usuario/go-server.git
$ cd go-server
$ go build
Isso irá criar um executável go-server na pasta raiz do projeto.

Configuração
A aplicação utiliza as seguintes variáveis de ambiente para se configurar:

DB_USER: o nome de usuário do banco de dados
DB_PASSWORD: a senha do banco de dados
DB_HOST: o endereço do host do banco de dados
DB_PORT: a porta do banco de dados
DB_NAME: o nome do banco de dados
Certifique-se de configurar essas variáveis de acordo com o seu ambiente antes de executar a aplicação.

Uso
Para executar a aplicação, execute o seguinte comando:

shell
Copy code
$ ./go-server
Isso irá iniciar o servidor na porta 8080.

Para criar um novo usuário, faça uma solicitação POST para o endpoint /users com os seguintes campos no corpo da solicitação:

name: o nome do usuário
email: o e-mail do usuário
password: a senha do usuário
role: o papel do usuário (pode ser admin ou user)
Exemplo de solicitação:

bash
Copy code
POST /users HTTP/1.1
Content-Type: application/json

{
    "name": "João da Silva",
    "email": "joao.silva@example.com",
    "password": "senha123",
    "role": "user"
}
Se o usuário for criado com sucesso, o servidor irá retornar um código de status 201 Created.

Contribuindo
Sinta-se à vontade para enviar pull requests para este projeto. Toda contribuição é bem-vinda!

Licença
Este projeto está licenciado sob a licença MIT. Consulte o arquivo LICENSE para obter mais detalhes.