package gools_test

import (
	"strings"
	"testing"
)

func TestFind_Positive(t *testing.T) {
	inputSlice := []string{"This is an example", "This is a test", "Type of Error", "Closing Socket"}

	findFunction := func(value string) bool {
		return strings.Contains(value, "This")
	}

	expectedOutput := "This is an example"

	foundItem := Find(inputSlice, findFunction)

	if foundItem == nil {
		t.Fatalf("Output did not match the expected output :: ERR")
	}

	if *foundItem != expectedOutput {
		t.Fatalf("Output did not match the expected output :: ERR")
	}
}

func TestFind_Negative(t *testing.T) {
	inputSlice := []string{"this is an example", "this is a test", "Type of Error", "Closing Socket"}

	findFunction := func(value string) bool {
		return strings.Contains(value, "This")
	}

	foundItem := Find(inputSlice, findFunction)

	if foundItem != nil {
		t.Fatalf("Output did not match the expected output :: ERR")
	}
}
