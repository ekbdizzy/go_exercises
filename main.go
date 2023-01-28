package main

import (
	"fmt"
	"strings"
)

const PassStatus = "pass"
const FailStatus = "fail"

type HealthCheck struct {
	ServiceId int
	status    string
}

func GenerateCheck() (checks []HealthCheck) {
	for i := 0; i < 5; i++ {
		if i%2 == 0 {
			checks = append(checks, HealthCheck{ServiceId: i, status: PassStatus})
		} else {
			checks = append(checks, HealthCheck{ServiceId: i, status: FailStatus})
		}
	}
	return
}

func showGenerateCheck() {
	fmt.Println("Here would be shown ID")
	checks := GenerateCheck()
	for _, check := range checks {
		if check.status == "pass" {
			fmt.Println(check.ServiceId)
		}
	}
}

func hasSameValues(nums []int) bool {
	// Дан целочисленный массив nums, выведите на экран true,
	// если любое значение встречается в массиве как минимум дважды,
	// и false, если каждый элемент различен.
	quantity := map[int]int{}
	for _, value := range nums {
		quantity[value]++
		if quantity[value] > 1 {
			return true
		}
	}
	return false
}

func isSorted(words []string) bool {
	// Вам дан слайс состоящий из строк.
	// Необходимо напечатать true тогда и только тогда,
	// когда все слова в слайсе отсортированы лексикографически
	// относительно друг друга.
	var length int = len(words)
	for i, value := range words {
		if i != length-1 {
			if value > words[i+1] {
				return false
			}
		}
	}
	return true
}

func countLetters(str string) {
	// Нам дана строка следующего вида
	// “съешь ещё этих мягких французских булок, да выпей чаю”.
	// Используя тип данных map посчитайте количество повторений символов в этой строке.
	// В результате выведите на экран список символ - количество повторений.
	str = strings.ReplaceAll(str, " ", "")
	lettersCounter := map[string]int{}
	for _, value := range str {
		lettersCounter[string(value)]++
	}
	for key, value := range lettersCounter {
		fmt.Printf("%v - %v\n", key, value)
	}
}

func main() {
	//showGenerateCheck()
	str := "съешь ещё этих мягких французских булок, да выпей чаю"
	countLetters(str)
}
