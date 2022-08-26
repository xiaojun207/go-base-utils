package array

func IndexOf[T comparable](arr []T, item T) int {
	for i, s := range arr {
		if s == item {
			return i
		}
	}
	return -1
}

func Contains[T comparable](arr []T, item T) bool {
	idx := IndexOf[T](arr, item)
	return idx > -1
}
