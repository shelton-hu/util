package boolutil

func Bool2Int(v bool) int {
	if v {
		return 1
	}
	return 0
}

func Int2Bool(v int) bool {
	return v != 0
}
