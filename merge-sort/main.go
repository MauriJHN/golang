package main

import "fmt"

func main() {
  array1 := []int{1,3021,15,7,910};
  fmt.Println("Unsorted array: " + fmt.Sprint(array1) + "\n");
  merge_sort(array1);
  fmt.Println("Sorted array: " + fmt.Sprint(array1) + "\n");
}
