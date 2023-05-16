package arrayUtils

func Map[T, U any](ts []T, f func(T) U) []U {
	us := make([]U, len(ts))
	for i := range ts {
		us[i] = f(ts[i])
	}
	return us
}
func Filter[T any](ts []T, f func(T) bool) []T {
	filtered := make([]T, 0)
	for _, v := range ts {
		if f(v) {
			filtered = append(filtered, v)
		}
	}
	return filtered
}
