// Write a method that replaces all spaces in a string with '%20', assuming that the string has sufficient space at the end for the new characters.
// java chapter-1/questions/urlify.java

public class Urlify {
  public static void urlify(String string) {
    char[] characterArray = string.toCharArray();
    String replacedString = "";

    for (int i = 0; i < characterArray.length; i++) {
      if (replacedString.length() >= characterArray.length - 2) {
        break;
      }
  
      if (characterArray[i] == ' ') {
        replacedString += "%20";
      } else {
        replacedString += characterArray[i];
      }
    }

    System.out.println(replacedString);
  }

  public static void main(String[] args) {
    String string1 = "Scott Domes   ";
    String string2 = "Mr Scott Domes      ";

    urlify(string1);
    urlify(string2);
  }
}