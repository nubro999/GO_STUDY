// 문제 03-extra2: 첫 글자별 그룹화
//
// [같은 패턴]
//   - map[K][]V, append, range
//
// [문제]
//   문자열 슬라이스 words를 첫 글자(byte)별로 그룹화한 map[byte][]string 을 반환.
//   각 그룹 내 순서는 입력 순서를 유지.
//   빈 문자열은 무시.
//
// [예시]
//   GroupByFirst(["apple","banana","avocado","blueberry","cherry"])
//     -> {'a': ["apple","avocado"], 'b': ["banana","blueberry"], 'c': ["cherry"]}
//
// [힌트]
//   m[k] = append(m[k], v) — 키 없으면 nil에 append (Go에서 안전)
package main

import "fmt"

func GroupByFirst(words []string) map[byte][]string {
	// TODO: 구현하세요.
	return nil
}

func main() {
	got := GroupByFirst([]string{"apple", "banana", "avocado", "blueberry", "cherry", ""})

	want := map[byte][]string{
		'a': {"apple", "avocado"},
		'b': {"banana", "blueberry"},
		'c': {"cherry"},
	}

	pass := len(got) == len(want)
	for k, vs := range want {
		if len(got[k]) != len(vs) {
			pass = false
			break
		}
		for i := range vs {
			if got[k][i] != vs[i] {
				pass = false
			}
		}
	}
	fmt.Printf("GroupByFirst = %v  | pass=%v\n", got, pass)
}
