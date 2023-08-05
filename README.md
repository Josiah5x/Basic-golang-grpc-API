# Basic-golang-grpc-API

### *gRPC stand for Google Remote Procedure Call*
Is a framework developed by Google to handle remote procedure calls (RPCs).

>Definition - gRPC is a modern, open source remote procedure call (RPC) framework that can run anywhere. This is something that we use within distributed systems that allow us to communicate between applications. More specifically, it allows us to expose methods within our application that we want other applications to be able to invoke.
#
## Before going for detailing let run the code. I beleived you already have Basic of Golang.

### Clone or pull this gitbub repo
https://github.com/Josiah5x/Basic-golang-grpc-API.git

## step1
You need to install Golang

Next, go to your directory in terminal type below code
```
go mod tidy
```
Next: Run this command on the terminal

>Server: go run main.go
>Client: go run client.go

## Testing the API
### Addition
```

curl -X POST -F "num1=2" -F "num2=6" 'http://localhost:3333/addition'
```
### Multiplication
```
curl -X POST -F "num1=2" -F "num2=6" 'http://localhost:3333/multiplication'
```
