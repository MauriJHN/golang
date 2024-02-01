package main

import (
  "fmt"
  "math"
)

func getUserInput() string {
  var input string
  fmt.Scanln(&input)
  return input
}

func getCaseInput() (float64, float64) {
  var collection1, collection2 float64
  fmt.Scanln(&collection1, &collection2)
  return collection1, collection2
}

func getCaseCount() int {
  var caseCount int
  fmt.Scanln(&caseCount)
  return caseCount
}

func checkDenom(maxDenom float64, collection1 float64, collection2 float64) bool {
  if (math.Mod(collection1, maxDenom) == 0 && math.Mod(collection2, maxDenom) == 0) {
    return true
  }
  return false
}

func getMaximumDenominator(collection1 float64, collection2 float64) float64 {
  var maxDenom float64
  maxDenom = math.Mod(math.Max(collection1, collection2), math.Min(collection1, collection2))
  if (maxDenom == 0) {
    return math.Min(collection1, collection2)
  }

  // the denom will only go down, so it makes sense to use it for the for loop
  for (maxDenom > 0) {
    if (checkDenom(maxDenom, collection1, collection2)) {
      return maxDenom
    }
    maxDenom = math.Min(math.Mod(collection1, maxDenom), math.Mod(collection2, maxDenom))
  }

  return maxDenom
}

func processCase() float64 {
  collection1, collection2 := getCaseInput()
  maximumDenom := getMaximumDenominator(collection1, collection2)
  return maximumDenom
}

func main() {
  caseCount := getCaseCount()
  maxDenoms := make([]float64, caseCount)

  for i := 0; i < caseCount; i++ {
    maxDenom := processCase()
    maxDenoms[i] = maxDenom
  }

  for i := 0; i < caseCount; i++ {
    fmt.Println(maxDenoms[i])
  }
}
