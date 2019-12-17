package sortutil

// 打乱顺序
func Jumble(array []interface{}) []interface{} {
	tmpMap := make(map[int]interface{}, len(array))
	for i, ele := range array {
		tmpMap[i] = ele
	}

	result := make([]interface{}, len(array))
	k := 0
	for _, value := range tmpMap {
		result[k] = value
		k++
	}
	return result
}
