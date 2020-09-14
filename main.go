package main

import "fmt"
import "reflect"

func main() {
	name := "Alexandre"
	version := 1.1

	fmt.Println("hello ", name)
	fmt.Println("versão: ", version)

	fmt.Println("1- Start monitoring")
	fmt.Println("2- Show logs")
	fmt.Println("0- Exit")

	var cmd int
	fmt.Scan
}
