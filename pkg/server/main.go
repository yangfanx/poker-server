package server

import (
	"log"

	"github.com/gin-gonic/gin"
	"github.com/yangfanx/poker-server/internal/handlers"
)

var host, port, gqlPath, gqlPgPath string
var isPgEnabled bool

func init() {
	host = "localhost"
	port = "7777"
	gqlPath = "/graphql"
	gqlPgPath = "/"
	isPgEnabled = true
}

// Run spins up the server
func Run() {
	endpoint := "http://" + host + ":" + port

	r := gin.Default()

	// Handlers
	// Simple keep-alive/ping handler
	r.GET("/ping", handlers.Ping())

	// GraphQL handlers
	// Playground handler
	if isPgEnabled {
		r.GET(gqlPgPath, handlers.PlaygroundHandler(gqlPath))
		log.Println("GraphQL Playground @ " + endpoint + gqlPgPath)
	}
	r.POST(gqlPath, handlers.GraphqlHandler())
	log.Println("GraphQL @ " + endpoint + gqlPath)

	// Run the server
	// Inform the user where the server is listening
	log.Println("Running @ " + endpoint)
	// Print out and exit(1) to the OS if the server cannot run
	log.Fatalln(r.Run(host + ":" + port))
}
