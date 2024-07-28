package main

import (
	"net/http"
	"os"
	"xpm-auth/controller"
	"xpm-auth/helper"
	model "xpm-auth/models"
	repository "xpm-auth/repositories/auth"
	"xpm-auth/router"
	service "xpm-auth/service/auth"

	"github.com/go-playground/validator/v10"
	"github.com/joho/godotenv"
	"github.com/rs/zerolog/log"
)

func main() {
	envError := godotenv.Load()
	if envError != nil {
		log.Fatal().Msg("Error loading .env file")
		helper.Error(envError)
	}
	db := DatabaseConnection()
	validate := validator.New()

	db.Table("users").AutoMigrate(&model.User{})

	// Repository
	userRepository := repository.NewAuthRepositoryImpl(db)

	// Service
	userService := service.NewAuthServiceImpl(userRepository, validate)

	// Controller
	userController := controller.NewAuthController(userService)

	// Router
	routes := router.NewRouter(userController)
	log.Info().Msg("Server started.")
	server := &http.Server{
		Addr:    os.Getenv("PORT"),
		Handler: routes,
	}
	err := server.ListenAndServe()
	helper.Error(err)
}
