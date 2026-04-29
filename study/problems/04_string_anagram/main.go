// 문제 04: 애너그램 판별
//
// [학습 포인트]
//   - rune 자료형 (UTF-8 대응)
//   - []rune 변환과 정렬
//   - 문자열 vs []byte vs []rune 차이 이해
//
// [문제]
//   두 문자열 a, b가 애너그램인지 판별하시오.
//   - 애너그램: 같은 문자를 같은 횟수만큼 가진 두 문자열
//   - 대소문자/공백은 무시
//
// [예시]
//   IsAnagram("listen", "silent")        -> true
//   IsAnagram("Dirty Room", "DormItory") -> true
//   IsAnagram("hello", "world")          -> false
//
// [구현 방법 두 가지]
//   1) 양쪽을 정렬해서 비교 ([]rune → sort)
//   2) map[rune]int 으로 빈도수 카운트 후 비교
//   둘 다 구현해보면 좋음. 본 파일은 한 가지를 채워주세요.
package main

import "fmt"

func IsAnagram(a, b string) bool {
	// TODO: 구현하세요.
	return false
}

func main() {
	tests := []struct {
		a, b string
		want bool
	}{
		{"listen", "silent", true},
		{"Dirty Room", "DormItory", true},
		{"hello", "world", false},
		{"", "", true},
		{"a", "ab", false},
	}

	for _, tc := range tests {
		got := IsAnagram(tc.a, tc.b)
		fmt.Printf("IsAnagram(%q, %q) = %v  | pass=%v\n", tc.a, tc.b, got, got == tc.want)
	}
}
