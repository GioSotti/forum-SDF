package Func

func ContainsID(slice []int, id int) bool {
	for _, item := range slice {
		if item == id {
			return true
		}
	}
	return false
}
