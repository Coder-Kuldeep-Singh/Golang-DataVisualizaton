package main

import (
	"Dashboard/routes"
)

func main() {
	//setup routes
	r := routes.SetupRouter()

	// running
	r.Run(":8080")

}
