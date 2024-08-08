package main

import (
	"fmt"
	"os"

	"github.com/gocarina/gocsv"
)

// this setup is not not very unusual when working with gorm (maybe also other ORMs)
type Wrapper struct {
	Name string `csv:"name"`
	Age  uint   `csv:"age"`
	// Inner creates a problem, both as pointer and as value
	Inner *Inner // add `csv:"-"` to avoid rapid increase in consumed memory and your program being killed
}

type Inner struct {
	Position string
	Salary   uint
	Wrappers *Wrapper
}

func main() {
	file, err := os.OpenFile("./data.csv", os.O_RDWR|os.O_CREATE, os.ModePerm)
	if err != nil {
		panic(err)
	}
	defer file.Close()

	var wraps []Wrapper
	if err := gocsv.UnmarshalFile(file, &wraps); err != nil {
		panic(err)
	}
	fmt.Printf("%+v\n", wraps)
}
