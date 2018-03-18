
// weight of a nonnegative integer x is the number of bits that a re set to one in its binary representation.
// e.g. 92 in binary is 1011100 and therefore its weight is 4.
// write a program which takes a positive int x and returns a value(y) which is not equal to x but |x-y| is minimum. 

public class ClosestIntegerWeight {
    public static void main(String[] args) {
        System.out.println("closestWeight(92) = " + closestWeight(92));
        System.out.println("closestWeight(5) = " + closestWeight(5));
        System.out.println("closestWeight(6) = " + closestWeight(6));

    }

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
