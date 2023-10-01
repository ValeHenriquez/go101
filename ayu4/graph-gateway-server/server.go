package main

import (
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/99designs/gqlgen/graphql/handler"
	"github.com/99designs/gqlgen/graphql/playground"
	"github.com/ValeHenriquez/graph-gateway-server/config"
	"github.com/ValeHenriquez/graph-gateway-server/graph"
	"github.com/joho/godotenv"
)

const defaultPort = "8080"

func main() {

	fmt.Println("API Gateway starting...")
	godotenv.Load()
	fmt.Println("Loaded env variables...")

	fmt.Println("Configuring RabbitMQ Connection...")
	config.SetupRabbitMQ()
	fmt.Println("RabbitMQ Connection configured...")

	port := os.Getenv("PORT")
	if port == "" {
		port = defaultPort
	}

	srv := handler.NewDefaultServer(graph.NewExecutableSchema(graph.Config{Resolvers: &graph.Resolver{}}))

	http.Handle("/", playground.Handler("GraphQL playground", "/query"))
	http.Handle("/query", srv)

	log.Printf("connect to http://localhost:%s/ for GraphQL playground", port)
	log.Fatal(http.ListenAndServe(":"+port, nil))
}










