package lab2

func Filter(arr []string, cond func(string) bool) []string {
	var result []string
	for i := range arr {
		if cond(arr[i]) {
			result = append(result, arr[i])
		}
	}
	return result
}
