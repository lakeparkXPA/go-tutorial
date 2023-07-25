package integers

import (
	"fmt"
	"testing"
)

// TDD - 먼저 테스트 코드 생성 -> 그다음 기능 구현
func TestAdder(t *testing.T) {
	sum := Add(2, 2)
	expected := 4

	if sum != expected {
		t.Errorf("expected '%d' but got '%d'", expected, sum)
	}
}

func ExampleAdd() {
	sum := Add(2, 10)
	fmt.Println(sum)
}
