package main

import "fmt"

// 캐릭터 구조체 생성
type Character struct {
	Name  string
	Level int
}

// 잘못된 예시) 원본 구조체를 변경하는것이 아닌, 복사본을 생성 후 변경
func LevelUpFake(c Character) {
	c.Level++
	fmt.Println("-- 복사본 캐릭터 레벨", c.Level)
}

// 올바른 예시) 원본 구조체를 변경
// *를 사용하여 원본을 이용함을 명시
func LevelUpReal(c *Character) {
	c.Level++
	fmt.Println("-- 원본 캐릭터 레벨", c.Level)
}

func main() {

	// 내 캐릭터 레벨
	myC := Character{Name: "민수", Level: 1}
	fmt.Println("초기 레벨", myC.Level)

	LevelUpFake(myC)
	fmt.Println("가짜 레벨업 후 레벨:", myC.Level)

	LevelUpReal(&myC)
	fmt.Println("실제 레벨업 후 레벨:", myC.Level)
}
