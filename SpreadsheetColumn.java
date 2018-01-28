
/*
Given a column title as appear in an electronic spreadsheet, return its corresponding column number.

For example:

    A -> 1
    B -> 2
    C -> 3
    ...
    Z -> 26
    AA -> 27
    AB -> 28 
*/

import java.util.*;

public class SpreadsheetColumn {
    public static void main(String[] args) {
        if(args.length <= 0) {
            System.out.println("Usage: $> java SpreadsheetColumn <columnname>");
            return;
        }

        System.out.println(titleToNumber(args[0].trim()));
    }
    
    public static int titleToNumber(String s) {
        // create map structure assigning A->1, B->2,..., Z->26
        Map<Character,Integer> alphabet = new HashMap<Character,Integer>();

        char letter = 'A';
        for(int i = 1; i <= 26; i++) {
            alphabet.put(letter, i);
            letter++;
        }
        
        // treat s as a base-26 number
        int sLen = s.length();
        int pow = 0;
        int base = 26;
        int total = 0;
        for(int i = sLen-1; i >= 0; i--) {
            total += (int)(alphabet.get(s.charAt(i)) * Math.pow((double)base, (double)pow));
            pow++;
        }
        
        return total;
    }
}
