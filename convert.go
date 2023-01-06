package zi

import "strconv"

type Int interface {
	int | int32 | int64
}

func StringToInt[T Int](s string) T {
	tmp, _ := strconv.ParseInt(s, 10, 0)

	return T(tmp)
}
