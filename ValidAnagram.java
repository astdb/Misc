/*
Given two strings s and t , write a function to determine if t is an anagram of s.

Example 1:

Input: s = "anagram", t = "nagaram"
Output: true
Example 2:

Input: s = "rat", t = "car"
Output: false
Note:
You may assume the string contains only lowercase alphabets.

Follow up:
What if the inputs contain unicode characters? How would you adapt your solution to such case?
*/

import java.util.*;

class ValidAnagram {
    public static void main(String[] args) {
        String[][] tests = new String[][]{{"anagram", "nagaram"}};

        for (int i = 0; i < tests.length; i++) {
            System.out.println(isAnagram(tests[i][0], tests[i][1]));
        }
    }

    public static boolean isAnagram(String s, String t) {
        // both strings must be of same length to be anagrams
        if(s.length() != t.length()) {
            return false;
        }

        // convert both strings to char arrays and sort
        char[] s_ch = s.toCharArray();
        char[] t_ch = t.toCharArray();

        Arrays.sort(s_ch);
        Arrays.sort(t_ch);

        // compare sorted char array version char by char to ensure they're the same
        for(int i  = 0; i < s_ch.length; i++) {
            if(s_ch[i] != t_ch[i]) {
                return false;
            }
        }

        return true;
    }
}
