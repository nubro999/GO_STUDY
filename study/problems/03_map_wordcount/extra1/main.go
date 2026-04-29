// 문제 03-extra1: 첫 등장 인덱스 매핑
//
// [같은 패턴]
//   - map[K]V, comma-ok idiom, range
//
// [문제]
//   문자열 슬라이스 words가 주어진다.
//   각 단어가 처음 등장한 인덱스를 매핑한 map을 반환하시오.
//   같은 단어가 다시 나타나도 첫 인덱스를 유지.
//
// [예시]
//   FirstIndex(["go","is","fun","go","fun"])
//     -> {"go":0, "is":1, "fun":2}
//
// [힌트]
//   comma-ok 로 키가 이미 있는지 검사 후 없을 때만 저장.
package main

import "fmt"
func FirstIndex(words []string) map[string]int {
	// TODO: 구현하세요.
	return nil
}

func main() {
	tests := []struct {
		words []string
		want  map[string]int
	}{
		{[]string{"go", "is", "fun", "go", "fun"}, map[string]int{"go": 0, "is": 1, "fun": 2}},
		{[]string{}, map[string]int{}},
		{[]string{"a", "a", "a"}, map[string]int{"a": 0}},
	}

	for _, tc := range tests {
		got := FirstIndex(tc.words)
		fmt.Printf("FirstIndex(%v) = %v  | pass=%v\n", tc.words, got, equalMap(got, tc.want))
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
