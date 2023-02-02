package initializing

import (
	"TestTask/checktable"
	"bytes"
	"log"
)

// CountRows Подсчет строк и проверка формата
func CountRows(file []byte) int {
	size := len(file)
	if size == 0 {
		log.Fatal("Пустой файл ")
	}
	sep := []byte{'\n'}
	row := bytes.Count(file, sep)
	if file[size-1] != '\n' {
		row++
	}
	return row
}

// CountColumn Подсчет столбцов и проверка формата (44-запятая 10-\n 13-возрат коретки, 32-пробел)
func CountColumns(file []byte) int {
	sepCountHendler := 1
	sepCount := 1
	i := 0
	for ; file[i] != 10; i++ {
		if file[i] == 44 {
			sepCountHendler++
		}
		if file[i] == 44 && file[i+1] == 44 || file[i] == 44 && file[i+1] == 13 {
			log.Fatal("Ошибка заголовка ")
		}
		if file[i] == 32 || file[i] == 10 && file[i+2] == 10 {
			return sepCountHendler
		}
	}
	i++
	for ; i < len(file); i++ {
		if file[i] == 32 || file[i] == 10 && file[i+2] == 10 {
			log.Fatal("Лишнии эскейп последовательности ")
		}
		if file[i] == 44 && file[i+1] == 44 || file[i] == 44 && file[i+1] == 13 {
			log.Fatal("Колличество столбцов неверно ")
		}
		if file[i] == 44 {
			sepCount++
		}
		if file[i] == 10 && sepCount == sepCountHendler || i == len(file)-1 && sepCount == sepCountHendler {
			sepCount = 1
		} else if file[i] == 10 && sepCount != sepCountHendler || i == len(file)-1 && sepCount != sepCountHendler || file[len(file)-1] == 44 {
			log.Fatal("Колличество столбцов неверно ")
		}
	}
	return sepCountHendler
}

// InitializingArray Инициализация 2d массива (44-запятая 10-\n 13-возрат коретки, 32-пробел)
func InitializingArray(file []byte) [][]string {

	row := CountRows(file)
	if file[0] != 44 || row < 2 {
		log.Fatal("Ошибка формата")
	}
	column := CountColumns(file)
	array := make([][]string, row)
	for i := range array {
		array[i] = make([]string, column)
	}
	i := 0
	j := 1
	for count := 1; count < len(file); count++ {
		if file[count] == 44 {
			j++
		} else if file[count] == 10 {
			i++
			j = 0
		} else if file[count] == 13 {
			continue
		} else {
			array[i][j] += string(file[count])
		}
	}
	checktable.CheckStruct(array)

	return array
}
