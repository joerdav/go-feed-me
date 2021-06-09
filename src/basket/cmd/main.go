package main

import (
	"basket/app"
	"os"
)

func main() {
	app.Run(os.Getenv("env"))
}
