package main

import "fmt"

func main() {
	q, r := divide(10, 3)
	fmt.Printf("몫: %d, 나머지: %d\n", q, r)

	// _로 버리고 싶은 반환값을 버릴 수 있음
	q2, _ := divide(30, 8)
	fmt.Printf("몫: %d", q2)
}

// 1. 변수명 자료형 순으로 parameter 선언
// 2. 반환값을 여러개 선언 가능, 단일선언의 경우 ()괄호 생략 가능
// 실제 개발에서 사용 시, 반환값과 별개로 error용으로 이용할 수 있을 것 같아서 굉장히 매력적인 방식으로 보임
func divide(a int, b int) (int, int) {
	quotient := a / b
	remainder := a % b
	return quotient, remainder
}
