// 문제 07: 인터페이스와 다형성
//
// [학습 포인트]
//   - interface 정의와 만족 (implicit satisfaction)
//   - 메서드 셋 (값 리시버 vs 포인터 리시버)
//   - 인터페이스 슬라이스를 통한 다형적 처리
//
// [문제]
//   다음 인터페이스를 만족하는 도형 타입 3개를 구현하시오.
//
//     type Shape interface {
//         Area() float64
//         Perimeter() float64
//     }
//
//   구현할 타입:
//     - Circle{Radius float64}
//     - Rectangle{Width, Height float64}
//     - Triangle{A, B, C float64}   // 세 변의 길이, 헤론의 공식 사용
//
//   그리고 TotalArea(shapes []Shape) float64 함수를 작성하시오.
//
// [헤론의 공식]
//   s = (a+b+c)/2
//   Area = sqrt(s(s-a)(s-b)(s-c))
package main

import (
	"fmt"
	"math"
)

type Shape interface {
	Area() float64
	Perimeter() float64
}

type Circle struct {
	Radius float64
}

// TODO: Circle의 Area, Perimeter 구현
func (c Circle) Area() float64 {
	return math.Pi * c.Radius * c.Radius
}

func (c Circle) Perimeter() float64 {
	return 2 * math.Pi * c.Radius
}

type Rectangle struct {
	Width, Height float64
}

func (r Rectangle) Area() float64 {
	return r.Width * r.Height
}

func (r Rectangle) Perimeter() float64 {
	return 2 * (r.Width + r.Height)
}

// TODO: Rectangle의 Area, Perimeter 구현

type Triangle struct {
	A, B, C float64
}

func (t Triangle) Area() float64 {
	s := (t.A + t.B + t.C) / 2
	return math.Sqrt(s * (s - t.A) * (s - t.B) * (s - t.C))
}

func (t Triangle) Perimeter() float64 {
	return t.A + t.B + t.C
}

// TODO: Triangle의 Area, Perimeter 구현 (헤론의 공식)

func TotalArea(shapes []Shape) float64 {
	// TODO: 모든 도형의 Area 합계를 반환
	var total float64
	for _, s := range shapes {
		total += s.Area()
	}
	return total
}

func main() {
	// 아래 주석을 해제하고 Circle/Rectangle/Triangle의 메서드를 구현한 뒤 실행하세요.
	//
	shapes := []Shape{
		Circle{Radius: 1},              // π ≈ 3.14159
		Rectangle{Width: 3, Height: 4}, // 12
		Triangle{A: 3, B: 4, C: 5},     // 6 (3-4-5 직각삼각형)
	}
	
	for _, s := range shapes {
		fmt.Printf("%-25T  Area=%.4f  Perimeter=%.4f\n", s, s.Area(), s.Perimeter())
	}
	
	total := TotalArea(shapes)
	want := math.Pi + 12 + 6
	fmt.Printf("TotalArea = %.4f  (want %.4f)  pass=%v\n",
		total, want, math.Abs(total-want) < 1e-9)

	_ = math.Pi
	fmt.Println("Shape 구현체를 작성한 뒤 main()의 주석을 해제하세요.")
}
