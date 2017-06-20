import java.util.*;

public class EgocentricNumbers {
    public static void main(String[] args) {
        System.out.println(egocentricNumbers(5));
        System.out.println(egocentricNumbers(153));
        System.out.println(egocentricNumbers(9982));
        System.out.println(egocentricNumbers(407));
    }

    private static int egocentricNumbers(int input) {
        // find length of number
        int n = 0;
        int remainder = 0;
        int in = input;
        ArrayList<Integer> digits = new ArrayList<Integer>();

        do {
            remainder = input % 10;
            input = input / 10;
            digits.add(remainder);
            n++;
        } while(input > 0);

        if(in == 0) {
            n = 1;
        }

        int total = 0;
        for (Integer digit: digits) {
            total += Math.pow(digit, n);
        }

        if(total == in) {
            return 1;
        }
        
        return 0;
    }
}
