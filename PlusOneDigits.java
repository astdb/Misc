
// Given a non-negative integer represented as a non-empty array of digits, return the value of adding one to it. Return value must be an int array too.
// Assume it contains no leading zeros, except for the whole number 0 itself. 

import java.util.*;

public class PlusOneDigits {
    public static void main(String[] args) {
        // int[][] tests = {{1,2,3}, {1}, {0}, {0,1,0}, {}, {0,0,0,0,0,0,0,0,2,3}, {1,0}, {9}, {9,9}, {8,9,9,9}};
        int [][] tests = {{7}, {8}, {9}, {9,9}, {7,9,9}, {8,9,9,9}};

        for(int[] testcase: tests) {
            System.out.printf("%s: %s\n", Arrays.toString(testcase), Arrays.toString(plusOne(testcase)));
        }
    }

    public static int[] plusOne(int[] digits) {
        int digit = 0;
        int base = 10;
        int position = 0;
        int carry = 0;

        System.out.println("--------------------------\nOriginal: " + Arrays.toString(digits) + "\nModding digits[]...");
        for(int i = digits.length-1; i >= 0; i--) {
            int positionResult = 0;
            if(i == digits.length-1) {
                positionResult = digits[i] + 1 + carry;
            } else {
                positionResult = digits[i] + carry;
            }
            
            // carry = 0;
            System.out.println("\t\tPositionResult: " + positionResult);
            
            if(positionResult >= base) {
                digits[i] = positionResult % base;
                carry = positionResult / base;
                System.out.printf("\t\tcarry (%d) = positionResult (%d) / base (%d)\n", carry, positionResult, base);
            } else {
                digits[i] = positionResult;
                carry = 0;
            }

            System.out.println("\t" + Arrays.toString(digits) + " | Carry = " + carry);
        }

        System.out.println("Done modifying digits[] in-place.\n");

        if(carry == 0) {
            System.out.println("Carry == 0 - returning digits[]");
            return digits;
        }

        System.out.println("Carry == " + carry + " - constructing newdigits[]...");
        int carryLen = String.valueOf(carry).length();
        int[] newdigits = new int[carryLen + digits.length];

        int carryPos = carryLen-1;
        int carryRem = carry % base;
        carry = carry / base;
        newdigits[carryPos] = carryRem;
        System.out.println(Arrays.toString(newdigits));
        carryPos--;

        while(carry > 0){
            carryRem = carry % base;
            carry = carry / base;
            newdigits[carryPos] = carryRem;
            System.out.println(Arrays.toString(newdigits));
            carryPos--;
        }

        int j = 0;
        for(int i = carryLen; i < newdigits.length; i++) {
            newdigits[i] = digits[j];
            j++;
        }

        return newdigits;
    }

    public static int[] plusOne1(int[] digits) {
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
