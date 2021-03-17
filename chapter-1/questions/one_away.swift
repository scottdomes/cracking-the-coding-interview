// There are three types of edits that can be performed on strings: insert a character, remove a character, or replace a character. Given two strings, write a function to check if they are one or zero edits away.

func is_one_insert_away(string1: String, string2: String) -> Bool {
  return string2.count - 1 == string1.count
}

func is_one_delete_away(string1: String, string2: String) -> Bool {
  return string2.count + 1 == string1.count
}

func count_different_characters(string1: String, string2: String) -> Int {
  var different_characters = 0
  for character in string1 {
    if !string2.contains(character) {
      different_characters += 1
    }
  }

  return different_characters
}

func is_one_replace_away(string1: String, string2: String) -> Bool {
  if string1.count != string2.count {
    return false
  }

  return count_different_characters(string1: string1, string2: string2) == 1
}

func is_one_away(string1: String, string2: String) -> Bool {
  return string1 == string2 || is_one_insert_away(string1: string1, string2: string2) || is_one_delete_away(string1: string1, string2: string2) || is_one_replace_away(string1: string1, string2: string2)
}

let example_string = "pale"
let string1 = "pales"
let string2 = "ple"
let string3 = "bale"
let string4 = "bake"

print(is_one_away(string1: example_string, string2: example_string)) // true
print(is_one_away(string1: example_string, string2: string1)) // true
print(is_one_away(string1: example_string, string2: string2)) // true
print(is_one_away(string1: example_string, string2: string3)) // true
print(is_one_away(string1: example_string, string2: string4)) // false