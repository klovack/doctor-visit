package main

//go:generate sqlboiler --wipe psql

import (
	"fmt"
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/graphql-go/graphql"
	"github.com/graphql-go/handler"
	"github.com/klovack/doctor-visit/db"
	"github.com/klovack/doctor-visit/gql"
	"github.com/klovack/doctor-visit/server"
	_ "github.com/lib/pq"
)

func main() {
	// Initialize API and close the connection to the db when the program is finished
	router, database := initializeAPI()
	defer database.Close()

	fmt.Println("Preparation is done.")
	fmt.Println("Launching server on port 6789...")

	log.Fatal(http.ListenAndServe(":6789", router))
}

func initializeAPI() (*gin.Engine, *db.DB) {
	// Create a new Router
	router := gin.Default()

	// Create a connection to the database
	database, err := db.New(
		db.NewConnectionConfig("172.20.0.1", 5432, "postgres", "my-super-secret", "doctor_visit"),
	)
	if err != nil {
		log.Fatalln(err)
	}

	// // Create root query for graphqll
	rootQuery := gql.NewRoot(database)

	schema, err := graphql.NewSchema(
		graphql.SchemaConfig{Query: rootQuery.Query},
	)
	if err != nil {
		fmt.Println("Failed to create schema", err)
	}

	h := handler.New(&handler.Config{
		Schema:     &schema,
		Pretty:     true,
		GraphiQL:   true,
		Playground: true,
	})

	s := server.NewServer(h)

	router.POST("/graphql", s.GraphQL)

	return router, database
}
