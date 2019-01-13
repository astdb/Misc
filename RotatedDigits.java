/*
X is a good number if after rotating each digit individually by 180 degrees, we get a valid number that is different from X.  Each digit must be rotated - we cannot choose to leave it alone.

A number is valid if each digit remains a digit after rotation. 0, 1, and 8 rotate to themselves; 2 and 5 rotate to each other; 6 and 9 rotate to each other, and the rest of the numbers do not rotate to any other number and become invalid.

Now given a positive number N, how many numbers X from 1 to N are good?

Example:
Input: 10
Output: 4
Explanation: 
There are four good numbers in the range [1, 10] : 2, 5, 6, 9.
Note that 1 and 10 are not good numbers, since they remain unchanged after rotating.
Note:

N  will be in range [1, 10000].
*/

import java.util.*;

public class RotatedDigits {
  public static void main(String[] args) {
    int[] tests = new int[]{1,2,3,4,5,6,7,8,9,10};

    for(int test: tests) {
      System.out.println("rotatedDigits(" + test + ") == " + rotatedDigits(test));
    }
  }

    public static int rotatedDigits(int x) {
        int totalGood = 0;
        for(int i = 1; i <= x; i++) {
            if(good(i)) {
                totalGood++;
            }
        }

        return totalGood;
    }

    public static boolean good(int x) {
        List<Integer> xDigits = getDigits(x);
        List<Integer> xDigitsRotated = new ArrayList<Integer>();

        for(Integer d: xDigits) {
            if(rotatable(d)) {
                xDigitsRotated.add(rotate(d));
            } else {
                // unrotatable digit found in x - invalid
                return false;
            }
        }

        if(x != getValue(xDigitsRotated)) {
            // rotated x generates valid digit different to x - x is 'good'
            return true;
        }

        return false;
    }

  public static List<Integer> getDigits(int n) {
      // TODO: cater for non-negatives
      List<Integer> digits = new ArrayList<Integer>();
      int rem = 0;

      while(n > 0) {
        rem = n % 10;
        n = n / 10;
        digits.add(rem);
      }

      return digits;
  }

  public static int getValue(List<Integer> n) {
      double power = 0.0;
      int val = 0;
      for(int x: n) {
          val += x * (int)Math.pow(10.0, power);
          power++;
      }

      return val;
  }

  public static boolean rotatable(int n) {
    if(n == 3 || n == 4 || n == 7) {
        return false;
    }

    return true;
  }

  public static int rotate(int n) {
    Map<Integer, Integer> rotations = new HashMap<Integer, Integer>();
    rotations.put(0,0);
    rotations.put(1,1);
    rotations.put(8,8);
    rotations.put(2,5);
    rotations.put(5,2);
    rotations.put(6,9);
    rotations.put(9,6);

    return rotations.get(n);
  }
}
