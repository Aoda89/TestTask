package checktable

import (
	calculation "TestTask/calculations"
	"log"
	"regexp"
	"strconv"
	"strings"
	"unicode"
)

// CheckStruct Проверка таблицы на наличие недопустимых символов
func CheckStruct(array [][]string) {
	matched := false
	for _, row := range array {
		for _, val := range row {
			matched, _ = regexp.MatchString(`[^0-9A-Za-z\s=,+*\-/]`, val)
			if matched == true {
				log.Fatal("Недопустимые символы")
			}
			CheckingUniqueElementsNumbers(array)
			CheckingUniqueElementsHandler(array)
			CheckHeading(array)
			CheckFirstColumn(array)
			Search(array)
		}
	}
}

// CheckHeading Проверка заголовка на наличие недопустимых символов
func CheckHeading(array [][]string) {
	matched := false
	for i := 0; i < len(array[0]); i++ {
		matched, _ = regexp.MatchString(`[^A-Za-z\s]`, array[0][i])
		if array[0][i] == "," {
			matched = false
		}
		if matched == true {
			log.Fatal("Недопустимое название заголовка")
		}
	}
}

// CheckFirstColumn Проверка столбца с номерами строк  на наличие недопустимых символов
func CheckFirstColumn(array [][]string) {
	matched := false
	for i := 0; i < len(array); i++ {
		matched, _ = regexp.MatchString(`[^0-9\s]`, array[i][0])
		if array[0][i] == "," {
			matched = false
		}
		if matched == true {
			log.Fatal("Недопустимый номер строки")
		}
	}
}

// Search Поиск формулы
func Search(array [][]string) {
	matchedCorrect := false
	matchedIncorrect := false
	formula := new(string)
	linkArray := new([][]string)
	linkArray = &array
	row := len(array)
	column := len(array[0]) - 1
	for i := 1; i < row; i++ {
		for j := 1; j < column; j++ {
			matchedCorrect, _ = regexp.MatchString(`^=.*\d{1}.*`, array[i][j])
			matchedIncorrect, _ = regexp.MatchString(`(?:.=)(?:).*|(?:^[A-Za-z]).*|(?:^[^=]\D).*|(?:^=[A-Za-z]+$).*`, array[i][j])
			if matchedCorrect == true {
				formula = &array[i][j]
				SearchColumn(formula, linkArray)
			}
			if matchedIncorrect == true {
				log.Fatal("Ошибочные данные в ячейке")
			}
		}
	}
}

// SearchColumn Подстановка значения в формулу
func SearchColumn(formula *string, array *[][]string) string {

	mas := *formula
	var handler string
	var number string
	var storage string
	var recursVolum string
	result := new(string)

	for i := 1; i < len(mas); i++ {
		if unicode.IsLetter(rune(mas[i])) {
			handler += string(mas[i])
		} else if unicode.IsNumber(rune(mas[i])) {
			number += string(mas[i])
		}
		if mas[i] == '+' || mas[i] == '/' || mas[i] == '*' || i == len(mas)-1 {
			if handler == "" {
				storage += number
				if i != len(mas)-1 {
					storage += string(mas[i])
				}
				number = ""
			} else {
				result = SearchVolume(handler, number, array)
				_, err := strconv.Atoi(*result)
				if err != nil {
					recursVolum = SearchColumn(result, array)
					storage += recursVolum

				} else {
					storage += *result
					if i != len(mas)-1 {
						storage += string(mas[i])
					}
					handler = ""
					number = ""
				}
			}
		}
	}
	storageNew := storageSlice(storage)
	answer := calculation.Calculate(storageNew)
	*formula = answer
	return answer
}

// Разделение строки с выражением на массив символов
func storageSlice(storage string) []rune {
	storageNew := make([]rune, len(storage))
	for i := 0; i < len(storage); i++ {
		storageNew[i] = rune(storage[i])
	}
	return storageNew
}

// SearchVolume Поиск значений
func SearchVolume(handler, number string, array *[][]string) *string {
	var receivedPosition string
	mas := *array
	row := len(mas)
	column := len(mas[0])
	countRow := 0
	countColumn := 0
	countMatch := 0
	for countRow = 0; countRow < row; countRow++ {
		if mas[countRow][0] == number {
			countMatch++
			break
		}
	}
	for countColumn = 0; countColumn < column; countColumn++ {
		if mas[0][countColumn] == handler {
			countMatch++
			break
		}
	}
	if countMatch != 2 {
		log.Fatal("Неверное имя ячейки ")
	}
	resalt := mas[countRow][countColumn]
	receivedPosition += handler
	receivedPosition += number
	check := strings.FieldsFunc(resalt, Split)
	for i := 0; i < len(check); i++ {
		if check[i] == receivedPosition {
			log.Fatal("Ячейка ссылается сама на себя")
		}
	}
	return &resalt
}

// Split Функия для проверки ссылается ли ячейка сама на себя
func Split(r rune) bool {
	return r == '+' || r == '-' || r == '/' || r == '*' || r == '='
}

// Проверка которая определяет являеются ли все названия столбцов уникальными
func CheckingUniqueElementsNumbers(array [][]string) {
	row := len(array)
	checkArray := make([]string, row)
	unique := map[string]bool{}
	for i := 0; i < row; i++ {
		checkArray[i] = array[i][0]
	}
	for _, w := range checkArray {
		unique[w] = true
	}
	if len(checkArray) != len(unique) {
		log.Fatal("В первом столбце имеются одинаковые значения")
	}
}

// Проверка которая определяет являеются ли все названия строк уникальными
func CheckingUniqueElementsHandler(array [][]string) {
	column := len(array[0]) - 1
	checkArray := make([]string, column)
	unique := map[string]bool{}
	for i := 0; i < column; i++ {
		checkArray[i] = array[0][i]
	}
	for _, w := range checkArray {
		unique[w] = true
	}
	if len(checkArray) != len(unique) {
		log.Fatal("Столбцы имеют одинаковое имя")
	}
}
