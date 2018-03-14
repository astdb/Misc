
// given a 64-bit number and two bit locations (in ints), swap the bit values in those positions. 
// NOTE: index 0 is LSB, 63 is MSB
// the bit swap will be necessary when the values are different

public class SwapBits {
    public static void main(String[] args) {
        swapBits(24L, 1, 2);
        swapBits(1, 2, 24L);
    }

    // swap bits in ith and jth locations of x and return resulting x
    public static long swapBits(long x, int i, int j) {
        // extract bits into array
        long[] bits = new long[64];
        int count = 0;
        long rem = x % 2;
        x = x / 2;
        bits[count] = rem;
        count++;

        while(x > 0) {
            rem = x % 2;
            x = x / 2;
            bits[count] = rem;
            count++;
        }

        // swap bits if required
        if(bits[i] != bits[j]) {
            long temp = bits[i];
            bits[i] = bits[j];
            bits[j] = temp;
        }

        // reconstruct & return
        long total = 0L;
        for(int k = 0; k < bits.length; k++) {
            total += bits[k] * Math.pow(2, i);
        }
        return total;
    }

    // swap bits using bitwise operations (O(1))
    public static long swapBits(int i, int j, long x) {
        // extract the ith and jth bits and check if they differ
        if(((x >>> i) & 1) != ((x >>> j) & 1)) {
            // bits differ - swap by flipping values

            // select bits to flip w/ bitmask. Since x ^ 1 = 0 when x = 1 and 1
            // when x = 0, we can perform the flip XOR
            long bitMask = (1L << i) | (1L << j);
            x ^= bitMask;
        }

        return x;
    }

    // print array representing binary number in string format
    public static String printBin(long[] x) {
        String n = "";
        for(int i = x.length-1; i >= 0; i--) {
            n = n + x[i];
        }
        
        return n;
    }
}
