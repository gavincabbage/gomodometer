package main

import (
	"fmt"
	"os"
)

func main() {
	f, err := os.Open("/dev/input/mice")
	if err != nil {
		panic(err)
	}

	readBuffer := make([]byte, 3)
	for {
		num, err := f.Read(readBuffer)
		if err != nil {
			panic(err)
		}
		fmt.Println("read", num, "bytes: ", readBuffer[0], readBuffer[1], readBuffer[2])
	}

}
