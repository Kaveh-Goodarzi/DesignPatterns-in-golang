package main

import (
	"fmt"
	"sync"
)

// Singleton:

var mutex = &sync.Mutex{}

type Config struct {
	// Config Variables
}

var counter int = 1
var instance *Config

func GetConfigInstance() *Config {
	if instance == nil {
		mutex.Lock()
		defer mutex.Unlock()
		if instance == nil {
			fmt.Println("Creating new Config instance")
			instance = &Config{}
			counter++
		} else {
			fmt.Println("Config instance already created")
		}
	} else {
		fmt.Println("Config instance already created")
	}
	return instance
}

func main() {
	for i := 0; i < 100; i++ {
		go GetConfigInstance()
	}
	fmt.Scanln()
}	