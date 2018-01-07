
// Given a string s consists of upper/lower-case alphabets and empty space characters ' ', return the length of last word in the string.

public class LastLiteral {
    public static void main(String[] args) {
        String[] tests = {"Hello World", "", "           h", "        jhgh       hgjhg l ", " space                    ", null};

        for(String testcase: tests) {
            System.out.printf("%s -> %s\n", testcase, lengthOfLastWord(testcase));
        }
    }

    public static int lengthOfLastWord(String s) {
        if(s == null || (s.trim().equals("") == true)) {
            return 0;
        }

        String[] s_split = s.split(" ");

        return s_split[s_split.length-1].trim().length();
    }
}
