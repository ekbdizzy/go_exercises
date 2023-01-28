package main

import "fmt"

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

func main() {
	showGenerateCheck()
}
