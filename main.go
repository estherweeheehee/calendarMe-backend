package main

import (
	"calendarme-backend/config"
	"calendarme-backend/routes"
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	// "log"
	// "fmt"
	// "os"
	
)

func main() {
	// Set the router as the default one shipped with Gin
	router := gin.Default()
	router.Use(cors.Default())
	config.Connect()
	routes.NoteRoute(router)

	// Start and run the server
// 	port, err := getPort()
//   	if err != nil {
//     log.Fatal(err)
//   }
	// router.Run(port)
	router.Run(":3000")
}

// func getPort() (string, error) {
// 	// the PORT is supplied by Heroku
// 	port := os.Getenv("PORT")
// 	if port == "" {
// 	  return "", fmt.Errorf("$PORT not set")
// 	}
// 	return ":" + port, nil
//   }