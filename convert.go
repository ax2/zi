package zi

import "strconv"

type IntType interface {
	int | int32 | int64
}

func Int[T IntType](s string) T {
	tmp, _ := strconv.ParseInt(s, 10, 0)

	return T(tmp)
}

func Float(s string) float64 {
	tmp, _ := strconv.ParseFloat(s, 10)

	return float64(tmp)
}
