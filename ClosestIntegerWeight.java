
// weight of a nonnegative integer x is the number of bits that a re set to one in its binary representation.
// e.g. 92 in binary is 1011100 and therefore its weight is 4.
// write a program which takes a positive int x and returns a value(y) which is not equal to x but |x-y| is minimum. 

public class ClosestIntegerWeight {
    public static void main(String[] args) {
        System.out.println("closestWeight(92) = " + closestWeight(92));
        System.out.println("closestWeight(5) = " + closestWeight(5));
        System.out.println("closestWeight(6) = " + closestWeight(6));
        // System.out.println("closestWeight(1073741824) = " + closestWeight(1073741824));
        System.out.println("closestIntSameWeight(1073741824L) = " + closestIntSameWeight(1073741824L));
    }

    // if bits ant indexes k1 and k2 are flipped (where k! > k2), the the absolute value between the 
    // original integer and the result of the swaps would be (2^k1 - 2^k2). To minimize this,
    // k1 must be made as small as possible and k2 as close as possible to k1. 
    // the most efficient technique is to swap the two rightmost consecutive bits that differ. 
    public static long closestIntSameWeight(long x) {
        // x is given to be nonnegative with leading bit 0
        // therefore only 63 LSB's to consider
        final int NUM_UNSIGNED_BITS = 63;

        for(int i = 0; i < NUM_UNSIGNED_BITS; ++i) {
            if((((x >>> i) & 1) != ((x >>> (i + 1)) & 1))) {
                x ^= (1L << i) | (1L << (i + 1));   // swaps bit-i and bit-(i + 1)
                return x;
            }
        }

        // error if all bits of x are 0 or 1
        throw new IllegalArgumentException("All bits are 0 or 1.");
    }

    // works horribly on certain inputs (e.g. large full powers of 2)
    private static int closestWeight(int n) {
        int n_weight = weight(n);
        int x = 1;
        while(true) {            
            System.out.println("\n----------------------------------------");
            System.out.println("n = " + n);
            System.out.println("n_weight = " + n_weight);
            System.out.println("x = " + x);

            // check x below n
            System.out.println("(n - x) = " + (n - x));
            if((n - x) > 0) {
                System.out.println("weight(n - x) = " + weight(n - x));
                if(weight(n - x) == n_weight) {
                    return (n - x);
                }
            }

            // check x above n
            System.out.println("(n + x) = " + (n + x));
            // if((Integer.MAX_VALUE - n) > x) {
            if((Integer.MAX_VALUE - x) > n) {
                System.out.println("weight(n + x) = " + weight(n + x));
                if(weight(n + x) == n_weight) {
                    return (n + x);
                }
            }
            x++;
        }
    }

    private static int weight(int x) {
        int weight = 0;
        int rem = 0;

        while(x > 0) {
            rem = x % 2;
            x = x / 2;
            weight += rem;
        }

        return weight;
    }
}
