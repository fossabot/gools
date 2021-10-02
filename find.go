package gools

// Find ...
/*
	Iterates over elements of collection, returning the first element predicate returns truthy for. 
	The predicate is invoked with three arguments: (value, index|key, collection).
*/
func Find[DType any](collection []DType, predicate(element DType) bool) *DType {
	for _, each := range collection {
		if predicate(each) {
			return each
		}
	}

	return nil
}