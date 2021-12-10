package main

import (
	rl "github.com/gen2brain/raylib-go/raylib"
	"strconv"
	"sync"
)

type Points struct {
	points []rl.Vector3
	lock   sync.RWMutex
}

// Set to 1920x1200 for real camera and 1024x640 for simulator
const picWidth = int32(1920)
const picHeight = int32(1200)

//To make the camera image smaller
const drawEveryXThPixel = int32(5)

type Pic struct {
	image     [picWidth][picHeight]rl.Color
	lock      sync.RWMutex
	rowLength int
}

var (
	camera = rl.Camera3D{}
	points = Points{[]rl.Vector3{}, sync.RWMutex{}}
	image  = Pic{[picWidth][picHeight]rl.Color{}, sync.RWMutex{}, 1}
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

	//ToDo: Display HZ of camerafeed
	//ToDo: Display HZ of Lidarfeed

	for !rl.WindowShouldClose() {
		rl.BeginDrawing()
		rl.ClearBackground(rl.Color{45, 47, 48, 255})
		Update()
		rl.EndDrawing()
	}
}

func Update() {
	CameraControls()
	rl.BeginMode3D(camera)
	DrawCubes()
	rl.DrawGrid(100, 1.0)
	rl.EndMode3D()

	DrawImage()
	rl.DrawText("camera ms: "+strconv.Itoa(int(cameraMs)), 1350, 10, 30, rl.Blue)
	rl.DrawText("lidar ms: "+strconv.Itoa(int(lidarMs)), 1350, 50, 30, rl.Blue)
	rl.DrawText("Press F1 to toggle mouse-catching and FPS controls", 500, 10, 20, rl.White)
	rl.DrawFPS(10, 10)

}

func DrawCubes() {
	points.lock.RLock()
	for _, point := range points.points {
		rl.DrawCube(point, 0.05, 0.05, 0.05, rl.Blue)
	}
	points.lock.RUnlock()
}

func DrawImage() {

	image.lock.RLock()
	for xi := int32(0); xi < int32(len(image.image)); xi += drawEveryXThPixel {
		for yi := int32(0); yi < int32(len(image.image[0])); yi += drawEveryXThPixel {
			rl.DrawPixel(xi/drawEveryXThPixel, yi/drawEveryXThPixel, image.image[xi][yi])
		}
	}
	image.lock.RUnlock()

}
