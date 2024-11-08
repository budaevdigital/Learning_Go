// Дан массив целых чисел A и целое число k.
// Нужно найти и вывести индексы пары чисел, сумма которых равна k.
//
// Если таких чисел нет, то вернуть пустой слайс.
// Индексы можно вернуть в любом порядке

package main

import "fmt"

func find(arr []int, k int) []int {
	m := make(map[int]int)
	// будем складывать в неё индексы массива, а в качестве ключей использовать само значение
	for i, a := range arr {
		if j, ok := m[k-a]; ok { // если значение k-a уже есть в массиве, значит, arr[j] + arr[i] = k и мы нашли, то что нужно
			return []int{i, j}
		}
		// если искомого значения нет, то добавляем текущий индекс и значение в map
		m[a] = i
	}
	// не нашли пары подходящих чисел
	return nil
}

func main() {
	var arrayCount [100]int
	valueK := 23
	indexCount := 0

	// просто заполняем массив
	for index := range arrayCount {
		arrayCount[index] = indexCount
		// fmt.Printf("Index: %v Value: %v \n", index, indexCount)
		indexCount++
	}

	// берем слайс от массива
	sliceArray := arrayCount[:]
	fmt.Println(sliceArray)

	newValue := find(sliceArray, valueK)

	fmt.Println(newValue)

}
