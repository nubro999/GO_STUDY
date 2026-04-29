// 문제 06-extra2: 이벤트 시간순 정렬
//
// [같은 패턴]
//   - struct 슬라이스 정렬, 3-way 비교
//
// [문제]
//   Event 슬라이스를 다음 규칙으로 정렬하시오.
//     1차: Start 오름차순
//     2차: End 오름차순 (시작이 같으면 빨리 끝나는 게 먼저)
//     3차: Title 사전순
//
// [예시]
//   [{B 10 20} {A 10 15} {C 5 30} {D 10 15}]
//   →
//   [{C 5 30} {A 10 15} {D 10 15} {B 10 20}]
package main

import "fmt"

type Event struct {
	Title string
	Start int
	End   int
}

func SortEvents(events []Event) {
	// TODO: 구현하세요.
}

func main() {
	events := []Event{
		{"B", 10, 20},
		{"A", 10, 15},
		{"C", 5, 30},
		{"D", 10, 15},
	}
	want := []Event{
		{"C", 5, 30},
		{"A", 10, 15},
		{"D", 10, 15},
		{"B", 10, 20},
	}

	SortEvents(events)

	pass := len(events) == len(want)
	for i := range events {
		if events[i] != want[i] {
			pass = false
		}
	}
	fmt.Printf("SortEvents = %v\n", events)
	fmt.Printf("pass=%v\n", pass)
}
