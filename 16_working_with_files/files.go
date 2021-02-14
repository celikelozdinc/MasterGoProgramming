package main

import (
	"log"
	"os"
	"time"
)

func main() {
	fd, err := os.Create("dummy.log")
	defer fd.Close()
	if err != nil {
		log.Fatal(err)
	} else {
		log.Printf("Type of file descriptor : %T\n", fd)
	}

	var fileInfo os.FileInfo
	fileInfo, err = os.Stat("dummy.log")
	if err != nil {
		log.Println("File does not exist")
	} else {
		log.Println("File exists, here is the info:")
		log.Printf("Modification time : %v", fileInfo.ModTime())
		log.Printf("Underlying data structure : %v", fileInfo.Sys())
	}

	time.Sleep(5 * time.Second)

	oldPath, newPath := "dummy.log", "dummy_updated.log"
	err = os.Rename(oldPath, newPath)
	if err != nil {
		log.Fatal(err)
	} else {
		log.Println("Renamed successfully")
	}

}
