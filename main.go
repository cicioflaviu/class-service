package main

import (
	"cicio.dev/class-service/system"
	"fmt"
)

func main() {
	database := system.InitDatabase()
	defer database.Close()

	server := system.InitServer(database)
	err := server.Run(":8085")
	if err != nil {
		fmt.Println("Failed to start server, details: ", err)
	}
}