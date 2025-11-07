package main

import (
	"log"
	"os"

	"github.com/LucasPaulo001/ToDo-GO/src/routes"
	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
)

func main(){
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Erro ao carregar arquivo .env")
	}

	router := gin.Default()
	routes.InitRoutes(router)

	PORT := os.Getenv("PORT")
	router.Run(":"+PORT)
}