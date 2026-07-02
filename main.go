package main

import (
	"log"
	"os"

	"go-graphql-api/db"
	"go-graphql-api/graph"

	"github.com/99designs/gqlgen/graphql/handler"
	"github.com/99designs/gqlgen/graphql/playground"
	"github.com/gofiber/fiber/v3"
	"github.com/gofiber/fiber/v3/middleware/adaptor"
	"github.com/joho/godotenv"
)

func main() {
	db.ConnectDB()

	if err := godotenv.Load(); err != nil {
		log.Println("Error reading platform environment")
	}
	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}

	app := fiber.New()

	// 3. Assemble GQL schema with database pointer injection
	srv := handler.NewDefaultServer(graph.NewExecutableSchema(graph.Config{
		Resolvers: &graph.Resolver{DB: db.DB},
	}))

	// 4. Bind Endpoint Handlers
	app.Get("/", adaptor.HTTPHandler(playground.Handler("Library Dashboard Workspace", "/query")))
	app.Post("/books", adaptor.HTTPHandler(srv))

	log.Printf("Library Network Engine processing tasks on http://localhost:%s/", port)
	log.Fatal(app.Listen(":" + port))
}
