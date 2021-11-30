package main

import (
	rl "github.com/gen2brain/raylib-go/raylib"
	"sync"
)

type Points struct {
	points []rl.Vector3
	lock   sync.RWMutex
}

type Pic struct {
	image rl.Image
	lock  sync.RWMutex
}

var (
	camera        = rl.Camera3D{}
	cameraForward = rl.Vector3{0, 0, 1}
	points        = Points{[]rl.Vector3{}, sync.RWMutex{}}
	image         = Pic{rl.Image{}, sync.RWMutex{}}
)

var (
	cameraMoveSpeed = float32(0.1)
	cameraLookSpeed = float32(0.0001)
)

func main() {
	go HandleConnection()
	Init()
	rl.CloseWindow()
}

func Init() {
	rl.InitWindow(1000, 1000, "raylib [core] example - 3d mode")

	camera.Position = rl.NewVector3(0.0, 5.0, -10.0)
	camera.Target = rl.NewVector3(0.0, 0.0, 0.0)
	camera.Up = rl.NewVector3(0.0, 1.0, 0.0)
	camera.Fovy = 70
	camera.Projection = rl.CameraPerspective

	rl.SetTargetFPS(75)
	//rl.SetCameraMode(camera, rl.CameraFirstPerson)

	for !rl.WindowShouldClose() {
		rl.BeginDrawing()
		rl.ClearBackground(rl.RayWhite)
		Update()
		rl.EndDrawing()
	}
}

func Update() {
	rl.BeginMode3D(camera)
	DrawCubes()
	rl.DrawGrid(100, 1.0)
	rl.EndMode3D()

	DrawImage()
	rl.DrawFPS(10, 10)

}



func DrawCubes() {
	points.lock.RLock()
	for _, point := range points.points {
		//rl.DrawCube(point, 1.0, 1.0, 1.0, rl.Red)
		rl.DrawCubeWires(point, 0.1, 0.1, 0.1, rl.Blue)
	}
	points.lock.RUnlock()
}

func DrawImage() {
	image.lock.RLock()
	texture := rl.LoadTextureFromImage(&image.image)
	points.lock.RUnlock()

	rl.DrawTexture(texture, 0,0,rl.Color{0,0,0,0})
}

