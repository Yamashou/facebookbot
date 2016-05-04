package main

import (
	"time"
	"log"
	"./reqCafe"
)

func main() {
	for d := 2; d <= 6; d++ {
		t := time.Date(2016, 5, d, 0, 0, 0, 0, time.Local)
		log.Println(reqCafe.RtCafeInfo(t))
	}
}
