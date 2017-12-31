
/*
    If we list all the natural numbers below 10 that are multiples of 3 or 5, we get 3, 5, 6 and 9. The sum of these multiples is 23.
    Find the sum of all the multiples of 3 or 5 below 1000.
*/


public class MultThreeFive {
    public static void main(String[] args) {
        // read upper limit from command line
        if(args.length < 1) {
            System.out.println("Usage: $> java MultThreeFive <UpperLimit> (UpperLimit must be numeric)");
            return;
        }

        int upperLimit = -1;
        try {
            upperLimit = Integer.parseInt(args[0]);
        } catch(NumberFormatException e) {
            System.out.println("UpperLimit must be numeric");
            return;
        }

        if(upperLimit < 0) {
            System.out.println("UpperLimit must be a positive numeric value");
            return;
        }

        int total = 0;
        for(int i = 1; i < upperLimit; i++) {
            // if multiple of 3 or 5
            if(i % 3 == 0 || i % 5 == 0) {
                total = total + i;
            }            
        }

        System.out.println(total);
    }
}
