package main

import (
	"browse/app"
	"os"
)

func main() {
	app.Run(os.Getenv("ENV"))
}
