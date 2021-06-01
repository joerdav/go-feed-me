package main

import (
	"content/app"
	"os"
)

func main() {
	app.Run(os.Getenv("env"))
}
