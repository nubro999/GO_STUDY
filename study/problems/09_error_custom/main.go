// 문제 09: 커스텀 에러와 에러 래핑
//
// [학습 포인트]
//   - error 인터페이스 만족 (Error() string 메서드)
//   - fmt.Errorf("%w", err) 로 에러 래핑
//   - errors.Is / errors.As 로 에러 식별
//   - sentinel error vs custom error type 의 차이
//
// [문제]
//   1) DivisionByZeroError 라는 커스텀 에러 타입을 정의하시오.
//      - 분자(Numerator)를 필드로 갖는다.
//      - Error() string 메서드 구현. 메시지 예: "cannot divide 10 by zero"
//
//   2) SafeDivide(a, b int) (int, error) 함수를 구현하시오.
//      - b == 0 인 경우 *DivisionByZeroError 반환
//      - 그 외에는 a/b 와 nil 반환
//
//   3) ComputeAverage(nums []int, divisor int) (int, error) 를 구현하시오.
//      - 합을 divisor로 나눈 결과를 반환.
//      - SafeDivide를 호출하고, 에러가 발생하면
//        fmt.Errorf("compute average failed: %w", err) 로 래핑하여 반환.
//
//   main에서는 errors.As 로 *DivisionByZeroError 를 추출할 수 있어야 함.
package main

import (
	"errors"
	"fmt"
)

// TODO: DivisionByZeroError 타입과 Error() 메서드 정의

func SafeDivide(a, b int) (int, error) {
	// TODO: 구현하세요.
	return 0, nil
}

func ComputeAverage(nums []int, divisor int) (int, error) {
	// TODO: 구현하세요. (합산 후 SafeDivide 호출, 에러 래핑)
	return 0, nil
}

func main() {
	// 정상 케이스
	if v, err := SafeDivide(10, 2); err == nil {
		fmt.Printf("SafeDivide(10, 2) = %d  | pass=%v\n", v, v == 5)
	}

	// 에러 케이스 — errors.As 로 커스텀 타입 추출
	_, err := ComputeAverage([]int{1, 2, 3, 4}, 0)
	fmt.Printf("ComputeAverage error: %v\n", err)

	// var dze *DivisionByZeroError
	// if errors.As(err, &dze) {
	//     fmt.Printf("  -> 분자값 추출 성공: %d  | pass=%v\n", dze.Numerator, dze.Numerator == 10)
	// } else {
	//     fmt.Println("  -> errors.As 실패: 타입을 추출하지 못함")
	// }
	// 위 주석을 해제하고 DivisionByZeroError 를 정의한 뒤 실행하세요.
	_ = errors.As
}
