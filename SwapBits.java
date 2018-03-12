
// given a 64-bit number and two bit locations (in ints), swap the bit values in those positions. 
// NOTE: index 0 is LSB, 63 is MSB
// the bit swap will be necessary when the values are different

public class SwapBits {
    public static void main(String[] args) {
        swapBits(24L, 1, 2);
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

    // print array representing binary number in string format
    public static String printBin(long[] x) {
        String n = "";
        for(int i = x.length-1; i >= 0; i--) {
            n = n + x[i];
        }
        
        return n;
    }
}
