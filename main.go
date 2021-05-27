package main

import (
	"fiber-ordering/config"
	"fiber-ordering/controller"
	"fiber-ordering/repository"
	"fiber-ordering/service"
	"fmt"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"github.com/gofiber/fiber/v2/middleware/recover"
	"log"
	"os"
	"strconv"
	"time"
)

func main() {
	log.SetFlags(log.LstdFlags | log.Lshortfile)

	reader, writer := config.ConfigureMySQL()
	server := config.ServerGen{
		Reader:      reader,
		Writer:      writer,
		Port:        config.Server.Port,
	}

	readTimeoutSecondsCount, _ := strconv.Atoi(os.Getenv("SERVER_READ_TIMEOUT"))
	app := fiber.New(fiber.Config{
		ReadTimeout: time.Second * time.Duration(readTimeoutSecondsCount),
	})
	app.Use(cors.New(), recover.New())

	// account
	accountRepo := repository.NewAccountRepo(reader, writer)
	accountService := service.NewAccountService(accountRepo)
	controller.NewAccountAPI(app, accountService)

	app.Listen(fmt.Sprintf(":%s", server.Port))
}