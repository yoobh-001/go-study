package main

import (
	"fmt"
	// 가져올 패키지 경로
	"go-study/08_package/mymath"
)

func main() {
	result := mymath.Add(1, 2)
	fmt.Println(result)

	// 해당 패키지는 외부패키지고, 함수이름이 소문자이므로 사용불가
	// 함수 첫글자가 소문자냐 대문자냐에따라 외부패키지에서 사용가능 여부가 결정됨
	// result2 := mymath.minus(1, 2)
}
