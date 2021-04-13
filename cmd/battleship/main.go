package main

import "battleship/internal/app"

const configPath = "configs/main"

func main() {
	app.Run(configPath)
}