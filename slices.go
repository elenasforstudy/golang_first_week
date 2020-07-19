package main

import "fmt"

var (
	arr                  = make([]int, 10, 12)
	arrInt               = []int{1, 2, 3, 4, 5, 6, 7, 8, 9}
	secondSlice          = []int{1, 2, 3}
	arrString   []string = []string{"John", "Hannah", "Victor", "Ann", "Michael"}
)

// 1 - К каждому элементу []int прибавить 1
func addOne(s []int) {
	for index := range s {
		s[index]++
	}
}

// 2 - Добавить в конец slice число 5
func appendInt(s *[]int, i int) {
	*s = append(*s, i)
}

// 3 - Добавить в начало slice число 5
func prependInt(s []int, i int) []int {
	s = append(s, 0)
	copy(s[1:], s)
	s[0] = i
	return s
}

// 4 - Взять последнее число slice, вернуть его пользователю, а из slice этот элемент удалить
func returnAndDeleteLastItem(s *[]int) int {
	lengthS := len(*s) - 1
	item := (*s)[lengthS]
	*s = append((*s)[:lengthS])
	return item
}

// 5 - Взять первое число slice, вернуть его пользователю, а из slice этот элемент удалить
func returnAndDeleteFirstItem(s *[]int) int {
	item := (*s)[0]
	*s = append((*s)[1:])
	return item
}

// 6 - Взять i-е число slice, вернуть его пользователю, а из slice этот элемент удалить. Число i передает пользователь в функцию
func returnAndDeleteAnyItem(s *[]int, index int) int {
	var itemValue int
	switch {
	case index == 1:
		{
			itemValue = (*s)[0]
			*s = append((*s)[1:])

		}
	case index > 1 && index < len(*s)-1:
		{
			itemValue = (*s)[index-1]
			*s = append((*s)[:index-1], (*s)[index:]...)
		}
	case index == len(*s)-1:
		{
			itemValue = (*s)[len(*s)-1]
			*s = append((*s)[:len(*s)-1])
		}
	}
	return itemValue
}

// 7 - Объединить два slice и вернуть новый со всеми элементами первого и второго
func uniteSlices(s1 []int, s2 []int) []int {
	var unite []int
	unite = append(s1, s2...)
	return unite
}

// 8 - Из первого slice удалить все числа, которые есть во втором
func deleteBySecondSlice(s1 *[]int, s2 *[]int) []int {
	for _, v := range *s2 {
		for j, w := range *s1 {
			if v == w {
				*s1 = append((*s1)[:j], (*s1)[j+1:]...)
			}
		}
	}
	return *s1
}

// 9 - Сдвинуть все элементы slice на 1 влево. Нулевой становится последним, первый - нулевым, последний - предпоследним.
func moveBackOne(s *[]int) []int {
	*s = append((*s)[1:], (*s)[0])
	return *s
}

// 10 - Тоже, но сдвиг на заданное пользователем i
func moveBackAny(s *[]int, shift int) []int {
	if shift < len(*s) {
		*s = append((*s)[shift:], (*s)[:shift]...)
	}
	return *s
}

// 11 - Тоже, что 9, но сдвиг вправо
func moveForwardOne(s *[]int) []int {
	*s = append((*s)[len(*s)-1:], (*s)[:len(*s)-1]...)
	return *s
}

// 12 - Тоже, но сдвиг на i
func moveForwardAny(s *[]int, shift int) []int {
	if shift < len(*s) {
		*s = append((*s)[len(*s)-shift:], (*s)[:len(*s)-shift]...)
	}
	return *s
}

// 13 - Вернуть пользователю копию переданного slice
func returnCopy(s []int) []int {
	newS := make([]int, len(s))
	copy(newS, s)
	return newS
}

// 14 - В slice поменять все четные с ближайшими нечетными индексами. 0 и 1, 2 и 3, 4 и 5...
func swapElements(s *[]int) []int {
	for i := 0; i < len(*s)-1; {
		tmp := (*s)[i]
		(*s)[i] = (*s)[i+1]
		(*s)[i+1] = tmp
		i += 2
	}
	return *s
}

