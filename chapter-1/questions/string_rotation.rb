# Assume you have a method is_substring which checks if one word is a substring or another. Given two strings, write code to check if one is a rotation of another, using only one call to is_substring. E.g. "waterbottle" is a rotation of "erbottlewat"
# Note: is_substring == include? in Ruby

string1 = "waterbottle"
string2 = "erbottlewat"

def is_rotation?(string1, string2)
  return false if string1.length != string2.length
  concat = string1 + string1

  return concat.include?(string2)
end

print(is_rotation?("waterbottle", "erbottlewat")) # True
print(is_rotation?("abcde", "deabc")) # True
print(is_rotation?("abcd", "deabc")) # False
print(is_rotation?("abede", "deabc")) # False
