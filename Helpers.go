package main

import (
	rl "github.com/gen2brain/raylib-go/raylib"
	"github.com/ungerik/go3d/quaternion"
	"github.com/ungerik/go3d/vec3"
)

func Add(v1, v2 rl.Vector3) rl.Vector3 {
	v3 := rl.Vector3{
		X: v1.X + v2.X,
		Y: v1.Y + v2.Y,
		Z: v1.Z + v2.Z,
	}
	return v3
}

func AddMultiple(vs []rl.Vector3) rl.Vector3 {
	v3 := rl.Vector3{}
	for _, v := range vs {
		Add(v3, v)
	}
	return v3
}

func Subtract(v1, v2 rl.Vector3) rl.Vector3 {
	v3 := rl.Vector3{
		X: v1.X - v2.X,
		Y: v1.Y - v2.Y,
		Z: v1.Z - v2.Z,
	}
	return v3
}

func Multiply(v1 rl.Vector3, num float32) rl.Vector3 {
	v3 := rl.Vector3{
		X: v1.X * num,
		Y: v1.Y * num,
		Z: v1.Z * num,
	}
	return v3
}

func Dot(v1, v2 rl.Vector3) float32 {
	d := v1.X*v2.X + v1.Y*v2.Y + v1.Z*v2.Z
	return d
}

func Cross(v1, v2 rl.Vector3) rl.Vector3 {
	v3 := rl.Vector3{
		X: v1.Y*v2.Z - v1.Z*v2.Y,
		Y: v1.Z*v2.X - v1.X*v2.Z,
		Z: v1.X*v2.Y - v1.Y*v2.X,
	}
	return v3
}

func RlVecTogo3dVec(v rl.Vector3) vec3.T {
	return vec3.T{v.X, v.Y, v.Z}
}

func RlQuatTogo3dQuat(q rl.Quaternion) quaternion.T {
	return quaternion.T{q.X, q.Y, q.Z, q.W}
}

func Go3DVecToRlVec(v vec3.T) rl.Vector3 {
	return rl.Vector3{v[0], v[1], v[2]}
}

func Go3dQuatToRlQuat(q quaternion.T) rl.Quaternion {
	return rl.Quaternion{q[0], q[1], q[2], q[3]}
}

func CameraRotation(currentForward rl.Vector3) rl.Quaternion {
	go3dV := RlVecTogo3dVec(currentForward)
	q := quaternion.Vec3Diff(&vec3.T{0, 0, 1}, &go3dV)
	return Go3dQuatToRlQuat(q)
}

//ToDO. wtf,janis delete this
func CameraRight() rl.Vector3 {
	q := CameraRotation(cameraForward)
	v := RotateVectorByQuaternion(rl.Vector3{1, 0, 0}, q)
	return v
}

func CameraLeft() rl.Vector3 {
	q := CameraRotation(cameraForward)
	v := RotateVectorByQuaternion(rl.Vector3{-1, 0, 0}, q)
	return v
}

func CameraUp() rl.Vector3 {
	q := CameraRotation(cameraForward)
	v := RotateVectorByQuaternion(rl.Vector3{0, 1, 0}, q)
	return v
}

func CameraDown() rl.Vector3 {
	q := CameraRotation(cameraForward)
	v := RotateVectorByQuaternion(rl.Vector3{0, -1, 0}, q)
	return v
}

func CameraBack() rl.Vector3 {
	q := CameraRotation(cameraForward)
	v := RotateVectorByQuaternion(rl.Vector3{0, 0, -1}, q)
	return v
}

func QuaternionFromEuler(euler rl.Vector3) rl.Quaternion {
	//ToDo: Mistake probably here
	q := quaternion.FromEulerAngles(euler.Y*rl.Rad2deg, euler.X*rl.Rad2deg, euler.Z*rl.Rad2deg)

	return Go3dQuatToRlQuat(q)
}

func RotateVectorByQuaternion(v rl.Vector3, q rl.Quaternion) rl.Vector3 {

	go3dQ := RlQuatTogo3dQuat(q)
	go3dV := RlVecTogo3dVec(v)
	result := go3dQ.RotatedVec3(&go3dV)

	return Go3DVecToRlVec(result)
}
