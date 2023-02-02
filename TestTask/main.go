package main

import (
	"TestTask/operation"
	"log"
	"os"
)

func main() {
	args := os.Args
	if len(args) == 1 {
		log.Fatal("Отсутствуют аргумены")
	}
	operation.OpenFile(args[1])
}
