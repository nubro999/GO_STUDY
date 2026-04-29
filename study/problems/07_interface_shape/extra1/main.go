// 문제 07-extra1: Animal 인터페이스
//
// [같은 패턴]
//   - 인터페이스 정의 + 구체 타입 + 다형적 슬라이스 순회
//
// [문제]
//   다음 인터페이스를 만족하는 Dog, Cat, Cow 구조체를 작성하시오.
//
//     type Animal interface {
//         Sound() string
//         Name() string
//     }
//
//   - Dog{Nickname string} → Sound() = "Woof", Name() = Nickname
//   - Cat{Nickname string} → Sound() = "Meow", Name() = Nickname
//   - Cow{Nickname string} → Sound() = "Moo",  Name() = Nickname
//
//   그리고 Chorus(animals []Animal) []string 함수를 작성:
//     - 각 동물별로 "{Name}: {Sound}" 형식의 문자열을 만들어 슬라이스 반환
//
// [예시]
//   Chorus([]Animal{ Dog{"Rex"}, Cat{"Whiskers"} })
//     -> ["Rex: Woof", "Whiskers: Meow"]
package main

import "fmt"

type Animal interface {
	Sound() string
	Name() string
}

// TODO: Dog, Cat, Cow 정의 및 메서드 구현

func Chorus(animals []Animal) []string {
	// TODO: 구현하세요.
	return nil
}

func main() {
	// 아래 주석을 해제하고 Dog/Cat/Cow를 구현한 뒤 실행하세요.
	//
	// got := Chorus([]Animal{
	// 	Dog{Nickname: "Rex"},
	// 	Cat{Nickname: "Whiskers"},
	// 	Cow{Nickname: "Bessie"},
	// })
	// want := []string{"Rex: Woof", "Whiskers: Meow", "Bessie: Moo"}
	// pass := len(got) == len(want)
	// for i := range got {
	// 	if got[i] != want[i] { pass = false }
	// }
	// fmt.Printf("Chorus = %v  | pass=%v\n", got, pass)

	fmt.Println("Dog/Cat/Cow를 구현한 뒤 main()의 주석을 해제하세요.")
}
