/**
    2520 is the smallest number that can be divided by each of the numbers from 1 to 10 without any remainder.
    What is the smallest positive number that is evenly divisible by all of the numbers from 1 to 20?
*/

public class SmallestMultiple {
    public static void main(String[] args) {
        for(int i = 1; i > 0; i++) {
            if(divisible(i)) {
                System.out.println(i);
                return;
            }
        }
    }

    private static boolean divisible(int n) {
        for(int i = 1; i <= 20; i++) {
            if(n % i != 0) {
                return false;
            }
        }

        return true;
    }
}
