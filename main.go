package main

import (
	"pinpong.co/rest-go/deliveries/routes"
)

func main() {
	r := routes.SetupRouter()
	r.Run()
}
