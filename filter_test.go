package gools

import (
	"strings"
	"testing"
)

func TestFilter_Simple(t *testing.T) {
	inputSlice := []string{"This is an example", "This is a test", "Type of Error", "Closing Socket"}

	filterFunction := func(value string) bool {
		return strings.Contains(value, "This")
	}

	expectedOutput := []string{"This is an example", "This is a test"}

	filteredItems := Filter(inputSlice, filterFunction)

	if len(filteredItems) != len(expectedOutput) {
		t.Fatalf("Output did not match the expected output :: ERR1")
	}

	for i := range filteredItems {
		if filteredItems[i] != expectedOutput[i] {
			t.Fatalf("Output did not match the expected output :: ERR2")
		}
	}
}
