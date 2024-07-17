package main

import (
	"bufio"
	"fmt"
	"io/ioutil"
	"log"
	"os"
)

func CreateFile(filename, text string) {
	fmt.Printf("This line goes into the file\n")

	//create
	//user will specify filename
	file, er := os.Create(filename)
	if er != nil {
		log.Fatalf("Couldn't create file %s", er)
	}
	//defer as before
	defer file.Close()

	//Writing string to file
	len, er := file.WriteString(text)
	if er != nil {
		log.Fatalf("Failed writing to file %s", er)
	}
	fmt.Printf("\nFilename is %s", file.Name())
	fmt.Printf("\nFile size is %d", len)
}

func ReadFile(filename string) {
	//reading as before
	fmt.Printf("\nReading file function...\n")
	data, er := ioutil.ReadFile(filename) // again, demonstrates ioutil's ReadFile
	if er != nil {
		log.Panicf("failed reading file: %s", er)
	}
	fmt.Printf("\nFilename %s", filename)
	fmt.Printf("\nFile size %d", len(data))
	fmt.Printf("\nContents are: %s", data)
}

func main() {
	//scanln hereuses reference (denoted by &) to grab filename
	fmt.Println("Enter filename: ") 
	var filename string
	fmt.Scanln(&filename) //looks a lot like C, huh?
	fmt.Println("Write some text to put in the file: ")
	//define a NewReader ""Object"", then user inputReader.ReadString to grab input
	inputReader := bufio.NewReader(os.Stdin) //stdin makes sense...
	input, _ := inputReader.ReadString('\n')
	CreateFile(filename, input)
	ReadFile(filename)
}
