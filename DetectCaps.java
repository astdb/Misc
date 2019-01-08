/*
Given a word, you need to judge whether the usage of capitals in it is right or not.

We define the usage of capitals in a word to be right when one of the following cases holds:

All letters in this word are capitals, like "USA".
All letters in this word are not capitals, like "leetcode".
Only the first letter in this word is capital if it has more than one letter, like "Google".
Otherwise, we define that this word doesn't use capitals in a right way.
Example 1:
Input: "USA"
Output: True
Example 2:
Input: "FlaG"
Output: False
Note: The input will be a non-empty word consisting of uppercase and lowercase latin letters.
*/

class DetectCaps {
    public static void main(String[] args) {
       String[] tests = new String[]{"USA", "FlaG", "a", "A", ""};

       for(String test: tests) {
           System.out.println("detectCapitalUse(\"" + test + "\") == " + detectCapitalUse(test));
       }
    }

    public static boolean detectCapitalUse(String word) {
        // check if all letters in this string are caps
        boolean allUpper = true;
        boolean allLower = true;

        for(int i = 0; i < word.length(); i++) {
            if(Character.isUpperCase(word.charAt(i))) {
                allLower = false;
            } else {
                allUpper = false;
            }
        }

        if(allUpper) {
            return true;
        }

        if(allLower) {
            return true;
        }

        if(word.length() > 1 && Character.isUpperCase(word.charAt(0))) {
            boolean allSmall = true;
            for(int j = 1; j < word.length(); j++) {
                if(Character.isUpperCase(word.charAt(j))) {
                    // found an uppercase in the rest
                    allSmall = false;
                }
            }

            return allSmall;
        }

        return false;
    }
}
