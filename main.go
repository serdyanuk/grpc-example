package main

import (
	"math/rand"
	"time"
)

func main() {
	rand.Seed(time.Now().UnixNano())

}

func randRange(min, max int) int {
	return rand.Intn(max-min) + min
}
