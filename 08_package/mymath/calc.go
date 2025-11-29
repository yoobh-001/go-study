package mymath // 상위 폴더이름과 맞춰야함

// 함수가 대문자일경우 외부 패키지에서 사용가능
func Add(a int, b int) int {
	return a + b
}

// 함수가 소문자일경우 외부 패키지에서 사용불가
// 단, 같은패키지(mymath)에서는 사용가능
func minus(a int, b int) int {
	return a - b
}

