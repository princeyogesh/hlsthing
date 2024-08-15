package main

import (
	"flag"
	"fmt"
)

var storeToRemote *bool
var CloudMode bool = false
var CloudObjectProvider string = ""

func main() {
	inputFile := flag.String("filename", "/home/app/assets/BigBunny.mp4", "Please enter file name which you would like to proceed")
	storeToRemote = flag.Bool("storeRemote", false, "used if we need to store .ts file remotely on some remote machine ")
	fmt.Println("Processing the CLI input ", *inputFile)
	fmt.Println("hello worlds")

	//
	breakdwonFile(*inputFile)
	//
	if CloudMode {
		uploadChunkstoCloud(*inputFile, CloudObjectProvider)
		docleanup()

	} else if *storeToRemote {
		docleanup()
		panic("not able to send on any machine , not implemented ")
	} else {
		fmt.Println("no need to upload anywhere, please do manual cleanup")
	}

}

func docleanup() {
	fmt.Println("doing cleanup")
}
