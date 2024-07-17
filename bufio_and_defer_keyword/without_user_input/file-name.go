package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"os"
)

func CreateFile() {
	//fmt package has more than Println - Printf and Scanf are other examples
	fmt.Printf("I am writing a file\n")
	//usage of Fatalf
	file, er := os.Create("test1.txt")
	if er != nil { 
		log.Fatalf("failed creating file %s", er) 
	}
	//usage of defer - e.g. closing a running file if it completes before main function
	defer file.Close()
	//len - strlen self explanatory 
	len, er := file.WriteString("Reading and writing operations in golang")
	if er != nil {
		log.Fatalf("Failed writing string to %s", er)
	}
		fmt.Printf("\nFile: %s", file.Name())
		fmt.Printf("\nLength: %d bytes", len)
}

func ReadFile() {
	fmt.Printf("\n\nNow Reading a file\n")
	fileName := "test1.txt"
	//ioutil can return filename and contents
	data, er := ioutil.ReadFile("test1.txt")
	if er != nil {
		log.Panicf("Failed reading from file!! %s", er)
	}
	fmt.Printf("\nFile Name %s", fileName)
	fmt.Printf("\nSize is %d", len(data))
	fmt.Printf("\nData read: %s", data)
}

func main() { 
	CreateFile()
	ReadFile() 
}


