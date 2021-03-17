// Turn a string like aabccccccaaa into a2b1c5a3. If the compressed string is not smaller, return the original string. Assume only lowercase letters.
package main

import (
  "fmt"
  "strings"
  "strconv"
)

func compressString(input string) string {
  var currentCharacter rune = -1
  var currentCharacterCount = 0
  var output strings.Builder

  for _, character := range input {
    if currentCharacter == character {
      currentCharacterCount += 1
    } else {
      output.WriteString(strconv.Itoa(currentCharacterCount))
      output.WriteString(string(currentCharacter))

      currentCharacter = character
      currentCharacterCount = 0
    }
  }
  
  return output.String()
}

func main() {
  const string1 = "abc" // abc
  const string2 = "aabbccaa" // aabbccaa
  const string3 = "aabbbccaa" // a2b3c2a2
  const string4 = "aabccccccaaa" // a2b1c5a3
  const string5 = "aaaabaaaabaaaa" // a4b1a4b1a4
  
  fmt.Println(compressString(string1))
  fmt.Println(compressString(string2))
  fmt.Println(compressString(string3))
  fmt.Println(compressString(string4))
  fmt.Println(compressString(string5))
}
