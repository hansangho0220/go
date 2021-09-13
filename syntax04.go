package main

import (
	"fmt"
	"math/rand"
	"time"
)

//난수 추출된 수의 소수 판정 프로그램 v0.1
func main() {
	seed := time.Now().Unix()
	rand.Seed(seed)

	number := rand.Intn(150) + 2
	fmt.Println("임의로 추출된 수: ", number)

	count := 0
	for i := 1; i <= number; i++ {
		if number%i == 0 {
			count++
		}
	}

	if count == 2 {
		fmt.Println(number, "는 소수입니다.")
	} else {
		fmt.Println(number, "는 소수가 아닙니다.")
	}
}
