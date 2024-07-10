package algorithms

func QuickSort(arr []int) {
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
    QuickSort(left)
    QuickSort(right)
    copy(arr, left)
    arr[len(left)] = pivot
    copy(arr[len(left)+1:], right)
}
