
# Lightning Strike RESTful API using Go and Gin
Based on:
* [https://go.dev/doc/tutorial/web-service-gin](https://go.dev/doc/tutorial/web-service-gin)
* [Build a REST API with Golang from scratch: PostgreSQL with Gorm and Gin Web Framework](https://dc1888.medium.com/build-a-rest-api-with-golang-from-scratch-postgresql-with-gorm-and-gin-web-framework-3d3f95ccf2e7)
* [Automatically generate RESTful API documentation with Swagger 2.0 for Go](https://golangexample.com/automatically-generate-restful-api-documentation-with-swagger-2-0-for-go/)

![Swagger](swagger.png?raw=true)

## Get dependencies
`go get .`

## Swagger documents
Initiate creation
`swag init`

Hosted at
`http://localhost:8080/swagger/index.html`

# dotenv file
Contains
```
PG_USER=<user>
PG_PASSWORD=<password>
PG_DB=<database>
```

## Run the code
`go run .`

## View all lightning stikes
`curl http://localhost:8080/api/v1/strikes`

## Add a lightning strike
`curl http://localhost:8080/api/v1/strikes --include --header "Content-Type: application/json" --request "POST" --data '{"strikeTime": "2021-11-10T11:00:00Z","xCoord": 52.961441,"yCoord": 1.067521}'`
