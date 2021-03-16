fn is_palindrome_permutation(string: &str) -> bool {
  let lowercased = string.to_lowercase();
  let chars = lowercased.chars();
  chars.filter(|c| !c.is_whitespace());
  // let filtered = cleaned_string.filter(|c| !c.is_whitespace());
  let mut odd_numbered_letter = false;
  for character in chars {
    let occurences = string.matches(character).count();
    let is_odd = occurences % 2 != 0;
    // println!("{}", is_odd);
    // println!("{}", character);

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
