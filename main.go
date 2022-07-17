package main

import (
	"fmt"
	"os"

	"pinpong.co/rest-go/deliveries/routes"
)

func main() {
	var port = os.Getenv("PORT")
	if port == "" {
		port = "8080"
		fmt.Println("No port in heroku" + port)
	}
	r := routes.SetupRouter()
	r.Run(":" + port)
}
