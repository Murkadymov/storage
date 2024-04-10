package main

import (
	"fmt"
	"modstorage/internal/storage"
)

func main() {

	fmt.Println("Process started")
	st := storage.NewStorage()

	fmt.Println(st)
}
