// Turn a string like aabccccccaaa into a2b1c5a3. If the compressed string is not smaller, return the original string. Assume only lowercase letters.
package main

import (
  "fmt"
  "strings"
  "strconv"
)

func compressString(input string) string {
  var currentCharacterCount = 0
  var output strings.Builder
  var currentCharacter = []rune(input)[0]
  output.WriteString(string(currentCharacter))

  for index, character := range input {
    if currentCharacter == character {
      currentCharacterCount += 1
    } else {
      output.WriteString(strconv.Itoa(currentCharacterCount))
      output.WriteString(string(character))

      currentCharacter = character
      currentCharacterCount = 1
    }

    if index == (len(input) - 1) {
      output.WriteString(strconv.Itoa(currentCharacterCount))
    }
  }
  
  var outputString = output.String()
  if len(outputString) < len(input) {
    return outputString
  } else {
    return input
  }
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
