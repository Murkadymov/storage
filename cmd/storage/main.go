package main

import (
	"fmt"
	"modstorage/internal/storage"
)

func main() {

	fmt.Println("Process started")
	st := storage.NewStorage()
	fileGot, err := st.Upload("TestingFile", []byte("TextInsideFile"))
	fmt.Println(fileGot, string(fileGot.Data), err)
}
