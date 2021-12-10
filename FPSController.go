package main

import rl "github.com/gen2brain/raylib-go/raylib"

var (
	cameraForward      = rl.Vector3{0, 0, 1}
	cameraMoveSpeed    = float32(0.1)
	cameraLookSpeed    = float32(0.0001)
	cursorCatched      = false
	f1PressedLastFrame = false
)

func CameraControls() {
	if rl.IsKeyDown(rl.KeyW) {
		camera.Position = rl.Vector3Add(camera.Position, rl.Vector3Multiply(cameraForward, cameraMoveSpeed))
	}
	if rl.IsKeyDown(rl.KeyD) {
		camera.Position = rl.Vector3Add(camera.Position, rl.Vector3Multiply(CameraLeft(), cameraMoveSpeed))
	}
	if rl.IsKeyDown(rl.KeyS) {
		camera.Position = rl.Vector3Add(camera.Position, rl.Vector3Multiply(CameraBack(), cameraMoveSpeed))

	}
	if rl.IsKeyDown(rl.KeyA) {
		camera.Position = rl.Vector3Add(camera.Position, rl.Vector3Multiply(CameraRight(), cameraMoveSpeed))
	}
	if rl.IsKeyDown(rl.KeyQ) {
		camera.Position = rl.Vector3Add(camera.Position, rl.Vector3Multiply(CameraDown(), cameraMoveSpeed))
	}
	if rl.IsKeyDown(rl.KeyE) {
		camera.Position = rl.Vector3Add(camera.Position, rl.Vector3Multiply(CameraUp(), cameraMoveSpeed))

	}

	if rl.IsKeyDown(rl.KeyF1) {
		if !f1PressedLastFrame {
			if cursorCatched {
				rl.EnableCursor()
			} else {
				rl.DisableCursor()
			}

			cursorCatched = !cursorCatched
		}
		f1PressedLastFrame = true
	} else {
		f1PressedLastFrame = false
	}

	if cursorCatched {
		mouseDelta := rl.GetMouseDelta()
		yRotSign := float32(1)
		if rl.Vector3DotProduct(cameraForward, rl.Vector3{0, 0, 1}) < 0 {
			yRotSign = -1
		}
		cameraRotEuler := rl.Vector3{mouseDelta.Y * yRotSign * cameraLookSpeed, -mouseDelta.X * cameraLookSpeed, 0}
		cameraRotQ := QuaternionFromEuler(cameraRotEuler)
		cameraForward = RotateVectorByQuaternion(cameraForward, cameraRotQ)
		cameraForward = rl.Vector3Normalize(cameraForward)

		camera.Target = Add(camera.Position, cameraForward)
	}
}
