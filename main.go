package main

import (
	"fmt"

	"github.com/TheAnswer16/rapper_phrases_generator/api"

	"github.com/joho/godotenv"
)

func main() {
	err := godotenv.Load()
	if err != nil {
		fmt.Println("Erro ao carregar as variáveis de ambiente:", err)
		return
	}

	api.Init()

	if err != nil {
		fmt.Println("Erro ao iniciar o servidor:", err)
	}

}
