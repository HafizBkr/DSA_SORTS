package main

import "fmt"

func bubbleSort(arr []int) {
    n := len(arr)
    for i := 0; i < n-1; i++ {
        swapped := false
        for j := 0; j < n-i-1; j++ {
            if arr[j] > arr[j+1] {
                arr[j], arr[j+1] = arr[j+1], arr[j]
                swapped = true
            }
        }
        if !swapped {
            break
        }
    }
}

func main() {
    arr := []int{64, 25, 12, 22, 11}
    fmt.Println("Array before sorting:", arr)
    bubbleSort(arr)
    fmt.Println("Array after sorting:", arr)
}
