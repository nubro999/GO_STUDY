// 문제 04-extra1: 회문(Palindrome) 판별
//
// [같은 패턴]
//   - rune, []rune 변환, 양 끝에서 좁혀가는 투포인터
//
// [문제]
//   문자열 s가 회문인지 판별하시오.
//   - 대소문자 구분 안 함
//   - 공백/문장부호 무시 (영문/숫자만 고려)
//
// [예시]
//   IsPalindrome("A man, a plan, a canal: Panama") -> true
//   IsPalindrome("race a car")                     -> false
//   IsPalindrome("")                               -> true
//
// [힌트]
//   1) 정규화: 소문자 + 영숫자만 남기기
//   2) 양끝 인덱스 i, j를 좁혀가며 비교
package main

import "fmt"

func IsPalindrome(s string) bool {
	// TODO: 구현하세요.
	return false
}

func main() {
	tests := []struct {
		s    string
		want bool
	}{
		{"A man, a plan, a canal: Panama", true},
		{"race a car", false},
		{"", true},
		{"a", true},
		{"ab", false},
		{"No lemon, no melon", true},
	}

	for _, tc := range tests {
		got := IsPalindrome(tc.s)
		fmt.Printf("IsPalindrome(%q) = %v  | pass=%v\n", tc.s, got, got == tc.want)
	}
}
