package d3

import (
	"reflect"
	"testing"
)

func TestGetMatches(t *testing.T) {
	t.Skip()
	in := "~-$'&&^!<mul(959,544)@where()#:?%)/)-mul(730,399)::"
	out := [][]int{{959, 544}, {730, 399}}
	result := getMatches(in)
	if !reflect.DeepEqual(result, out) {
		t.Errorf("got %v, expected %v", result, out)
	}
}
