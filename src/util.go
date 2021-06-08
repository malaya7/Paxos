package main

import (
	"fmt"
	"log"
)

func Logs(s string) {
	fmt.Println("======")
	log.Println("Node-"+PORT, s)
}
