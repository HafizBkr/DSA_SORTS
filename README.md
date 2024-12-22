Sorting Algorithms in Go

This repository contains implementations of various sorting algorithms in the Go programming language, along with tests to verify their correctness.
Algorithms Implemented
1. Insertion Sort:

Insertion sort is a simple sorting algorithm that builds the final sorted array one item at a time.

Implementation:

    File: insertsort/insertion_sort.go

Test File: 

    File: insertsort/insert_sort_test.go

Test Usage:


    go test ./insertsort -run TestInsertionSort

2. Selection Sort

Selection sort sorts an array by repeatedly finding the minimum element from the unsorted part and putting it at the beginning.

Implementation:

    File: selectionsort/selection_sort.go

Test File:

    File: selectionsort/selection_sort_test.go

Test Usage:



    go test ./selectionsort -run TestSelectionSort

3. Bubble Sort

Bubble sort repeatedly steps through the list, compares adjacent elements, and swaps them if they are in the wrong order.

Implementation:

    File: bubblesort/bubble_sort.go

Test File:

    File: bubblesort/bubble_sort_test.go

Test Usage:


    go test ./bubblesort -run TestBubbleSort

4. Merge Sort

Merge sort is a divide and conquer algorithm that divides the input array into two halves, recursively sorts them, and then merges the sorted halves.

Implementation:

    File: mergesort/merge_sort.go

Test File:

    File: mergesort/merge_sort_test.go

Test Usage:

    go test ./mergesort -run TestMergeSort

5. Quick Sort

Quick sort is a divide and conquer algorithm that picks an element as a pivot and partitions the array around the pivot.

Implementation:

    File: quicksort/quick_sort.go

Test File:

    File: quicksort/quick_sort_test.go

Test Usage:

    go test ./quicksort -run TestQuickSort

Running the Tests

Each sorting algorithm is implemented in a separate Go file, and corresponding tests are provided to verify their correctness. To run the tests for all the algorithms, navigate to the root directory of the project and use the following command::


    go test ./...

Contributing

Contributions and improvements are welcome! If you have any suggestions or find issues, please feel free to open an issue or submit a pull request..
