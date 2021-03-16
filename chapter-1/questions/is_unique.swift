// Implement an algorithm to determine if a string has all unique characters.

let string1 = "Scott" // -> False
let string2 = "Aaaa" // -> False
let string3 = "Asdfghkl" // -> True
let string4 = "Aa" // -> False
let string5 = "AbCdEfG" // -> True

func has_all_unique_characters(string: String) -> Bool {
  var character_hash: [String: Bool] = [:]
  for character in string {
    if character_hash[character.lowercased()] != nil {
      return false
    } else {
      character_hash[character.lowercased()] = true
    }
  }
  return true
}

print(has_all_unique_characters(string: string1))
print(has_all_unique_characters(string: string2))
print(has_all_unique_characters(string: string3))
print(has_all_unique_characters(string: string4))
print(has_all_unique_characters(string: string5))


// What if you cannot use additional data structures?

// Note: for some reason this method wasn't working in replit
func replacingOccurrences(string: String, of: Character, with: String) -> String {
  return string.map { $0 == of ? with : String($0) }.joined()
}

func has_all_unique_characters2(string: String) -> Bool {
  let lowercase = string.lowercased()
  for character in lowercase {
    let new_string = replacingOccurrences(string: lowercase, of: character, with: "")

    if new_string.count != string.count - 1 {
      return false
    }
  }

  return true
}
    
print(has_all_unique_characters2(string: string1))
print(has_all_unique_characters2(string: string2))
print(has_all_unique_characters2(string: string3))
print(has_all_unique_characters2(string: string4))
print(has_all_unique_characters2(string: string5))
