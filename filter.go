package gools


// Filter ...
/*
	- Allows you to filter a slice of elements based on a Filter Function that is passed as a second
	paramater.

	* elements -> Array of any data type
	* filterFunction -> Function used for filtering. Should return a boolean upon return
*/
func Filter[DType any](elements []DType, filterFn func(element DType) bool) []DType {
	indexedItems := make([]Dtype, 0)

	for _, each := range elements {
		if filterFn(each) {
			indexedItems = append(indexedItems, each)
		}
	}

	return indexedItems
}