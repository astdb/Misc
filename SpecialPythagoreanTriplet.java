
// Find the product of the Pythagorean triplet whose addition is 1000.

public class SpecialPythagoreanTriplet {
    public static void main(String[] args) {

        for(int a = 1; a < 1000; a++) {
            for(int b = a+1; b < 1000; b++) {
                for(int c = b+1; c < 1000; c++) {
                    if( ((a * a) + (b * b) == (c * c)) && (a + b + c == 1000)) {
                        System.out.printf("a = %d, b = %d, c = %d\n", a, b, c);
                        System.out.printf("abc = %d\n", a * b * c);
                    }
                }
            }
        }
    }
}
