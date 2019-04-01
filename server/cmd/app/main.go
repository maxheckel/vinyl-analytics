package main

import (
	"app/internal/app"
	"fmt"
)

func main() {
	fmt.Println("Running")
	application := app.NewApp()
	application.Serve()

}
