// Given an image represented by an NxN matrix, write a method to rotate the image by 90 degrees.
package main

import (
  "fmt"
  "math/rand"
)

func generateMatrix() [3][3]int {
  matrix := [3][3]int{}
  for i:=0; i<3; i++ {
      for j:=0; j<3; j++ {
          matrix[i][j] = rand.Intn(10)
      }
  }

  return matrix
}

func rotateMatrix(matrix [3][3]int) [3][3]int {
  return matrix
}

func main() {
  matrix := generateMatrix()
  
  fmt.Println(matrix[0])
  fmt.Println(matrix[1])
  fmt.Println(matrix[2])

  rotatedMatrix := rotateMatrix(matrix)

  fmt.Println("Rotated:")
  fmt.Println(rotatedMatrix[0])
  fmt.Println(rotatedMatrix[1])
  fmt.Println(rotatedMatrix[2])
}
