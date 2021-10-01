package gools

// Find ...
/*
	- Allows you to find an element inside a slice of elements based on a Find Function that is passed as a second
	paramater.

	* elements -> Array of any data type
	* findFn -> Function used for filtering. Should return a boolean upon return

	=> nil if not found
	=> Element found of DType
*/
func Find[DType any](elements []DType, findFn(element DType) bool) *DType {
	for _, each := range elements {
		if findFn(each) {
			return each
		}
	}

	return nil
}