package main

import "testing"

func TestMergeSort1Element(t *testing.T) {
	a := [1]int{1}

	mergeSort(a[0:len(a)])

	exoected := [1]int{1}

	if a != exoected {
		t.Errorf("%v != %v", a, exoected)
	}
}

func TestMergeSort2SortedElements(t *testing.T) {
	a := [2]int{1, 2}

	mergeSort(a[0:len(a)])

	exoected := [2]int{1, 2}

	if a != exoected {
		t.Errorf("%v != %v", a, exoected)
	}
}

func TestMergeSort2UnSortedElements(t *testing.T) {
	a := [2]int{2, 1}

	mergeSort(a[0:len(a)])

	exoected := [2]int{1, 2}

	if a != exoected {
		t.Errorf("%v != %v", a, exoected)
	}
}

func TestMergeSortLongArray(t *testing.T) {
	a := [8]int{5, 2, 4, 7, 1, 3, 2, 6}

	mergeSort(a[0:len(a)])

	exoected := [8]int{1, 2, 2, 3, 4, 5, 6, 7}

	if a != exoected {
		t.Errorf("%v != %v", a, exoected)
	}
}
