package algorithms

import (
	"reflect"
	"testing"
)

func TestInsertionSort(t *testing.T) {
	arr := []int{64, 25, 12, 22, 11}
	expected := []int{11, 12, 22, 25, 64}
    InsertionSort(arr)
	if !reflect.DeepEqual(arr, expected) {
		t.Errorf("got %v, want %v", arr, expected)
	}
}
