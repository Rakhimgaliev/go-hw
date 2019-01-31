package main

import (
	"fmt"
	"io/ioutil"
	"log"
)

func main() {
	files, err := ioutil.ReadDir("testdata")
	if err != nil {
		log.Fatal(err)
	}

	for _, file := range files {
		fmt.Printf("%T\n", file)
		fmt.Println(file.Name())
	}
}