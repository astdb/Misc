
// write a function that takes an unsigned 64-bit integer and returns the unsigned 64-bit integer value
// consisting of the input of the bits of the input in reverse order. 

public class ReverseBits {
    public static void main(String[] args) {
        long[] testcases = new long[]{1L, 50L, 0L, 60L, 5648L, Long.MAX_VALUE, Long.MIN_VALUE};

        for(long test: testcases) {
            System.out.println("reverseBits(" + test + ") == " + reverseBits(test));
        }
    }

    private static long reverseBits(long x) {
        long rev = 0;
        double pow = 63;
        long rem = x % 2;
        x = x / 2;
        rev = (long)(rem * Math.pow(2.0, pow));
        pow--;

        while(x > 0) {
            rem = x % 2;
            x = x / 2;
            rev = (long)(rem * Math.pow(2.0, pow));
            pow--;
        }

        return rev;
    }
}
