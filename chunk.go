package gools

// Chunks ...
/*
	Creates an array of elements split into groups the length of size. 
	If array can't be split evenly, the final chunk will be the remaining elements.
*/
func Chunk[DType any](array []DType, size int) [][]DType {
	chunks := [][]DType{}

	var index = 0
	arr := []DType{}
	for _, each := range array {
		if index < size {
			arr = append(arr, each)
		} else {
			chunks = append(chunks, arr)
			arr = []DType{}
		}

		index += 1
	}

	if len(arr) > 0 {
		chunks = append(chunks, arr)
	}

	return chunks
}