package binary_search

/*
BinarySearch performs a binary search on a sorted array.

@param sortedArray is a slice of integers sorted in ascending order.
@param search is the integer value we are searching for.
@return the index of the search value if found; otherwise, -1.
*/
func BinarySearch(sortedArray []int, search int) int {
	var (
		low  = 0
		high = len(sortedArray) - 1
	)

	// Iterate until low value exceeds high value
	for low <= high {
		// Calculate the middle index
		var middle int = (low + high) / 2

		// If the middle element equals the search value,
		// return the middle index
		if sortedArray[middle] == search {
			return middle
		}

		// If the middle element is less than the search value,
		// move the low index to the right of middle
		if sortedArray[middle] < search {
			low = middle + 1
		} else {
			// If the middle element is greater than the search value,
			// move the high index to the left of middle
			high = middle - 1
		}
	}
	// If the search value is not found, return -1
	return -1
}

/*
RecursiveBinarySearch performs a binary search on a sorted array recursively.

@param sortedArray is a slice of integers sorted in ascending order.
@param search is the integer value we are searching for.
@param low is the lower index of the subarray to search.
@param high is the higher index of the subarray to search.
@return the index of the search value if found; otherwise, -1.
*/
func RecursiveBinarySearch(sortedArray []int, search, low, high int) int {
	// Base case: If low value exceeds high value, return -1 indicating search value not found
	if low > high {
		return -1
	}

	// Calculate the middle index
	var middle int = (low + high) / 2

	// If the middle element equals the search value, return the middle index
	if sortedArray[middle] == search {
		return middle
	}

	// If the middle element is less than the search value, search the right half recursively
	if sortedArray[middle] < search {
		return RecursiveBinarySearch(sortedArray, search, middle+1, high)
	} else {
		// If the middle element is greater than the search value, search the left half recursively
		return RecursiveBinarySearch(sortedArray, search, low, middle-1)
	}
}
