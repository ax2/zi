package zi

import "testing"

func TestMapGet(t *testing.T) {
	m := map[string]interface{}{
		"a": 1,
		"b": "2",
		"c": map[string]interface{}{
			"c1": 3,
			"c2": map[string]interface{}{
				"c21": 4,
			},
		},
	}

	v1 := MapGet[int](m, "a")
	v2 := MapGet[string](m, "b")
	v3 := MapGet[int](m, "c", "c1")
	v4 := MapGet[int](m, "c", "c2", "c21")

	if v1 != 1 {
		t.Errorf("mapget failed:%v", v1)
	}
	if v2 != "2" {
		t.Errorf("mapget failed:%v", v2)
	}
	if v3 != 3 {
		t.Errorf("mapget failed:%v", v3)
	}
	if v4 != 4 {
		t.Errorf("mapget failed:%v", v4)
	}
}
