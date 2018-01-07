
// Given a non-negative integer represented as a non-empty array of digits, return the value of adding one to it. Return value must be an int array too.
// Assume it contains no leading zeros, except for the whole number 0 itself. 

import java.util.*;

public class PlusOneDigits {
    public static void main(String[] args) {
        int[][] tests = {{1,2,3}, {1}, {0}, {0,1,0}, {}, {0,0,0,0,0,0,0,0,2,3}, {1,0}};

        for(int[] testcase: tests) {
            System.out.printf("%s: %s\n", Arrays.toString(testcase), Arrays.toString(plusOne(testcase)));
        }
    }

    public static int[] plusOne(int[] digits) {
        int digit = 0;
        int base = 10;
        int position = 0;

        for(int i = digits.length-1; i >= 0; i--){
            digit = digit + (digits[i] * (int) Math.pow(base, position));
            position++;
        }

        int digitPlusOne = digit + 1;
        int[] result = new int[String.valueOf(digitPlusOne).length()];
        int resultPos = result.length - 1;

        int rem = digitPlusOne % base;
        digitPlusOne = digitPlusOne / base;
        result[resultPos] = rem;
        resultPos--;

        while(digitPlusOne > 0) {
            rem = digitPlusOne % base;
            digitPlusOne = digitPlusOne / base;
            result[resultPos] = rem;
            resultPos--;
        }

        return result;
    }
}
