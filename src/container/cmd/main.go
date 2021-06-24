package main

import (
	"container/app"
	"os"
)

func main() {
	app.Run(os.Getenv("ENV"))
}
