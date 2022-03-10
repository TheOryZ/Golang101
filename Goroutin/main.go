package main

import (
	"runtime"
	"time"
)

func main() {
	// basit gorouting kullanımı
	go xFunc()
	time.Sleep(100 * time.Millisecond)

	runtime.GOMAXPROCS(1)
	go zFunc()
	time.Sleep(100 * time.Millisecond)
}
func xFunc() {
	for i := byte('a'); i <= byte('z'); i++ {
		println(string(i))
	}
}

func zFunc() {
	for i := byte('a'); i <= byte('z'); i++ {
		go println(string(i))
	}
}
