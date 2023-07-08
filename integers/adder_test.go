package integers

import (
	"fmt"
	"testing"
)

func TestAdder(t *testing.T) {
	sum := Add(2, 2, "")
	expect := 4

	if sum != expect {
		t.Errorf("expect '%d' but got '%d'", expect, sum)
	}

}

func Add(x, y int, p string) int {
	return x + y
}

func ExampleAdd() {
	sum := Add(1, 5, "")
	fmt.Println(sum)
	// Output: 6
}
