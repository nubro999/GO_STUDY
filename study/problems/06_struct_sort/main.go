// 문제 06: 구조체 슬라이스 정렬
//
// [학습 포인트]
//   - struct 정의 및 슬라이스
//   - sort.Slice (less 함수)
//   - 다중 키 정렬 패턴 (1차/2차 키)
//
// [문제]
//   User 구조체 슬라이스를 다음 규칙으로 정렬하시오.
//     1차 정렬: Age 오름차순
//     2차 정렬: Name 사전순(오름차순) — Age가 같을 때
//   원본 슬라이스를 in-place 정렬한다.
//
// [예시]
//   입력: [{Bob 30} {Alice 25} {Charlie 25}]
//   출력: [{Alice 25} {Charlie 25} {Bob 30}]
package main

import "fmt"

type User struct {
	Name string
	Age  int
}

func SortUsers(users []User) {
	// TODO: 구현하세요.
	// 힌트: sort.Slice(users, func(i, j int) bool { ... })
}

func main() {
	users := []User{
		{"Bob", 30},
		{"Alice", 25},
		{"Charlie", 25},
		{"Dave", 40},
		{"Eve", 25},
	}
	want := []User{
		{"Alice", 25},
		{"Charlie", 25},
		{"Eve", 25},
		{"Bob", 30},
		{"Dave", 40},
	}

	SortUsers(users)

	pass := len(users) == len(want)
	for i := range users {
		if users[i] != want[i] {
			pass = false
		}
	}
	fmt.Printf("SortUsers result = %v\n", users)
	fmt.Printf("expected         = %v\n", want)
	fmt.Printf("pass=%v\n", pass)
}
