
public class CountBits {
	public static void main(String[] args) {
		int[] tests = new int[]{1,2,3,4,5,6,7,8,9,0};

		for(int i = 0; i < tests.length; i++) {
			System.out.println("coutnBits(" + tests[i] + ") == " + countBits(i));
		}
	}

	// count the number of bits set to '1' in the input
	public static short countBits(int x) {
		short numBits = 0;
		while(x != 0) {
			numBits += (x & 1);
			x >>>= 1;
		}

		return numBits;
	}
}
