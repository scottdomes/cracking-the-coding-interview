// Given two strings, write a method to decide if one is a permutation of another
package main

import (
  "fmt"
  "sort"
  "strings"
)

func isPermutation(stringA, stringB string) bool {
  if len(stringA) != len(stringB) {
    return false
  }

  var arrayA = strings.Split(stringA, "")
  var arrayB = strings.Split(stringB, "")
  sort.Strings(arrayA)
  sort.Strings(arrayB)

  var sortedStringA = strings.Join(arrayA, "")
  var sortedStringB = strings.Join(arrayB, "")

  if sortedStringA == sortedStringB {
    return true
  }

  return false
}

func main() {
  const string1 = "abc"
  const string2 = "bca"
  const string3 = "bcb"
  
  fmt.Println(isPermutation(string1, string2)) // True
  fmt.Println(isPermutation(string2, string1)) // True 
  fmt.Println(isPermutation(string1, string3)) // False
  fmt.Println(isPermutation(string3, string2)) // False
}
