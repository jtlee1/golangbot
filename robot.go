package main

import (
	"fmt"
	"time"

	"gocv.io/x/gocv"
	//"github.com/vcaesar/bitmap"
)

var lag time.Duration = 500

// cite code https://github.com/hybridgroup/gocv/blob/release/cmd/showimage/main.go
func display(img gocv.Mat) {
	//filename := os.Args[1]
	window := gocv.NewWindow("Hello")
	//img := gocv.IMRead(i, gocv.IMReadColor)
	if img.Empty() {
		fmt.Printf("Error reading image from:")
		return
	}
	for {
		window.IMShow(img)
		if window.WaitKey(1) >= 0 {
			break
		}
	}
}

func main() {
	moveToIcon()

}
