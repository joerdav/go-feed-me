package main

import (
	"details/app"
	"os"
)

func main() {
	app.Run(os.Getenv("ENV"))
}
