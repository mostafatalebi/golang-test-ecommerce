# Ecommerce test app

This project is for testing and demonstration purposes. It uses Golang.

## Usage
Simply run the following command to run the server on port `8010`.
The command needs to run in the project's root. 
`golang
go mod tidy && go run .
`

### API
The following endpoints are exposed:
```shell
GET /products/{id}
GET /products
GET /products?category=name
POST /products JSON
```

The postman collection in the repository's root contains all endpoints. You can use it to test and check the app.