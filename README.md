
# Lightening Strike RESTful API using Go and Gin
Based on [https://go.dev/doc/tutorial/web-service-gin](https://go.dev/doc/tutorial/web-service-gin)

## Get dependencies
`go get .`

## Run the code
`go run .`

## View all lightning stikes
`curl http://localhost:8080/strikes`

## Add a lightning strike
`curl http://localhost:8080/strikes --include --header "Content-Type: application/json" --request "POST" --data '{"id": "4","strikeTime": "2021-11-10T11:00:00Z","xCoord": 52.961441,"yCoord": 1.067521}'`
