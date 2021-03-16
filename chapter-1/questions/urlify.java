public class Urlify {
  public static void urlify(String string) {
    char[] characterArray = string.toCharArray();
    char[] replacedCharacterArray = new char[characterArray.length];

    for (int i = 0; i < characterArray.length; i++) {
      if (characterArray[i] == ' ') {
        replacedCharacterArray[i] = "%20";
      } else {
        replacedCharacterArray[i] = characterArray[i];
      }
    }

    String result = new String(replacedCharacterArray);

    System.out.println(result);
  }

  public static void main(String[] args) {
    String string1 = "Scott Domes";
    String string2 = "Mr Scott Domes";
    String string3 = " Mr Scott Domes      ";

    urlify(string1);
    urlify(string2);
    urlify(string3);
  }
}