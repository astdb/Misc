
/*
   A palindromic number reads the same both ways. The largest palindrome made from the product of two 2-digit numbers is 9009 = 91 × 99.
   Find the largest palindrome made from the product of two 3-digit numbers.
*/

import java.util.*;

public class LargestPalindrome {
    public static void main(String[] args) {
        // largest number made out of two three-digit numbers' product
        int topLimit = 998001;      // 999 * 999

        // smallest number made out of two three-digit numbers' product
        int lowLimit = 10000;       // 100 * 100

        // look for numbers from topLimit to lowLimit for palindrome numbers, and each of those for three digit products
        for(int i = topLimit; i >= lowLimit; i--) {
            // check if i is a palindrome number
            if(isPalindrome(i)) {
                // get factors
                int[] tgFactors = threeDigitFactors(i);

                boolean j = false;

                if(tgFactors[0] != -1 && tgFactors[1] != -1) {
                    j = true;
                }
            
                if(j == true) {
                    System.out.printf("Palindromic: %d, Factors: %d, %d\n", i, tgFactors[0], tgFactors[1]);
                    return;
                }
            }
	    }
    }

    // checks if i is a 'palindromic' number
    public static boolean isPalindrome(int i) {
        Vector<Integer> digits = new Vector<Integer>(new Integer(4), new Integer(4));

        int rem = i % 10;
        i = i / 10;
        digits.addElement(new Integer(rem));

        while (i > 0) {
            rem = i % 10;
            i = i / 10;
            digits.addElement(new Integer(rem));
            // digits = append(digits, rem)
        }

        int start = 0;
        int end = digits.size() -1;
        while(start < end) {
            // if digits[start] != digits[end] {
            if(digits.elementAt(start).intValue() != digits.elementAt(end).intValue()) {
                return false;
            }
            start++;
            end--;
        }

        return true;
    }

    // find a pair of three digit factors for i
    public static int[] threeDigitFactors(int i) {
        int k = 100;

        while (k <= 999) {
            if ((i % k) == 0) {
                int n = i / k;

                if (n >= 100 && n <= 999) {
                    // return true, k, n
                    int[] vals = {k, n};
                    return vals;
                }
            }

            k++;
        }

        // return false, -1, -1
        int[] vals = {-1, -1};
        return vals;
    }
}
