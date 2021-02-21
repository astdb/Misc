/*
Given an integer n, add a dot (".") as the thousands separator and return it in string format.

Example 1:

Input: n = 987
Output: "987"
Example 2:

Input: n = 1234
Output: "1.234"
Example 3:

Input: n = 123456789
Output: "123.456.789"
Example 4:

Input: n = 0
Output: "0"


Constraints:

0 <= n < 2^31

*/

import java.util.*;
import java.lang.StringBuilder;

public class ThousandSep {
  public static void main(String[] args) {
    int[] tests = {987, 1234, 123456789, 0};

    for(int test: tests) {
      System.out.printf("thousandSeparator(%d) = %s\n", test, thousandSeparator(test));
      // System.out.printf("getDigits(%d) = %s\n", test, getDigits(test).toString());

    }
  }

  private static String thousandSeparator(int n) {
    List<Integer> digs = getDigits(n);
    StringBuilder res = new StringBuilder();

    // for(Integer i: digs) {
    for(int i = digs.size()-1; i >= 0; i--) {
      res.append(digs.get(i).toString());

      int remLen = i;
      if((remLen > 0) && ((remLen % 3) == 0)) {
        res.append(",");
      }
    }
    
    return res.toString();
  }

  private static List<Integer> getDigits(int n) {
    List<Integer> digits = new ArrayList<Integer>();
    
    int rem = n % 10;
    n = n / 10;
    digits.add(rem);

    while(n > 0) {
      rem = n % 10;
      n = n / 10;
      digits.add(rem);
    }

    return digits;
  }
}

