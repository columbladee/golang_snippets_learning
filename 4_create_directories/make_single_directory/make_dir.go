package main

import (
	"log"
	"os"
)
func main() {
	// we will be using the os.Mkdir os.Modeperm and os.FileMode*
		//func Mkdir(name string, perm os.FileMode) 
		// type FileMode uint32
		// const ModePerm FileMode = 0777
	//Modeperm defaults to 0777 . 



	if er:= os.Mkdir("newdirectory", os.ModePerm); er != nil {
		log.Fatal(er)
	}
}
