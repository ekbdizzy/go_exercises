package main

import (
	"fmt"
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

func hasTwice(nums []int) bool {
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

func main() {
	//showGenerateCheck()
	nums := []int{1, 2, 3, 4, 5, 5}
	fmt.Println(hasTwice(nums))

	nums = []int{1, 2, 3, 4, 5, 6}
	fmt.Println(hasTwice(nums))
}
