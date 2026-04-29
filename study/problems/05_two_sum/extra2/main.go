// 문제 05-extra2: Group Anagrams
//
// [같은 패턴]
//   - 해시맵 키로 "정규화된 문자열" 활용, 그룹핑
//
// [문제]
//   문자열 슬라이스 strs를 애너그램끼리 그룹화한 [][]string 으로 반환.
//   각 그룹의 내부 순서는 입력 순서를 유지.
//   바깥 그룹들의 순서는 신경 쓰지 않는다 (테스트는 set 비교).
//
// [예시]
//   GroupAnagrams(["eat","tea","tan","ate","nat","bat"])
//     -> [["eat","tea","ate"], ["tan","nat"], ["bat"]]
//
// [힌트]
//   각 단어의 글자를 정렬해 만든 키를 사용:
//     "eat" → 정렬 → "aet"
//   같은 키를 가진 단어들이 같은 애너그램 그룹.
package main

import (
	"fmt"
	"sort"
)

func GroupAnagrams(strs []string) [][]string {
	// TODO: 구현하세요.
	_ = sort.Strings
	return nil
}

func main() {
	got := GroupAnagrams([]string{"eat", "tea", "tan", "ate", "nat", "bat"})
	fmt.Printf("GroupAnagrams = %v\n", got)
	fmt.Printf("그룹 개수 = %d (기대 3)\n", len(got))

	// 그룹 내용 검증 (집합 비교)
	expected := map[string]bool{
		"eat,tea,ate": true,
		"tan,nat":     true,
		"bat":         true,
	}
	pass := len(got) == 3
	for _, g := range got {
		key := join(g)
		if !expected[key] {
			pass = false
		}
	}
	fmt.Printf("pass=%v\n", pass)
}

func join(g []string) string {
	out := ""
	for i, s := range g {
		if i > 0 {
			out += ","
		}
		out += s
	}
	return out
}
