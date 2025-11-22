package main

import "fmt"

func main() {
	// 1. 정석 선언방식
	var name string = "김철수"
	var age int = 18

	// 2. 자료형 생략 선언방식
	var school = "ㅇㅇ고등학교"

	// 3. 축약 선언방식
	weight := 70.5
	height := 175.5

	fmt.Println(name, age, school, weight, height)
	fmt.Println(grade)
}

// 1,2번의 경우 함수 외부선언 가능
var grade = "2학년"

// 3번의 경우 함수 외부선언 불가능
// grade2: = "3학년"
