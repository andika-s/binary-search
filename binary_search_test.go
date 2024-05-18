package binary_search

import (
	"testing"
	"time"
)

type Data struct {
	name     string
	array    []int
	search   int
	expected int
}

func NewData() []Data {
	return []Data{
		{"should pass when search element value: 9 at index: 8", []int{1, 2, 3, 4, 5, 6, 7, 8, 9}, 9, 8},
		{"should pass when search element value: 188 at index: 4", []int{10, 12, 45, 98, 188, 276}, 188, 4},
	}
}

func TestBinarySearch(t *testing.T) {
	ticker := time.NewTicker(1 * time.Second).C

	for index, value := range NewData() {
		t.Run(value.name, func(t *testing.T) {
			actual := BinarySearch(value.array, value.search)
			if actual != value.expected {
				t.Errorf("the actual value: %v not equals with expected value: %v at index: %v", actual, value.search, index)
			}
			<-ticker
		})
	}
}

func TestRecursiveBinarySearch(t *testing.T) {
	ticker := time.NewTicker(1 * time.Second).C

	for index, value := range NewData() {
		t.Run(value.name, func(t *testing.T) {
			actual := RecursiveBinarySearch(value.array, value.search, 0, len(value.array))
			if actual != value.expected {
				t.Errorf("the actual value: %v not equals with expected value: %v at index: %v", actual, value.search, index)
			}
			<-ticker
		})
	}
}
