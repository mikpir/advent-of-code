package d1

import (
	"strings"
	"testing"
)

func Test2(t *testing.T) {
	got := RunB(strings.NewReader(`3   4
4   3
2   5
1   3
3   9
3   3`))
	expected := 31
	if got != expected {
		t.Errorf("Expected %d, got %d", expected, got)
	}
}
