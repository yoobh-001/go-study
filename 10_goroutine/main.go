package main

import (
	"fmt"
	"sync"
	"time"
)

func work(name string) {
	fmt.Println(name, "work start")
	time.Sleep(1 * time.Second) // 1초 sleep
	fmt.Println(name, "work end")
}

// sync.WaitGroup을 사용한 Goroutine
func workWithGoroutine(name string, wg *sync.WaitGroup) {
	// 함수가 끝나면 = defer, Done()을 호출하여 끝났음을 알림
	defer wg.Done()

	fmt.Println(name, "gorutine start")
	time.Sleep(1 * time.Second)
	fmt.Println(name, "gorutine end")
}

func main() {
	fmt.Println("일반 Sleep func 호출시작")
	start := time.Now() //현재시간
	work("직원A")
	work("직원B")
	fmt.Println("일반 Sleep func 소요시간", time.Since(start))

	fmt.Println("Goroutine 호출시작")
	start = time.Now()
	var wg sync.WaitGroup // 대기표 생성?
	wg.Add(2)             // 대기표 2개등록

	// 포인터전달
	go workWithGoroutine("직원A", &wg)
	go workWithGoroutine("직원B", &wg)

	// workWithGoroutine이 Done() 끝날때까지 기다림
	wg.Wait()
	fmt.Println("Goroutine 소요시간", time.Since(start))

	// 정리: go routine은 sync.WaitGroup을 사용하여 컨트롤할 수 있다.
	// 주의사항:
	// 1. sync.WaitGroup으로 컨트롤하지않으면 main함수가 끝나게되면
	// 호출한 goroutine이 끝나지않더라도 프로그램이 종료된다.
	// 2. go routine은 반환값을 받을 수 없으므로 포인터를 전달하여 실제 전달값이 변경되게 사용한다.
}
