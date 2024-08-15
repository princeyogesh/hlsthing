package main

import (
	"fmt"
	"os"
	"os/exec"
)

func breakdwonFile(input string) {

	fmt.Println("******Processing ", input, " **********")

	runFFMpegcommand(input)
}

// var outputPath string = "/home/app/assets/segment%03d.ts"
var outputPath string = "segment%03d.ts"
var hlspath string = "/home/app/assets/bunny.m3u8"

func runFFMpegcommand(inputfile string) {
	//x := "\""+outputPath+"/segment%03d.ts"+"\""
	//cmd := `ffmpeg -i ${videoPath} -codec:v libx264 -codec:a aac -hls_time 10 -hls_playlist_type vod -hls_segment_filename "${outputPath}/segment%03d.ts" -start_number 0 ${hlsPath}`;
	output_directory := inputfile + "_"
	createOutput_directory(output_directory)
	cmd := exec.Command("ffmpeg", "-i", inputfile, "-codec:v", "libx264", "-codec:a", "aac", "-hls_time", "10", "-hls_playlist_type", "vod",
		"-hls_segment_filename", output_directory+"/"+outputPath, "-start_number", "0", hlspath)

	//fmt.Println(inputfile)
	//cmd := exec.Command("ls", "-lrt", outputPath+"/")
	fmt.Println(cmd)
	_, err := cmd.Output()

	if err != nil {
		fmt.Println("ERROR occured ")
		fmt.Println(string(err.Error()))

	}

}

func createOutput_directory(foldername string) {
	fmt.Println("**************", foldername)
	_, err := os.Stat(foldername)
	if os.IsNotExist(err) {
		newDirErr := os.Mkdir(foldername, 0777)
		if newDirErr != nil {
			panic("Failed to create new directory")
		}
	}
}
