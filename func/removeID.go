package Func

func RemoveID(slice []int, id int) []int {
	index := -1
	for i, item := range slice {
		if item == id {
			index = i
			break
		}
	}
	if index >= 0 {
		slice = append(slice[:index], slice[index+1:]...)
	}
	return slice
}
