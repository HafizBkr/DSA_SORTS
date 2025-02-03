package algorithms

func SelectionSort(arr []int) {
    n := len(arr)
    for i := 0; i < n-1; i++ {
        MinIndex := i
        for j := i + 1; j < n; j++ {
            if arr[j] < arr[MinIndex] {
                MinIndex = j
            }
        }
        arr[i], arr[MinIndex] = arr[MinIndex], arr[i]
    }
}
