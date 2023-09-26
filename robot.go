package main

import (
	"time"
	//"github.com/vcaesar/bitmap"
)

var lag time.Duration = 500

func main() {
	moveRandom(1000, 3)
	moveRandomIcon(3)
	hellofunc()
	//closeWindow()

}
