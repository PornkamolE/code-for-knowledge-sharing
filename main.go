package main

import (
	"log"

	"github.com/gofiber/fiber/v2"

	"my-knowledge-sharing/app/adapters/handlers"
	"my-knowledge-sharing/app/adapters/repository/memory"
	"my-knowledge-sharing/app/core/services"
)

func main() {
	repo := memory.NewUserRepoMemory()
	svc := services.NewUserService(repo)
	h := handlers.NewUserHandler(svc)

	app := fiber.New()
	app.Post("/users", h.Create)
	app.Get("/users/:id", h.GetByID)

	log.Println("Server is starting at :3000")
	if err := app.Listen(":3000"); err != nil {
		log.Fatal(err)
	}
}