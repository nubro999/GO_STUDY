// 문제 03: 단어 빈도수 계산
//
// [학습 포인트]
//   - map[K]V 선언/초기화 (make vs literal)
//   - map의 zero-value 활용 (m[k]++ 패턴)
//   - strings.Fields, strings.ToLower
//
// [문제]
//
//	문장 s가 주어졌을 때 각 단어가 등장한 횟수를 map으로 반환하시오.
//	- 단어는 공백(스페이스/탭/개행)으로 구분된다.
//	- 대소문자는 구분하지 않는다 (모두 소문자로).
//	- 문장 양옆/사이의 공백은 무시.
//
// [예시]
//
//	WordCount("Go is fun. Go is fast.")
//	  -> map[string]int{"go":2, "is":2, "fun.":1, "fast.":1}
//
// [심화]
//
//	문장 부호(.,!?)도 제거해보고 싶다면 strings.Trim 또는 strings.Map 사용.
package main

import (
	"fmt"
	"strings"
)

func WordCount(s string) map[string]int {
	// TODO: 구현하세요.
	out := map[string]int{}
	for _, w := range strings.Fields(strings.ToLower(s)) {
		out[w]++
	}
	
	return out
}

func main() {
	tests := []struct {
		s    string
		want map[string]int
	}{
		{
			"Go is fun Go is fast",
			map[string]int{"go": 2, "is": 2, "fun": 1, "fast": 1},
		},
		{
			"  hello   WORLD hello  ",
			map[string]int{"hello": 2, "world": 1},
		},
		{
			"",
			map[string]int{},
		},
	}

	for _, tc := range tests {
		got := WordCount(tc.s)
		fmt.Printf("WordCount(%q) = %v  | pass=%v\n", tc.s, got, equalMap(got, tc.want))
	}
}

func equalMap(a, b map[string]int) bool {
	if len(a) != len(b) {
		return false
	}
	for k, v := range a {
		if b[k] != v {
			return false
		}
	}
	return true
}
