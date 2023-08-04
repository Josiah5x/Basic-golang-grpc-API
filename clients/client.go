package main

import (
	"context"
	"fmt"
	"grpc/example/interface/proto"
	"log"
	"net/http"
	"strconv"

	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
	"google.golang.org/grpc/reflection"
)

func main() {
	conn, err := grpc.Dial(":8080", grpc.WithTransportCredentials(insecure.NewCredentials()), grpc.WithBlock())
	if err != nil {
		log.Fatalln(err)
	}
	// creating the client server connection
	srv := grpc.NewServer()
	client := proto.NewAddServiceClient(conn)
	reflection.Register(srv)

	//using
	mux := http.NewServeMux()
	mux.HandleFunc("/additions", func(w http.ResponseWriter, r *http.Request) {
		// handling the request coming from the User
		num1, _ := strconv.Atoi(r.FormValue("num1"))
		num2, _ := strconv.Atoi(r.FormValue("num2"))

		//the user request is been response by multiplying the num1 and num2 the define standard if not it return error
		response, _ := client.Add(context.Background(), &proto.Request{A: int64(num1), B: int64(num2)})

		// Show the result to user
		fmt.Fprintf(w, "Result: %v\n", response)
	})
	mux.HandleFunc("/multiplication", func(w http.ResponseWriter, r *http.Request) {
		// handling the request coming from the User
		num1, _ := strconv.Atoi(r.FormValue("num1"))
		num2, _ := strconv.Atoi(r.FormValue("num2"))

		//the user request is been response by multiplying the num1 and num2 the define standard if not it return error
		response, _ := client.Multiply(context.Background(), &proto.Request{A: int64(num1), B: int64(num2)})

		// Show the result to user
		fmt.Fprintf(w, "Result: %v\n", response)
	})
	// this is HTTP-1.1 server serving on port 3333
	fmt.Println("Server Started at port 3333")
	http.ListenAndServe(":3333", mux)

}
