package main

// using os.Create . Syntax is as follows 
// func Create(filename string) (*File, error) 

import (
	"log"
	"os"
)

func main() {
	// es will be "error status"
	examplefile, es := os.Create("example.txt")
	if es != nil {
		log.Fatal(es)
	}
	log.Println(examplefile)
	examplefile.Close()
}
