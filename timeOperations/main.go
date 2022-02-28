package main

import (
	"fmt"
	"time"
)

func main() {
	// Konsol : Saat & Tarih İşlemleri
	// Date() metodu ile kendi tarih verimizi oluşturuyoruz.
	t := time.Date(2022, time.November, 10, 20, 0, 0, 0, time.UTC)
	// t adıyla tarih tipinde oluşturulan veriyi string tipinde ekrana yazıyoruz.
	fmt.Printf("Go Launch at %s\n", t)
	fmt.Println("-----")
	// time.Now() ile şu anın zaman bilgisini alıyoruz.
	now := time.Now()
	// Elde edilen zaman bilgisini ekrana yazıyoruz.
	fmt.Printf("The time now is %s\n", now)
	fmt.Println("-----")
	// flk başta oluşturduğumuz t adındaki zaman bilgisinden Ay, Gün ve Haftanın Günü bilgilerini ekrana yazıyoruz.
	fmt.Println("The month is", t.Month())
	fmt.Println("The day is", t.Day())
	fmt.Println("The weekday is", t.Weekday())
	fmt.Println("-----")
	// Tarihe 1 gün ekle!
	tomorrow := t.AddDate(0, 0, 1)
	fmt.Printf("Tomorrow is %v, %v, %v, %v\n",
		tomorrow.Weekday(),
		tomorrow.Month(),
		tomorrow.Day(),
		tomorrow.Year())
	fmt.Println("-----")
	longFormat := "Monday, Junuary 2, 2006"
	fmt.Println("Tomorrow is", tomorrow.Format(longFormat))
	fmt.Println("-----")
	shortFormat := "1/2/06"
	fmt.Println("Tomorrow is", tomorrow.Format(shortFormat))
}
