package main

import (
	"time"

	"github.com/go-vgo/robotgo"
)

func typing(s string) {
	for i := 0; i < len(s); i++ {
		robotgo.KeyTap(string(s[i]))
		time.Sleep(lag * time.Millisecond)
	}
}

func typeChar(s string, count int) {
	for i := 0; i < count; i++ {
		robotgo.KeyTap(s)
		time.Sleep(lag * time.Millisecond)
	}
}
