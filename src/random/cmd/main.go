package main

import (
	"os"
	"random/app"
)

func main() {
	app.Run(os.Getenv("env"))
}
