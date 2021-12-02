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
	image     map[[2]int32]rl.Color
	lock      sync.RWMutex
	rowLength int
}

var (
	camera        = rl.Camera3D{}
	cameraForward = rl.Vector3{0, 0, 1}
	points        = Points{[]rl.Vector3{}, sync.RWMutex{}}
	image         = Pic{make(map[[2]int32]rl.Color), sync.RWMutex{}, 1}
)

var (
	cameraMoveSpeed = float32(0.1)
	cameraLookSpeed = float32(0.0001)
)

func Init() {
	rl.InitWindow(1600, 1100, "raylib [core] example - 3d mode")

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
	points.lock.Lock()
	for _, point := range points.points {
		//rl.DrawCube(point, 1.0, 1.0, 1.0, rl.Red)
		rl.DrawCube(point, 0.05, 0.05, 0.05, rl.Blue)
	}
	points.lock.Unlock()
}

func DrawImage() {
	drawsize := int32(3)

	image.lock.RLock()
	for xi := int32(0); xi < int32(image.rowLength); xi += drawsize {
		for yi := int32(0); yi < int32(len(image.image)/image.rowLength); yi += drawsize {
			rl.DrawPixel(xi/drawsize, yi/drawsize, image.image[[2]int32{xi, yi}])
		}
	}
	image.lock.RUnlock()

}
