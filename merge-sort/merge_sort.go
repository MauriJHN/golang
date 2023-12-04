package main

import "fmt"

func merge(left []int, right []int, array []int) {
  fmt.Println("Merging: " + fmt.Sprint(left) + " and " + fmt.Sprint(right))
  var left_ix = 0
  var right_ix = 0
  var orig_ix = 0

  var left_size = len(left)
  var right_size = len(right)

  for left_ix < left_size && right_ix < right_size {
    if left[left_ix] <= right[right_ix] {
      fmt.Printf("Inserting left %d at index %d\n", left[left_ix], orig_ix)
      array[orig_ix] = left[left_ix]
      left_ix++
    } else {
      fmt.Printf("Inserting right %d at index %d\n", right[right_ix], orig_ix)
      array[orig_ix] = right[right_ix]
      right_ix++
    }
    orig_ix++
    fmt.Printf("Left array is now %v\n", left)
  }

  for left_ix < left_size {
    array[orig_ix] = left[left_ix]
    fmt.Printf("Inserting left %d at index %d\n", left[left_ix], orig_ix)
    left_ix++
    orig_ix++
  }

  for right_ix < right_size {
    array[orig_ix] = right[right_ix]
    fmt.Printf("Inserting right %d at index %d\n", right[right_ix], orig_ix)
    right_ix++
    orig_ix++
  }
}

func merge_sort(array []int) {
  // Base case. A list of zero or one elements is sorted, by definition.
  if len(array) < 2 {
    return
  }

  // Recursive case. First, divide the list into equal-sized sublists
  var mid = len(array) / 2
  // Copy left half of array into new array
  var left = make([]int, mid)
  copy(left, array[:mid])
  // Copy right half of array into new array
  var right = make([]int, len(array)-mid)
  copy(right, array[mid:])

  // Recursively sort both sublists
  merge_sort(left)
  merge_sort(right)
  merge(left, right, array)
}
