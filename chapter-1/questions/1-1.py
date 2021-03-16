# Implement an algorithm to determine if a string has all unique characters.

string1 = 'Scott' # -> False
string2 = 'Aaaa' # -> False
string3 = 'Asdfghkl' # -> True
string4 = 'Aa' # -> False
string5 = 'AbCdEfG' # -> True

def has_all_unique_characters(string):
  character_hash = dict()
  for character in string:
    if character.lower() in character_hash:
      return False
    else:
      character_hash[character.lower()] = True
  
  return True
    

print(has_all_unique_characters(string1))
print(has_all_unique_characters(string2))
print(has_all_unique_characters(string3))
print(has_all_unique_characters(string4))
print(has_all_unique_characters(string5))

# What if you cannot use additional data structures?

def has_all_unique_characters2(string):
  lowercase = string.lower()
  for character in lowercase:
    new_string = lowercase.replace(character, '')
    if len(new_string) != len(string) - 1:
      return False
  
  return True
    
print(has_all_unique_characters2(string1))
print(has_all_unique_characters2(string2))
print(has_all_unique_characters2(string3))
print(has_all_unique_characters2(string4))
print(has_all_unique_characters2(string5))