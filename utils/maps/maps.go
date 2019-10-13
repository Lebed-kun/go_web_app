package maps

func Copy(src map[string]interface{}) map[string]interface{} {
	newMap := make(map[string]interface{})
	for key, value := range src {
		newMap[key] = value
	}
	return newMap
}
