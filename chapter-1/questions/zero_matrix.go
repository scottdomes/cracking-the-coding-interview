// If an element in an MxN matrix is 0, its entire row and column are set to zero
package main

import (
  "fmt"
  "math/rand"
)

func generateMatrix(m, n int) [][]int {
  matrix := make([][]int, m) 

  for i:=0; i<m; i++ {
    matrix[i] = make([]int, n) 
  }

  for i:=0; i<m; i++ {
      for j:=0; j<n; j++ {
          matrix[i][j] = rand.Intn(10)
      }
  }

  return matrix
}

func zeroRow(matrix [][]int, row int) {
  n := len(matrix[0])

  for column:=0; column<n; column++ {
    matrix[row][column] = 0
  }
}

func zeroColumn(matrix [][]int, column int) {
  m := len(matrix)

  for row:=0; row<m; row++ {
    matrix[row][column] = 0
  }
}

func zeroMatrix(matrix [][]int) [][]int {
  m := len(matrix)
  n := len(matrix[0])

  newMatrix := generateMatrix(m, n)

  for row:=0; row<m; row++ {
    for column:=0; column<n; column++ {
      if matrix[row][column] == 0 {
        zeroRow(newMatrix, row)
        zeroColumn(newMatrix, column)
      }
    }
  }

  return newMatrix
}

func main() {
  matrix := generateMatrix(3, 4)
  
  fmt.Println(matrix[0])
  fmt.Println(matrix[1])
  fmt.Println(matrix[2])

  zeroedMatrix := zeroMatrix(matrix)

  fmt.Println("Zeroed:")
  fmt.Println(zeroedMatrix[0])
  fmt.Println(zeroedMatrix[1])
  fmt.Println(zeroedMatrix[2])
}
