package main

import (
	"goproxy/app"
	"os"
)

func main() {
	app.Run(os.Getenv("env"))
}