// 15 - Упорядочить slice в порядке: прямом, обратном, лексикографическом.
func mergeInt(left []int, right []int) []int {
	list := make([]int, 0)
	for len(left) > 0 && len(right) > 0 {
		if left[0] < right[0] {
			list = append(list, left[0])
			left = left[1:]
		} else {
			list = append(list, right[0])
			right = right[1:]
		}
	}
	if len(left) > 0 {
		list = append(list, left...)
	}
	if len(right) > 0 {
		list = append(list, right...)
	}

	return list
}

func mergesortInt(s []int) []int {
	length := len(s)
	if length >= 2 {
		middle := length / 2
		s = mergeInt(mergesortInt(s[:middle]), mergesortInt(s[middle:]))
	}
	return s
}

// 15 - Упорядочить slice в порядке: прямом, обратном, лексикографическом.
func orderInt(s []int, by string) []int {
	switch by {
	case "asc":
		{
			s = mergesortInt(s)
		}
	case "desc":
		{
			s = mergesortInt(s)
			for i, j := 0, len(s)-1; i < j; i, j = i+1, j-1 {
				s[i], s[j] = s[j], s[i]
			}
		}
	}

	return s
}

func mergesortLex(s []string) []string {
	length := len(s)
	if length >= 2 {
		middle := length / 2
		s = mergeLex(mergesortLex(s[:middle]), mergesortLex(s[middle:]))
	}
	return s
}

func mergeLex(left []string, right []string) []string {
	list := make([]string, 0)
	for len(left) > 0 && len(right) > 0 {
		if left[0] < right[0] {
			list = append(list, left[0])
			left = left[1:]
		} else {
			list = append(list, right[0])
			right = right[1:]
		}
	}
	if len(left) > 0 {
		list = append(list, left...)
	}
	if len(right) > 0 {
		list = append(list, right...)
	}

	return list
}

func orderLex(s []string) []string {
	s = mergesortLex(s)
	return s
}

func main() {
	fmt.Println("Начальный слайс -", arr)

	// 1
	addOne(arr)
	fmt.Println("1 - К каждому элементу []int прибавить 1 - ", arr)

	// 2
	appendInt(&arrInt, 5)
	fmt.Println("2 - Добавить в конец slice число -", arrInt)

	// 3
	arrInt = prependInt(arrInt, 5)
	fmt.Println("3 - Добавить в начало slice число -", arrInt)

	// 4
	fmt.Println("4 - Удаленный послдений элемент -", returnAndDeleteLastItem(&arrInt), "Измененный слайс -", arrInt)

	// 5
	fmt.Println("5 - Удаленный первый элемент -", returnAndDeleteFirstItem(&arrInt), "Измененный слайс -", arrInt)

	//6
	deletedItem := returnAndDeleteAnyItem(&arrInt, 3)
	fmt.Println("6 - Удален элемент по выбору -", deletedItem, "Измененный слайс -", arrInt)

	//7
	fmt.Println("7 - Объединение срезов:", uniteSlices(arrInt, secondSlice))

	//8
	fmt.Println("8 - Первый срез без элементов второго:", deleteBySecondSlice(&arrInt, &secondSlice))

	//9
	fmt.Println("Начальный слайс -", arrInt)
	fmt.Println("9 - Сдвиг влево на 1", moveBackOne(&arrInt))
	//10
	fmt.Println("Начальный слайс -", arrInt)
	fmt.Println("10 - Сдвиг влево введеное число", moveBackAny(&arrInt, 2))
	//11
	fmt.Println("Начальный слайс -", arrInt)
	fmt.Println("11 - Сдвиг влево на 1", moveForwardOne(&arrInt))
	//12
	fmt.Println("Начальный слайс -", arrInt)
	fmt.Println("12 - Сдвиг вправо на введеное число", moveForwardAny(&arrInt, 2))

	//13
	fmt.Printf("Адрес начального слайса %p\n\n", arrInt)
	fmt.Println("13 - Вернуть копию переданного слайса", returnCopy(arrInt))
	fmt.Printf("Адрес копии %p\n\n", returnCopy(arrInt))

	//14
	fmt.Println("14 - Поменять местами четные и нечетные элементы", swapElements(&arrInt))

	//15
	fmt.Println("15 - По возрастанию:", orderInt(arrInt, "asc"))
	fmt.Println("15 - По убыванию:", orderInt(arrInt, "desc"))

	fmt.Println(arrString)
	fmt.Println("15 - В лексикографическом порядке:", orderLex(arrString))
}
