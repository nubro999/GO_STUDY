// 문제 07-extra2: Storage 인터페이스 (모킹 패턴)
//
// [같은 패턴]
//   - 인터페이스로 의존성 분리, 구현체 교체 가능
//   - 실무에서 가장 자주 쓰이는 인터페이스 활용 (테스트 모킹)
//
// [문제]
//   다음 인터페이스를 만족하는 두 구현체를 작성하시오.
//
//     type Storage interface {
//         Save(key, value string) error
//         Load(key string) (string, error)
//     }
//
//   1) MemoryStorage: map[string]string 으로 저장
//   2) NullStorage: 아무것도 저장하지 않음. Save는 항상 nil, Load는 항상 ("", ErrNotFound)
//
//   ErrNotFound 는 sentinel error로 정의:
//     var ErrNotFound = errors.New("key not found")
//
//   그리고 SaveAll(s Storage, items map[string]string) error 함수를 작성:
//     - items의 모든 항목을 s에 저장
//     - 도중에 에러가 발생하면 즉시 반환
package main

import (
	"errors"
	"fmt"
)

var ErrNotFound = errors.New("key not found")

type Storage interface {
	Save(key, value string) error
	Load(key string) (string, error)
}

// TODO: MemoryStorage, NullStorage 정의 및 메서드 구현

func SaveAll(s Storage, items map[string]string) error {
	// TODO: 구현하세요.
	return nil
}

func main() {
	// 아래 주석을 해제하고 Storage 구현체를 작성한 뒤 실행하세요.
	//
	// mem := &MemoryStorage{}
	// _ = SaveAll(mem, map[string]string{"a": "1", "b": "2"})
	// v, err := mem.Load("a")
	// fmt.Printf("MemoryStorage Load(a) = %q, err=%v  | pass=%v\n", v, err, v == "1" && err == nil)
	//
	// _, err = mem.Load("missing")
	// fmt.Printf("MemoryStorage Load(missing) err=%v  | pass=%v\n", err, errors.Is(err, ErrNotFound))
	//
	// null := &NullStorage{}
	// _ = null.Save("x", "y")  // OK, 무시
	// _, err = null.Load("x")
	// fmt.Printf("NullStorage Load = err=%v  | pass=%v\n", err, errors.Is(err, ErrNotFound))

	fmt.Println("Storage 구현체를 작성한 뒤 main()의 주석을 해제하세요.")
}
