package zi

import "fmt"

func ExampleMapGet() {
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

	fmt.Printf("%v\n", v1)
	fmt.Printf("%v\n", v2)
	fmt.Printf("%v\n", v3)
	fmt.Printf("%v\n", v4)
	// Output: 1
	// 2
	// 3
	// 4
}
