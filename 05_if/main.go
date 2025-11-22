package main

import "fmt"

func main() {
	age := 15

	// if 문의경우 {} 중괄호 위치가 문법으로 정해져있으며,
	// () 괄호를 사용하지 않음

	if age >= 20 {
		fmt.Println("성인입니다.")
	} else if age >= 14 {
		fmt.Println("청소년입니다.")
	} else {
		fmt.Println("어린이입니다.")
	}

	if time := 23; time >= 22 {
		fmt.Printf("저녁입니다. %v시", time)
	}

	// for 선언 변수처럼 if 선언 변수는 밖에서 사용못함
	// fmt.Println(time)
}
