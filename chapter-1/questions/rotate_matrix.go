// Given an image represented by an NxN matrix, write a method to rotate the image by 90 degrees.
package main

import (
  "fmt"
  "math/rand"
)

func generateMatrix(n int) [][]int {
  matrix := make([][]int, n) 

  for i:=0; i<n; i++ {
    matrix[i] = make([]int, n) 
  }

  for i:=0; i<n; i++ {
      for j:=0; j<n; j++ {
          matrix[i][j] = rand.Intn(10)
      }
  }

  return matrix
}

func rotateMatrix(matrix [][]int) [][]int {
  n := len(matrix)

  newMatrix := generateMatrix(n)

  for row:=0; row<n; row++ {
    for column:=0; column<n; column++ {
      newMatrix[row][column] = matrix[n - column - 1][row]
    }
  }

  return newMatrix
}

func main() {
  matrix := generateMatrix(3)
  
  fmt.Println(matrix[0])
  fmt.Println(matrix[1])
  fmt.Println(matrix[2])

  rotatedMatrix := rotateMatrix(matrix)

  fmt.Println("Rotated:")
  fmt.Println(rotatedMatrix[0])
  fmt.Println(rotatedMatrix[1])
  fmt.Println(rotatedMatrix[2])

  matrix2 := generateMatrix(5)

  fmt.Println(matrix2[0])
  fmt.Println(matrix2[1])
  fmt.Println(matrix2[2])
  fmt.Println(matrix2[3])
  fmt.Println(matrix2[4])

  rotatedMatrix2 := rotateMatrix(matrix2)

  fmt.Println("Rotated:")
  fmt.Println(rotatedMatrix2[0])
  fmt.Println(rotatedMatrix2[1])
  fmt.Println(rotatedMatrix2[2])
  fmt.Println(rotatedMatrix2[3])
  fmt.Println(rotatedMatrix2[4])

}
