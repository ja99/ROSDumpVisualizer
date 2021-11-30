package main

import (
	"encoding/binary"
	"fmt"
	"github.com/aler9/goroslib"
	"github.com/aler9/goroslib/pkg/msgs/sensor_msgs"
	rl "github.com/gen2brain/raylib-go/raylib"
	"math"
	"os"
	"os/signal"
)

var (
	cameraSubscriber *goroslib.Subscriber
	lidarSubscriber  *goroslib.Subscriber
)

func onCameraMessage(msg *sensor_msgs.Image) {
	fmt.Printf("Incoming: %+v\n", msg)
	image.image = rl.Image{
		Width:   int32(msg.Width),
		Height: int32(msg.Height),
		Mipmaps: 0,
		Format:  rl.UncompressedR8g8b8a8,
	}

	for xi := int32(0); xi < int32(msg.Width); xi++ {
		for yi := int32(0); yi < int32(msg.Width); yi++ {
			val := xi+yi*image.image.Width
			color := rl.Color{msg.Data[val],msg.Data[val+1],msg.Data[val+1],255}
			image.lock.Lock()
			rl.ImageDrawPixel(&image.image,xi,yi,color)
			image.lock.Unlock()
		}
	}
}

func Float32frombytes(bytes []byte) float32 {
	bits := binary.LittleEndian.Uint32(bytes)
	float := math.Float32frombits(bits)
	return float
}

func onLidarMessage(msg *sensor_msgs.PointCloud2) {
	fmt.Printf("Incoming: %+v\n", msg)
	x := Float32frombytes(msg.Data[:4])
	y := Float32frombytes(msg.Data[4:8])
	z := Float32frombytes(msg.Data[8:12])

	v:= rl.Vector3{x,y,z}

	points.lock.Lock()
	points.points = append(points.points, v)
	points.lock.Unlock()

}

func HandleConnection() {
	// create a node and connect to the master
	n, err := goroslib.NewNode(goroslib.NodeConf{
		Name:          "RosDumpVisualizer",
		MasterAddress: "127.0.0.1:11311",
	})
	if err != nil {
		panic(err)
	}
	defer n.Close()

	// create a subscriber
	cameraSubscriber, err = goroslib.NewSubscriber(goroslib.SubscriberConf{
		Node:     n,
		Topic:    "/pylon_camera_node/image_raw",
		Callback: onCameraMessage,
	})
	if err != nil {
		panic(err)
	}
	defer cameraSubscriber.Close()

	// create a subscriber
	lidarSubscriber, err = goroslib.NewSubscriber(goroslib.SubscriberConf{
		Node: n,
		// ToDo: Make sure this is the correct topic
		Topic:    "/velodyne_points",
		Callback: onLidarMessage,
	})
	if err != nil {
		panic(err)
	}
	defer lidarSubscriber.Close()

	// wait for CTRL-C
	c := make(chan os.Signal, 1)
	signal.Notify(c, os.Interrupt)
	<-c
}
