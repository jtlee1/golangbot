package main

import (
	"fmt"
	"os"
	"os/exec"
	"time"

	"gocv.io/x/gocv"
)

func hellofunc() {
	exec.Command("open", "https://google.com").Start()
	time.Sleep(6 * lag * time.Millisecond)
	typing("halo")
	typeChar("backspace", 3)
	typing("ello justin")
}

func display(img gocv.Mat) {
	window := gocv.NewWindow("Hello")
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

// cite https://www.tutorialspoint.com/golang-program-to-get-the-list-the-name-of-files-in-a-specified-directory
// get file name under directory
func getFileName(path string) []string {
	var ret []string
	dir, err := os.Open(path)
	if err != nil {
		fmt.Println("Error:", err)
		os.Exit(3)
	}
	defer dir.Close()
	files, err := dir.Readdir(-1)
	if err != nil {
		fmt.Println("Error:", err)
		os.Exit(3)
	}
	for _, file := range files {
		ret = append(ret, file.Name())
	}
	return ret
}
