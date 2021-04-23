package arrutil

func Diff(base, exclude []string) (result []string) {
	excludeMap := make(map[string]bool)
	for _, s := range exclude {
		excludeMap[s] = true
	}
	for _, s := range base {
		if !excludeMap[s] {
			result = append(result, s)
		}
	}
	return result
}

func Unique(ss []string) (result []string) {
	smap := make(map[string]bool)
	for _, s := range ss {
		smap[s] = true
	}
	for s := range smap {
		result = append(result, s)
	}
	return result
}

func FindString(str string, array []string) int {
	for index, s := range array {
		if str == s {
			return index
		}
	}
	return -1
}

func StringIn(str string, array []string) bool {
	return FindString(str, array) > -1
}

func StringRomove(str string, array []string) ([]string, bool) {
	index := FindString(str, array)
	if index == -1 {
		return nil, false
	}

	result := []string{}
	if index > 0 {
		result = append(result, array[:index]...)
	}
	if index < len(array)-1 {
		result = append(result, array[(index+1):]...)
	}
	return result, true
}
