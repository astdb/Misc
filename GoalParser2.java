/*
You own a Goal Parser that can interpret a string command. The command consists of an alphabet of "G", "()" and/or "(al)" in some order. The Goal Parser will interpret "G" as the string "G", "()" as the string "o", and "(al)" as the string "al". The interpreted strings are then concatenated in the original order.

Given the string command, return the Goal Parser's interpretation of command.

 

Example 1:

Input: command = "G()(al)"
Output: "Goal"
Explanation: The Goal Parser interprets the command as follows:
G -> G
() -> o
(al) -> al
The final concatenated result is "Goal".
Example 2:

Input: command = "G()()()()(al)"
Output: "Gooooal"
Example 3:

Input: command = "(al)G(al)()()G"
Output: "alGalooG"
 

Constraints:

1 <= command.length <= 100
command consists of "G", "()", and/or "(al)" in some order.

*/

public class GoalParser2 {
  public static void main(String args[]) {
    String[] tests = {"", "G", "()", "(al)", "G()(al)","G()()()()(al)", "(al)G(al)()()G", "Goal"};

    for(String test: tests) {
      System.out.printf("interpret(%s) = %s\n", test, interpret(test));
    }
  }

  private static String interpret(String command) {
    int cmdLen = command.length();
    StringBuilder result = new StringBuilder();

    for(int i = 0; i < cmdLen; i++) {
      char c = command.charAt(i);

      if(c == 'G') {
        result.append("G");

      } else if(i+2 <= cmdLen &&  command.substring(i, i+2) == "()") {
        result.append("o");
        i++;

      } else if(i+4 <= cmdLen && command.substring(i, i+4) == "(al)") {
        result.append("al");
        i += 3;
      }
    }    

    return result.toString();
  }
}

