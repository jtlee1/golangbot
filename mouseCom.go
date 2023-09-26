package main

import (
	"fmt"
	"math/rand"
	"time"

	"github.com/go-vgo/robotgo"
	"github.com/vcaesar/gcv"
)

func pos() {
	for {
		x, y := robotgo.GetMousePos()
		fmt.Println("Current position X=", x, "Y=", y)
		if x == 540 && y == 960 {
			fmt.Println("Cursor @center")
		}
		time.Sleep(lag * time.Millisecond)
		//robotgo.Move(x, y)
	}
}

func moveRandomIcon(count int) {
	
	for i := 0; i < count; i++ {
		robotgo.MoveSmooth(rand.Intn(r), rand.Intn(r))
	}
}

func moveRandom(r int, count int) {
	for i := 0; i < count; i++ {
		robotgo.MoveSmooth(rand.Intn(r), rand.Intn(r))
	}
}

func moveToIcon(filename string) {
	img1, _, _ := robotgo.DecodeImg(filename)
	img := robotgo.CaptureImg()
	m1, _ := gcv.ImgToMat(img)
	m2, _ := gcv.ImgToMat(img1)
	rs := gcv.FindAllTemplate(m1, m2, 0.8)
	//fmt.Println("find: ", rs)
	robotgo.MoveSmooth(rs[0].Middle.X/2, rs[0].Middle.Y/2)
}

func scrollTop() {
	robotgo.Scroll(1000, 1000)
}
