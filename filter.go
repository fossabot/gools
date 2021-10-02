package gools


// Filter ...
/*
	Iterates over elements of collection, returning an array of all elements predicate returns truthy for. 
	The predicate is invoked with three arguments: (value, index|key, collection).
*/
func Filter[DType any](collection []DType, predicate func(element DType) bool) []DType {
	indexedItems := make([]Dtype, 0)

	for _, each := range collection {
		if predicate(each) {
			indexedItems = append(indexedItems, each)
		}
	}

	return indexedItems
}