package main

import "fmt"


func binarySearch(arr []int, target int) bool {
    low := 0
    high := len(arr) - 1

    for low <= high {
        mid := (low + high) / 2

        if arr[mid] == target {
            return true
        } else if arr[mid] < target {
            low = mid + 1
        } else {
            high = mid - 1
        }
    }
    return false
}

func main() {
    arr := []int{2, 4, 6, 8, 10, 12, 14, 16}
    target := 10

    found := binarySearch(arr, target)

    if found {
        fmt.Println("Target found in array")
    } else {
        fmt.Println("Target not found in array")
    }
}
