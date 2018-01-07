
/*
    Given a pattern and a string str, find if str follows the same pattern.

    Here follow means a full match, such that there is a bijection between a letter in pattern and a non-empty word in str.

    Examples:
    pattern = "abba", str = "dog cat cat dog" should return true.
    pattern = "abba", str = "dog cat cat fish" should return false.
    pattern = "aaaa", str = "dog cat cat dog" should return false.
    pattern = "abba", str = "dog dog dog dog" should return false.
    Notes:
    You may assume pattern contains only lowercase letters, and str contains lowercase letters separated by a single space.
*/

import java.util.*;

public class WordPattern {
    public static void main(String[] args) {
        if(args.length < 2) {
            System.out.println("Usage: $> java WordPattern <pattern> <input>");
        }

        String pattern = args[0].trim();
        String input = args[1].trim();

        System.out.println(wordPattern(pattern, input));
    }

    public static boolean wordPattern(String pattern, String input) {
        String[] input_array = input.split(" ");

        if(pattern.length() != input_array.length) {
            return false;
        }

        HashMap<Character,String> match = new HashMap<Character,String>();
        for(int i = 0; i < pattern.length(); i++) {
            Character patternChar = new Character(pattern.charAt(i));

            if(match.get(patternChar) == null) {
                // check if this value has been attached to any other key
                for(Map.Entry<Character, String> entry: match.entrySet()) {
                    if(Objects.equals(input_array[i], entry.getValue())) {
                        return false;
                    }
                }

                match.put(patternChar, input_array[i]);
            } else {
                if(!match.get(patternChar).equals(input_array[i])) {
                    return false;
                }
            }
        }

        return true;
    }
}
