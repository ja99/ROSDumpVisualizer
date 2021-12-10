package main

import (
	"encoding/binary"
	"github.com/aler9/goroslib"
	"github.com/aler9/goroslib/pkg/msgs/sensor_msgs"
	rl "github.com/gen2brain/raylib-go/raylib"
	"math"
	"os"
	"os/signal"
)

var (
	cameraSubscriber   *goroslib.Subscriber
	lidarSubscriber    *goroslib.Subscriber
	simLidarSubscriber *goroslib.Subscriber
	cameraStopwatch    Stopwatch
	cameraMs           int64
	lidarStopwatch     Stopwatch
	lidarMs            int64
)

func onCameraMessage(msg *sensor_msgs.Image) {
	cameraMs = cameraStopwatch.ElapsedMilliseconds()
	cameraStopwatch.Start()
	//fmt.Printf("Incoming: %+v\n", msg)
	image.rowLength = int(msg.Width)

	image.lock.Lock()
	for xi := int32(0); xi < int32(msg.Width); xi++ {
		for yi := int32(0); yi < int32(msg.Height); yi++ {
			index := (xi + yi*int32(msg.Width)) * 3
			color := rl.Color{msg.Data[index], msg.Data[index+1], msg.Data[index+2], 255}
			image.image[[2]int32{xi, yi}] = color
		}
	}
	image.lock.Unlock()
}

func Float32frombytes(bytes []byte) float32 {
	bits := binary.LittleEndian.Uint32(bytes)
	float := math.Float32frombits(bits)
	return float
}

func onLidarMessage(msg *sensor_msgs.PointCloud2) {
	lidarMs = lidarStopwatch.ElapsedMilliseconds()
	lidarStopwatch.Start()

	tempPoints := []rl.Vector3{}

	for i := 0; i < len(msg.Data)-12; i += 32 {
		x := Float32frombytes(msg.Data[i+0 : i+4])
		y := Float32frombytes(msg.Data[i+4 : i+8])
		z := Float32frombytes(msg.Data[i+8 : i+12])

		v := rl.Vector3{y, z, x}

		tempPoints = append(tempPoints, v)
	}
	points.lock.Lock()
	points.points = tempPoints
	points.lock.Unlock()
}

func onSimLidarMessage(msg *sensor_msgs.PointCloud) {
	lidarMs = lidarStopwatch.ElapsedMilliseconds()
	lidarStopwatch.Start()
	tempPoints := []rl.Vector3{}
	for _, point := range msg.Points {
		v := rl.Vector3{point.Y, point.Z, point.X}
		tempPoints = append(points.points, v)
	}
	points.lock.Lock()
	points.points = tempPoints
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

	// create a subscriber
	simLidarSubscriber, err = goroslib.NewSubscriber(goroslib.SubscriberConf{
		Node: n,
		// ToDo: Make sure this is the correct topic
		Topic:    "/velodyne",
		Callback: onSimLidarMessage,
	})
	if err != nil {
		panic(err)
	}
	defer simLidarSubscriber.Close()

	// wait for CTRL-C
	c := make(chan os.Signal, 1)
	signal.Notify(c, os.Interrupt)
	<-c
}
