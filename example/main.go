package main

import (
	"math/rand"
	"time"

	"github.com/valbeat/whatday"
)

func main() {
	rand.Seed(time.Now().UnixNano())
	whatday.Print()
}
