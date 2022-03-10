package main

import "time"

func main() {
	// basit gorouting kullanımı
	go xFunc()
	time.Sleep(100 * time.Millisecond)
	//...
}
func xFunc() {
	for i := byte('a'); i <= byte('z'); i++ {
		println(string(i))
	}
}
