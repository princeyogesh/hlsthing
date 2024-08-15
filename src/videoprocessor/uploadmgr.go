package main

import "strings"

/*
objectProvider is just one of cloud object storage provider like AWS s3, azure blob, google obkect

chunks will always be stored in folder with name  sourcefile_
*/

func uploadChunkstoCloud(sourcefile string, objectProvider string) {
	sourceFolder := sourcefile + "_"

	if strings.Compare(objectProvider, "S3") == 0 {
		AWSS3Handler(sourceFolder)
	}
	if strings.Compare(objectProvider, "AzureBlob") == 0 {
		AzureBlobHandler(sourceFolder)
	}
	if strings.Compare(objectProvider, "GoogleStorage") == 0 {
		GCPStorageHandler(sourceFolder)
	}
}

func GCPStorageHandler(sourceFolder string) {
	panic("unimplemented")
}

func AzureBlobHandler(sourceFolder string) {
	panic("unimplemented")
}

func AWSS3Handler(sourceFolder string) {
	panic("unimplemented")
}
