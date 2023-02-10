package main

import (
	"log"
	"runtime"
)

func loop() {
	for {
	}
}

func main() {
	cpus := runtime.NumCPU()
	//runtime.GOMAXPROCS(cpus)
	log.Println("[INFO]", "CPU Core Count:", cpus)
	for i := 0; i < cpus - 1; i++ {
		go loop()
	}
	loop()
}
