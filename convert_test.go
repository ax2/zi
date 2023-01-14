package zi

import "testing"

func TestStringToInt(t *testing.T) {
	s := "1234"
	n := Int[int64](s)
	if n != 1234 {
		t.Errorf("convert failed:%s => %d", s, n)
	}
}
