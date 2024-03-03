package api

import (
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/TheAnswer16/rapper_phrases_generator/controllers"
)

func Init() {

	http.HandleFunc("/api/get-phrase", controllers.GetPhrase)

	StartServer()
}

func StartServer() {

	port := os.Getenv("APP_PORT")
	fmt.Println("Starting server...")
	log.Fatal(http.ListenAndServe(fmt.Sprintf(":%s", port), nil))

}
