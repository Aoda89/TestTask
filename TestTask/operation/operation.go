package operation

import (
	"TestTask/initializing"
	"fmt"
	"log"
	"os"
)

// OpenFile Открытие файла
func OpenFile(name string) {
	file, err := os.ReadFile(name)
	if err != nil {
		log.Fatal("Ошибка чтения файла")
	}
	array := initializing.InitializingArray(file)
	PrintArray(array)
}

// PrintArray Печать результата
func PrintArray(array [][]string) {
	for _, row := range array {
		fmt.Print("\n")
		for i, val := range row {
			if val == "," {
				fmt.Print(",")
				continue
			}
			if i != len(row)-1 {
				fmt.Print(val, ",")
			} else {
				fmt.Print(val)
			}
		}
	}
}
