package main

import (
	"fmt"
	"strings"
)

var (
	text = "Some say the world will end in fire " +
		"Some say in ice "

	arrInt = []int{1, 10, 22, 3, 1, 5}

	arrIntTwo = []int{1, 2, 3, 4, 5, 6, 7, 8, 9}

	fibElement = 10
)

// 1 - Есть текст, надо посчитать сколько раз каждое слова встречается.
func countWordsEntrance(str string) map[string]int {
	wordsMap := make(map[string]int)
	text := strings.Fields(str)
	for _, val := range text {
		wordsMap[val]++
	}
	return wordsMap
}

// 2 - Есть очень большой массив или slice целых чисел, надо сказать какие числа в нем упоминаются хоть по разу.
func countNumbEntrance(s []int) map[int]int {
	numbMap := make(map[int]int)
	for index := range s {
		numbMap[s[index]]++
	}
	return numbMap
}

// 3 - Есть два больших массива чисел, надо найти, какие числа упоминаются в обоих
func countNumbIntersection(s1, s2 []int) map[int]int {

	firstMap := make(map[int]int)
	resultMap := make(map[int]int)

	for index := range s1 {
		firstMap[s1[index]]++
	}
	for index := range s2 {
		_, ok := firstMap[s2[index]]
		if ok {
			resultMap[s2[index]]++
		}
	}

	return resultMap
}

// 4 - Сделать Фибоначчи с мемоизацией
func _fibonacci(index int, m map[int]int) int {
	if index == 0 || index == 1 {
		return index
	}
	if m[index] == 0 {
		m[index] = _fibonacci(index-1, m) + _fibonacci(index-2, m)
	}
	return m[index]
}

func fibonacci(index int) map[int]int {
	fibMap := map[int]int{
		1: 1,
	}

	_fibonacci(index, fibMap)
	return fibMap
}

func main() {
	fmt.Println("\n\n1", "Текст:", text)
	fmt.Println("Надо посчитать, сколько раз каждое слова встречается\n", countWordsEntrance(text))

	fmt.Println("\n\n2", "Слайс/массив чисел", arrInt)
	fmt.Println("Какие числа появляются в слайсе (массиве) хоть по разу\n", countNumbEntrance(arrInt))

	fmt.Println("\n\n3 - Какие числа упоминаются в обоих слайсах - ", countNumbIntersection(arrInt, arrIntTwo))

	fmt.Println("\n\n4 -", "Элемент:", fibElement)
	fmt.Println("Фибоначчи с мемоизацией", fibonacci(fibElement))
}
