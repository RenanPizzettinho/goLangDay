# goLangDay
Repositório para aprendizado da linguagem Go.

## Rest API - Mux Routers and Postgres 

Uso de [github.com/gorilla/mux](https://github.com/gorilla/mux) com [github.com/lib/pq](https://github.com/lib/pq) para criação de api simples usando banco postgres.
Executar o comando `go run server.go` e acessar `localhost:8080/users/all`.

## Rest API - Mux Routers and Bongo

Rest API CRUD usando [github.com/gorilla/mux](https://github.com/gorilla/mux) para o gerenciamento de rotas e [github.com/go-bongo/bongo](https://github.com/go-bongo/bongo) para conexão com banco MongoDb.
Executar o comando `go run main.go usuario.go usuarioResource.go usuarioService.go util.go dbConnection.go` no diretorio `/src`.

### APIs 
* `localhost:8080/usuario` - GET - POST
* `localhost:8080/usuario/<id>` - GET - PUT - DELETE 

### Modelo de documento
```json
{
	"nome": "Renan Silvano",
	"login": "renan.silvano",
	"senha": "123456"
}
```
