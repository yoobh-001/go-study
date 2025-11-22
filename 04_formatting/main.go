package main

import "fmt"

func main() {
	name := "캐릭터"
	level := 100
	hp := 1000.591

	fmt.Println("이름: ", name, "레벨: ", level, "HP: ", hp)

	// %s: 문자열 출력
	// %d: 정수 출력
	// %f: 실수 출력
	fmt.Printf("이름: %s 레벨: %d HP: %f", name, level, hp)
	fmt.Println()
	fmt.Printf("HP %.2f", hp) // 소수점 2자리까지 출력
	fmt.Println()

	// %v: 자동으로 자료형 추론
	fmt.Printf("이름: %v, 레벨: %v, HP: %v", name, level, hp)

	fmt.Println()
	fmt.Printf("HP %.2v", hp) // 1e+03으로 출력되어 기대한대로 안됨
}
