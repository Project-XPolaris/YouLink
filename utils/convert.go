package utils

func MapConvert(from []interface{}) []map[string]interface{} {
	result := make([]map[string]interface{}, 0)
	for _, item := range from {
		result = append(result, item.(map[string]interface{}))
	}
	return result
}
