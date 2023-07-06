package main

import (
	"github.com/go-well/tts"
	"log"
	"os"
)

func main() {
	file, err := os.Create("out.mp3")
	if err != nil {
		log.Fatal(err)
	}
	err = tts.Generate("一号设备报警", file)
	if err != nil {
		log.Fatal(err)
	}
}
