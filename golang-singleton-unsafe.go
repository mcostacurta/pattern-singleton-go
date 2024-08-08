package main

// Singleton pattern - thread NOT safe
// command -race point 2 - 4 data races - it can cause a deadlock

import (
	"fmt"
	"math/rand"
	"time"
)

type UniqueMessage struct {
	msg string
}

var instance *UniqueMessage

func Unique() *UniqueMessage {

	// Traditional way to create a singleton
	if instance == nil {
		randomNum := rand.Intn(100)
		instance = &UniqueMessage{msg: fmt.Sprintf("That is the unique message {%d}", randomNum)}
	}

	return instance
}

func main() {

	// This will be called only once
	go func() {
		time.Sleep(time.Millisecond * 600)
		fmt.Println(*Unique())
	}()

	// Create 100 goroutines
	for i := 0; i < 100; i++ {
		go func(j int) {
			time.Sleep(time.Millisecond * 60)
			fmt.Println(*Unique(), " - ", j)
		}(i)
	}

	fmt.Scanln()
}
