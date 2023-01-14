package zi

func MapGetMap(m map[string]interface{}, k string) (v map[string]interface{}) {
	v, _ = m[k].(map[string]interface{})
	return
}

func MapGet[T any](m map[string]interface{}, k ...string) (v T) {
	var ok bool
	m1 := m
	n := len(k) - 1
	if n < 0 {
		return
	}

	for i := 0; i < n; i++ {
		m1, ok = m1[k[i]].(map[string]interface{})
		if !ok {
			return
		}
	}

	v, _ = m1[k[n]].(T)
	return
}
