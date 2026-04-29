// 문제 05-extra1: 가장 긴 중복 없는 부분문자열 길이
//
// [같은 패턴]
//   - 해시맵으로 "마지막으로 본 위치" 추적, 단일 패스
//
// [문제]
//   문자열 s에서 중복 문자가 없는 가장 긴 연속 부분문자열의 길이를 반환.
//
// [예시]
//   LongestUnique("abcabcbb") -> 3   // "abc"
//   LongestUnique("bbbbb")    -> 1   // "b"
//   LongestUnique("pwwkew")   -> 3   // "wke"
//   LongestUnique("")         -> 0
//
// [힌트]
//   슬라이딩 윈도우 + map[byte]int (마지막 등장 인덱스)
//   - left: 윈도우 시작
//   - i가 진행하며 s[i]가 윈도우 안에 이미 있으면 left를 갱신
package main

import "fmt"

func LongestUnique(s string) int {
	// TODO: 구현하세요.
	return 0
}

func main() {
	tests := []struct {
		s    string
		want int
	}{
		{"abcabcbb", 3},
		{"bbbbb", 1},
		{"pwwkew", 3},
		{"", 0},
		{"abcdef", 6},
	}

	for _, tc := range tests {
		got := LongestUnique(tc.s)
		fmt.Printf("LongestUnique(%q) = %d  | pass=%v\n", tc.s, got, got == tc.want)
	}
}
