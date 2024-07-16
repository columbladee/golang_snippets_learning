package main

import (
	"log"
	"os"
)

func main() {
	//same thing as single, but os.MkdirAll
	if er := os.MkdirAll("parent/child/grandchild", os.ModePerm); er != nil {
		log.Fatal(er)
	}
}


