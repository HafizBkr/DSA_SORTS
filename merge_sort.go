package main

import "fmt"

func mergeSort(arr []int) {
    if len(arr) > 1 {
        mid := len(arr) / 2
        left := make([]int, mid)
        right := make([]int, len(arr)-mid)

        copy(left, arr[:mid])
        copy(right, arr[mid:])

        mergeSort(left)
        mergeSort(right)

        i, j, k := 0, 0, 0

        for i < len(left) && j < len(right) {
            if left[i] <= right[j] {
                arr[k] = left[i]
                i++
            } else {
                arr[k] = right[j]
                j++
            }
            k++
        }

        for i < len(left) {
            arr[k] = left[i]
            i++
            k++
        }

        for j < len(right) {
            arr[k] = right[j]
            j++
            k++
        }
    }
}

func main() {
    arr := []int{64, 25, 12, 22, 11}
    fmt.Println("Array before sorting:", arr)
    mergeSort(arr)
    fmt.Println("Array after sorting:", arr)
}
