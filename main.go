package main

import rl "github.com/gen2brain/raylib-go/raylib"

func main() {
	go HandleConnection()
	Init()
	rl.CloseWindow()
}
