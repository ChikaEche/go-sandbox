# go-sandbox

# climate data

The purpose of this application is to show how to use grpc to send information between servers
This application uses three go applications and a proto file to define the methods to be used in getting data
- The client application gets information from the server application via grpc
- The main go application gets information from the client application via grpc and exposes that information via REST

# start locally

- Make sure go is installed on your local machine
- clone this repo https://github.com/ChikaEche/go-sandbox.git
- cd go-sandbox
- Run the command go run ./main.go
- Use postman or your browser to make a request to localhost:8080
