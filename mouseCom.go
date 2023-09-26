package main

import (
	"fmt"
	"math/rand"
	"path"
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
	dir := getFileName("img/")
	var s []string
	for _, n := range dir {
		if path.Ext(n) == ".png" {
			s = append(s, "img/"+n)
		}
	}
	for i := 0; i < count; i++ {
		moveToIcon(s[rand.Intn(len(s))])
		robotgo.Click()
		time.Sleep(2 * lag * time.Millisecond)
		closeWindow()
	}
}

func closeWindow() {
	moveToIcon("img/close/closeImg.png")
	robotgo.Click()
}

func moveRandom(r int, count int) {
	for i := 0; i < count; i++ {
		robotgo.MoveSmooth(rand.Intn(r), rand.Intn(r))
	}
}

func moveToIcon(filename string) {
	img1, _, _ := robotgo.DecodeImg(filename)
	img := robotgo.CaptureImg()
	//img1, _ := robotgo.Read(filename)
	m1, _ := gcv.ImgToMat(img)
	m2, _ := gcv.ImgToMat(img1)
	rs := gcv.FindAllTemplate(m1, m2, 0.8)
	//rs := gcv.FindAllImg(img1, img)
	//fmt.Println("find: ", rs, "HIGH: ", rs[0].MaxVal)
	if len(rs) != 0 {
		robotgo.MoveSmooth(rs[0].Middle.X/2, rs[0].Middle.Y/2)
	}
}

func scrollTop() {
	robotgo.Scroll(1000, 1000)
}
