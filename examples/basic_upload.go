package main

import (
	"fmt"
	"github.com/hduplooy/sftp"
	"log"
	"time"
)

func main() {
	config := sftp.Config{
		Username: "root",
		Password: "password", // required only if password authentication is to be used
		Server:   "localhost:22",
		Timeout:  time.Second * 30, // 0 for not timeout
	}
	client, err := sftp.New(config)
	if err != nil {
		log.Fatalln(err)
	}
	defer client.Close()

	uploaded, err := client.UploadIfNewer("/home/user/testfile", "testfile")
	if err != nil {
		log.Fatalln(err)
	}
	fmt.Printf("Uploaded file: %v\n", uploaded)
}
