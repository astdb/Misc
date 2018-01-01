
/**
    By listing the first six prime numbers: 2, 3, 5, 7, 11, and 13, we can see that the 6th prime is 13.
    What is the 10 001st prime number?
*/

import java.util.*;

public class XthPrime {
    public static void main(String[] args) {
        if(args.length < 1) {
            System.out.println("Usage: $> java XthPrime <input>");
            return;
        }

        int primeNum = 0;
        try {
            primeNum = Integer.parseInt(args[0]);
        } catch(NumberFormatException e) {
            System.out.println("Input must be numeric.");
            return;
        }

        if(primeNum <= 0) {
            System.out.println("Input must be a positive integer.");
        }

        int i = 1;      // counter of primes

        for(int j = 2; i <= primeNum; j++) {
            if(j == 2 || j == 3) {
                // 2 is prime by definition - the only 'even prime'
                if(i == primeNum) {
                    System.out.printf("%d. %d\n", i, j);
                }

                i++;
            } else {
                // test j for divisability upto sqrt(j)
                for(int k = 2; k < (int) Math.sqrt((double) j); k++) {
                    if(j % k == 0) {
                        // divisable by k - goto next j
                        break;
                    }

                    if(k > (int) Math.sqrt((double) j)) {
                        // j is prime - is it the one we're looking for?
                        if(i == primeNum) {
                            System.out.printf("%d. %d\n", i, j);
                        }

                        i++;
                    }
                }
            }
        }
    }
}
