// Given a string, write a function to check if it is a permutation of a palindrome.
// rustc chapter-1/questions/palindrome.rs && ./palindrome

fn is_palindrome_permutation(string: &str) -> bool {
  let mut cleaned = string.to_lowercase();
  cleaned.retain(|c| !c.is_whitespace());
  let mut odd_numbered_letter = false;

  for character in cleaned.chars() {
    let occurences = cleaned.matches(character).count();
    let is_odd = occurences % 2 != 0;

    if is_odd {
      if odd_numbered_letter {
        return false;
      } else {
        odd_numbered_letter = true;
      }
    }
  }

  return true;
}

fn main() {
  let string1 = "TtcctT"; // True
  let string2 = "TtctT"; // True
  let string3 = "Tact Coa"; // True
  let string4 = "taco catr"; // False
  println!("{}", is_palindrome_permutation(string1));
  println!("{}", is_palindrome_permutation(string2));
  println!("{}", is_palindrome_permutation(string3));
  println!("{}", is_palindrome_permutation(string4));
}
