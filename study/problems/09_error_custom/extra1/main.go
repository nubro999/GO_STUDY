// 문제 09-extra1: 다중 에러 타입 분기 처리
//
// [같은 패턴]
//   - 커스텀 에러 타입 + errors.As 로 분기
//
// [문제]
//   다음 두 에러 타입을 정의하시오.
//
//     type NotFoundError struct{ Resource string }
//     type ValidationError struct{ Field, Reason string }
//
//   각각 Error() 메서드를 구현 (메시지 자유, 단 Resource/Field/Reason 포함).
//
//   그리고 Classify(err error) string 함수를 작성:
//     - err 이 *NotFoundError → "not_found"
//     - err 이 *ValidationError → "validation"
//     - err 이 nil → "none"
//     - 그 외 (래핑된 위 두 타입 포함 안 됨) → "unknown"
//   래핑된 에러도 분류할 수 있도록 errors.As 사용.
//
// [예시]
//   Classify(nil)                                                -> "none"
//   Classify(&NotFoundError{"user"})                             -> "not_found"
//   Classify(fmt.Errorf("wrap: %w", &ValidationError{...}))      -> "validation"
//   Classify(errors.New("misc"))                                 -> "unknown"
package main

import (
	"errors"
	"fmt"
)

// TODO: NotFoundError, ValidationError 정의

func Classify(err error) string {
	// TODO: 구현하세요. errors.As 활용.
	_ = errors.As
	return ""
}

func main() {
	// 아래 주석을 해제하고 에러 타입을 구현한 뒤 실행하세요.
	//
	// tests := []struct {
	// 	err  error
	// 	want string
	// }{
	// 	{nil, "none"},
	// 	{&NotFoundError{Resource: "user"}, "not_found"},
	// 	{&ValidationError{Field: "email", Reason: "invalid"}, "validation"},
	// 	{fmt.Errorf("query failed: %w", &NotFoundError{Resource: "post"}), "not_found"},
	// 	{fmt.Errorf("submit: %w", &ValidationError{Field: "age", Reason: "negative"}), "validation"},
	// 	{errors.New("misc"), "unknown"},
	// }
	// for _, tc := range tests {
	// 	got := Classify(tc.err)
	// 	fmt.Printf("Classify(%v) = %q  | pass=%v\n", tc.err, got, got == tc.want)
	// }

	fmt.Println("에러 타입을 구현한 뒤 main()의 주석을 해제하세요.")
}
