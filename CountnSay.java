
/*
    The count-and-say sequence is the sequence of integers with the first five terms as following:

    1.     1
    2.     11
    3.     21
    4.     1211
    5.     111221
    1 is read off as "one 1" or 11.
    11 is read off as "two 1s" or 21.
    21 is read off as "one 2, then one 1" or 1211.
    Given an integer n, generate the nth term of the count-and-say sequence.

    Note: Each term of the sequence of integers will be represented as a string.
*/

import java.util.*;
import java.math.*;

public class CountnSay {
    public static void main(String[] args) {
        if(args.length != 1) {
            System.out.println("Usage: $> java CountnSay <input>");
            return;
        }

        int countTo = 0;
        try {
             countTo = Integer.parseInt(args[0]);
        } catch(NumberFormatException e) {
            System.out.println("Program input need to be a valid numeric value.");
            return;
        }

        if(countTo <= 0) {
            System.out.println("Program input need to be a positive numeric value.");
            return;
        }

        for(int i = 1; i < countTo + 1; i++) {
            System.out.println(countAndSay(i));
        }
    }

    public static String countAndSay(int n) {
        // int base = 10;
        BigInteger base = new BigInteger("10");
        BigInteger i = new BigInteger("1");  // starting term

        if(n == 1) {
            return "1";
        }

        for (int termCount = 1; termCount < n; termCount++) {
            BigInteger rem = i.remainder(base);
            i = i.divide(base);

            // keep track of digit groups
            String term = "";
            int groupCount = 1;
            BigInteger digit = rem;

            while(i.compareTo(BigInteger.valueOf(0)) == 1) {
                rem = i.remainder(base);
                i = i.divide(base);

                if(rem.compareTo(digit) != 0) {
                    // new digit group starting
                    term = groupCount + "" + digit.toString()  + term;

                    groupCount = 1;
                    digit = rem;

                } else {
                    groupCount++;
                }
            }

            // close off the final digit group
            term = groupCount + "" + digit.toString()  + term;

            if(termCount == (n - 1)) {
                return term;
            }

            i = new BigInteger(term);
        }        

        return "";
    }
}
