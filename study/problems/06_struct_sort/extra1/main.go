// 문제 06-extra1: 학생 랭킹
//
// [같은 패턴]
//   - struct, sort.Slice, 다중 키 정렬
//
// [문제]
//   Student 슬라이스를 다음 규칙으로 정렬하시오.
//     1차: Score 내림차순 (점수가 높은 학생이 앞)
//     2차: Name 사전순 오름차순 (점수가 같으면 이름순)
//   in-place 정렬.
//
// [예시]
//   입력: [{Bob 80} {Alice 90} {Carol 90} {Dan 80}]
//   출력: [{Alice 90} {Carol 90} {Bob 80} {Dan 80}]
package main

import "fmt"

type Student struct {
	Name  string
	Score int
}

func RankStudents(students []Student) {
	// TODO: 구현하세요.
}

func main() {
	students := []Student{
		{"Bob", 80},
		{"Alice", 90},
		{"Carol", 90},
		{"Dan", 80},
		{"Eve", 100},
	}
	want := []Student{
		{"Eve", 100},
		{"Alice", 90},
		{"Carol", 90},
		{"Bob", 80},
		{"Dan", 80},
	}

	RankStudents(students)

	pass := len(students) == len(want)
	for i := range students {
		if students[i] != want[i] {
			pass = false
		}
	}
	fmt.Printf("RankStudents = %v\n", students)
	fmt.Printf("pass=%v\n", pass)
}
