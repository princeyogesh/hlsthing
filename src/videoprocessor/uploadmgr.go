package main

import (
	"fmt"
	"strings"
	"sync"
)

/*
objectProvider is just one of cloud object storage provider like AWS s3, azure blob, google obkect

chunks will always be stored in folder with name  sourcefile_
*/

func uploadChunkstoCloud(sourcefile string, objectProvider string) {
	sourceFolder := sourcefile + "_"
	var objectProviderCount = len(objectProvider) // change later when actually implemented
	var wg sync.WaitGroup
	wg.Add(objectProviderCount)
	if strings.Compare(objectProvider, "S3") == 0 {
		go AWSS3Handler(sourceFolder, &wg)
	}
	if strings.Compare(objectProvider, "AzureBlob") == 0 {
		go AzureBlobHandler(sourceFolder, &wg)
	}
	if strings.Compare(objectProvider, "GoogleStorage") == 0 {
		go GCPStorageHandler(sourceFolder, &wg)
	}

	wg.Wait()
	fmt.Println("uploaded to all requested objectstorages ", objectProvider)
}

func GCPStorageHandler(sourceFolder string, wg *sync.WaitGroup) {
	defer wg.Done()
	panic("unimplemented")
}

func AzureBlobHandler(sourceFolder string, wg *sync.WaitGroup) {
	defer wg.Done()
	panic("unimplemented")
}

func AWSS3Handler(sourceFolder string, wg *sync.WaitGroup) {
	defer wg.Done()
	panic("unimplemented")
}
