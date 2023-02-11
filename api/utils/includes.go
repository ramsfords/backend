package utils

func IncludesService(arr []int32, value int32) bool {
	includes := false
	for _, j := range arr {
		if j == value {
			includes = true
			break
		}
	}
	return includes
}
