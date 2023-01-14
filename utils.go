package zi

// Ternary is a 1 line if/else statement.
// Example: result := zi.Ternary[int64](isOk, 1, 0)
// See: https://github.com/samber/lo
func Ternary[T any](condition bool, ifOutput T, elseOutput T) T {
	if condition {
		return ifOutput
	}

	return elseOutput
}
