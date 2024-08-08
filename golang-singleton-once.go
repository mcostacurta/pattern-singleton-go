package main

// Singleton pattern - thread safe using sync.Once

import (
	"fmt"
	"math/rand"
	"sync"
	"time"
)

var once sync.Once

type UniqueMessage struct {
	msg string
}

var instance *UniqueMessage

func Unique() *UniqueMessage {

	// This will be called only once
	once.Do(func() {
		randomNum := rand.Intn(100)
		instance = &UniqueMessage{msg: fmt.Sprintf("That is the unique message {%d}", randomNum)}
	})

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
