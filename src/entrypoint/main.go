package main

import (
	"flag"
	"fmt"
)

func main() {
	inputFile := flag.String("filename", "../../assets/BigBunny.mp4", "Please enter file name which you would like to proceed")

	fmt.Println("Processing the CLI input ", *inputFile)
	fmt.Println("hello worlds")

	//Take input with rest API
}
