// 문제 04-extra2: Run-Length Encoding (RLE)
//
// [같은 패턴]
//   - rune 순회, strings.Builder, fmt.Fprintf
//
// [문제]
//   문자열 s를 다음 규칙으로 압축한 결과를 반환.
//   - 연속된 같은 문자를 [문자][개수] 로 표현
//   - 단, 압축 결과가 원본보다 길면 원본을 반환
//
// [예시]
//   Compress("aaabbc")    -> "a3b2c1"
//   Compress("abcd")      -> "abcd"      // 압축이 더 길어짐
//   Compress("aaa")       -> "a3"
//   Compress("")          -> ""
//   Compress("aabcccccaaa") -> "a2b1c5a3"
//
// [힌트]
//   - strings.Builder 로 효율적으로 문자열 조립
//   - fmt.Fprintf(&b, "%c%d", ch, cnt) 패턴
package main

import "fmt"

func Compress(s string) string {
	// TODO: 구현하세요.
	return ""
}

func main() {
	tests := []struct {
		s    string
		want string
	}{
		{"aaabbc", "a3b2c1"},
		{"abcd", "abcd"},
		{"aaa", "a3"},
		{"", ""},
		{"aabcccccaaa", "a2b1c5a3"},
	}

	for _, tc := range tests {
		got := Compress(tc.s)
		fmt.Printf("Compress(%q) = %q  | pass=%v\n", tc.s, got, got == tc.want)
	}
}
