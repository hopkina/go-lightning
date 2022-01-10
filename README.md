
# Lightning Strike RESTful API using Go and Gin
Based on [https://go.dev/doc/tutorial/web-service-gin](https://go.dev/doc/tutorial/web-service-gin)
And [Build a REST API with Golang from scratch: PostgreSQL with Gorm and Gin Web Framework](https://dc1888.medium.com/build-a-rest-api-with-golang-from-scratch-postgresql-with-gorm-and-gin-web-framework-3d3f95ccf2e7)

## Get dependencies
`go get .`

## Run the code
`go run .`

## View all lightning stikes
`curl http://localhost:8080/strikes`

## Add a lightning strike
`curl http://localhost:8080/strikes --include --header "Content-Type: application/json" --request "POST" --data '{"id": "4","strikeTime": "2021-11-10T11:00:00Z","xCoord": 52.961441,"yCoord": 1.067521}'`
