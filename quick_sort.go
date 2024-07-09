package main

import "fmt"

func quickSort(arr []int) {
    if len(arr) < 2 {
        return
    }
    pivot := arr[0]
    var left, right []int
    for _, v := range arr[1:] {
        if v <= pivot {
            left = append(left, v)
        } else {
            right = append(right, v)
        }
    }
    quickSort(left)
    quickSort(right)
    copy(arr, left)
    arr[len(left)] = pivot
    copy(arr[len(left)+1:], right)
}

func main() {
    arr := []int{64, 25, 12, 22, 11}
    fmt.Println("Array before sorting:", arr)
    quickSort(arr)
    fmt.Println("Array after sorting:", arr)
}
