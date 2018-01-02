
/*
    The Hamming distance between two integers is the number of positions at which the corresponding bits are different.
    Given two integers x and y, calculate the Hamming distance.
*/

import java.util.*;

public class Hamming {
    public static void main(String[] args) {
        if(args.length < 2) {
            System.out.println("Usage: $> java Hamming <x> <y>");
            return;
        }

        int x, y = 0;
        try {
            x = Integer.parseInt(args[0]);
            y = Integer.parseInt(args[1]);
        } catch(NumberFormatException e) {
            System.out.println("Input must be numeric.");
            return;
        }

        if(x <= 0 || y <= 0) {
            System.out.println("Input must be positive integers.");
        }

        System.out.println(hammingDistance(x,y));
    }

    public static int hammingDistance(int x, int y) {
        int n = x ^ y;
        int dist = 0;

        int rem = n % 2;
        n = n / 2;
        if(rem == 1) {
            dist++;
        }

        while(n > 0) {
            rem = n % 2;
            n = n / 2;            
            if(rem == 1) {
                dist++;
            }
        }

        return dist;
    }
}
