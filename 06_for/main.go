package main

import "fmt"

func main() {
	fmt.Println("정석 for문")
	for i := 0; i < 10; i++ {
		fmt.Printf("%d ", i)
	}

	fmt.Println("\nwhile형태 for문")
	cnt := 0
	for cnt < 10 {
		fmt.Printf("%d ", cnt)
		cnt++
	}

	fmt.Println("\n무한루프형태 for문")
	infinity := 0
	for {
		fmt.Printf("%d ", infinity)
		infinity++

		// for문 안에서 break 처리
		if infinity >= 10 {
			break
		}
	}

}
