package gools_test

import "testing"

func TestChunks(t *testing.T) {
	arr := []int{1, 2, 3, 4, 5, 6, 7, 8, 9}

	expectedOutput := [][]int{{1, 2}, {3, 4}, {5, 6}, {7, 8}, {9}}
	output := Chunk(arr, 2)

	if len(output) != len(expectedOutput) {
		t.Fatalf("Chunk length is not valid [ERR1]")
	}

	for i := range output {
		outIndex := output[i]

		outExpectedIndex := expectedOutput[i]

		if len(outIndex) != len(outExpectedIndex) {
			t.Fatalf("Chunk item length is not valid [ERR2]")
		}

		for x := range outIndex {
			if outIndex[x] != outExpectedIndex[x] {
				t.Fatalf("Chunk item.inner length is not valid [ERR3]")
			}
		}
	}

}
