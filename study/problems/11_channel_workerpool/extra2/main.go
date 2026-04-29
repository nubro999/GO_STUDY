// 문제 11-extra2: Fan-In (여러 채널을 하나로 병합)
//
// [같은 패턴]
//   - chan + close, sync.WaitGroup, 멀티 producer 패턴
//
// [문제]
//   여러 입력 채널 (...<-chan int) 을 하나의 출력 채널로 병합한다.
//   모든 입력 채널이 close 되면 출력 채널도 close.
//
//   func FanIn(channels ...<-chan int) <-chan int
//
// [예시]
//   ch1: 1, 2, 3 (close)
//   ch2: 10, 20  (close)
//   ch3: 100     (close)
//   FanIn(ch1, ch2, ch3) → 위 6개 값을 (순서 무관하게) 모두 받고 close
//
// [힌트]
//   - 각 입력 채널을 별도 고루틴이 읽어서 out 으로 forward.
//   - WaitGroup 으로 모든 forwarder 종료 대기 → 그 후 close(out).
package main

import (
	"fmt"
	"sort"
	"sync"
)

func FanIn(channels ...<-chan int) <-chan int {
	out := make(chan int)
	// TODO: 구현
	_ = sync.WaitGroup{}
	close(out)
	return out
}

func main() {
	make1 := func(vals ...int) <-chan int {
		ch := make(chan int)
		go func() {
			defer close(ch)
			for _, v := range vals {
				ch <- v
			}
		}()
		return ch
	}

	merged := FanIn(make1(1, 2, 3), make1(10, 20), make1(100))

	var got []int
	for v := range merged {
		got = append(got, v)
	}
	sort.Ints(got)

	want := []int{1, 2, 3, 10, 20, 100}
	pass := len(got) == len(want)
	for i := range got {
		if got[i] != want[i] {
			pass = false
		}
	}
	fmt.Printf("FanIn (sorted) = %v  | pass=%v\n", got, pass)
}
