package main

// os.IsNotExist syntax (es is "error status" here)
// func IsNotExist (es error) bool

// os.Stat() syntax
// func Stat(FileName string) (os.FileInfo, error) . Uses are shown below


//try it with a file that isn't test.txt! 

import (
	"log"
	"os"
)

var (
	examplefile *os.FileInfo
	es error
)

func main() {
	examplefile, es := os.Stat("test.txt")
	if es != nil {
		if os.IsNotExist(es) {
			log.Fatal("No File Found")
		}
	}
	log.Println("Name: ", examplefile.Name())
	log.Println("Size: ", examplefile.Size())
	log.Println("File Permissions: ", examplefile.Mode())
	log.Println("Last modification time: ", examplefile.ModTime())
	log.Println("Is it a directory? :", examplefile.IsDir())
}



