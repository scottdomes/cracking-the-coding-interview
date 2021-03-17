// There are three types of edits that can be performed on strings: insert a character, remove a character, or replace a character. Given two strings, write a function to check if they are one or zero edits away.

func is_one_away(string1: String, string2: String) -> Bool {
  return true
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