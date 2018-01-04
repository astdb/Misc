
/*
    A self-dividing number is a number that is divisible by every digit it contains.
    For example, 128 is a self-dividing number because 128 % 1 == 0, 128 % 2 == 0, and 128 % 8 == 0.
    Also, a self-dividing number is not allowed to contain the digit zero.
    Given a lower and upper number bound, output a list of every possible self dividing number, including the bounds if possible.
*/

import java.util.*;

public class SelfDiv {
    public static void main(String[] args) {
        // System.out.println(selfDividing(140));
        if(args.length < 2) {
            System.out.println("Usage: $> java SelfDiv <lowerlimit> <upperlimit>");
            return;
        }

        int upper, lower = 0;
        try {
            lower = Integer.parseInt(args[0]);
            upper = Integer.parseInt(args[1]);
        } catch(NumberFormatException e) {
            System.out.println("The input has to be integer.");
            return;
        }

        

    }

    public List<Integer> selfDividingNumbers(int left, int right) {
        ArrayList<Integer> selfDividingNumbers = new ArrayList<Integer>();

        for(int i = left; i <= right; i ++) {
            if(selfDividing(i)) {
                selfDividingNumbers.add(new Integer(i));
            }
        }

       return selfDividingNumbers; 
    }

    public static boolean selfDividing(int n) {
        int ncopy = n;     // make copy of n, to extract digits
        
        // start extracting decimal digits of n
        int rem = ncopy % 10;
        ncopy = ncopy / 10;

        if(rem == 0) {
            return false;   // a self-dividing number cannot contain zero
        }

        if(n % rem != 0) {
            return false;
        }

        while(ncopy > 0) {
            rem = ncopy % 10;
            ncopy = ncopy / 10;

            if(rem == 0) {
                return false;   // a self-dividing number cannot contain zero
            }

            if(n % rem != 0) {
                return false;
            }
        }

        return true;
    }
}
