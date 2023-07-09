package iteration

import (
	"fmt"
	"testing"
)

const Count = 4

func TestRepeat(t *testing.T) {
	repeat := Repeat("a", Count)
	expect := "aaaa"

	if repeat != expect {
		t.Errorf("expect %q but got %q", expect, repeat)
	}
}

func BenchmarkRepeat(b *testing.B) {
	for i := 0; i < b.N; i++ {
		Repeat("a", Count)
	}
}

func ExampleRepeat() {
	repeat := Repeat("b", 3)
	fmt.Println(repeat)
	// Output: bbb
}

func Repeat(char string, count int) string {
	var repeat string
	for i := 0; i < count; i++ {
		repeat += char
	}
	return repeat
}
