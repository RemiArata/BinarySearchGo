package bs

func binarySearch(data []int, findVal int) (int, bool) {
	var notFound bool = true
	var lower int = 0
	var upper int = len(data)
	var mid int

	for notFound {
		if upper-lower == 1 {
			break
		}
		mid = (upper + lower) / 2
		if data[mid] == findVal {
			notFound = false
		} else if data[mid] < findVal {
			lower = mid
		} else {
			upper = mid
		}
	}
	return mid, data[mid] == findVal
}
