
public class PrimeFac {
    public static void main(String[] args) {
        if(args.length < 1) {
            System.out.println("Usage: $> java PrimeFac <NumberToFactor>");
            return;
        }

        long num = 0;
        try {
            num = Long.parseLong(args[0]);

        } catch(NumberFormatException e) {
            System.out.println("Input needs to be a valid numeric value.");
            return;
        }

        if(num <= 0) {
            System.out.println("Please enter a positive integer.");
            return;
        }

        long largestFactor = 0;

        while(num % 2 == 0) {
            System.out.printf("%d ", 2);
            num = num / 2;
            largestFactor = 2;
        }

        for(int i = 3; (double) i < Math.sqrt((double) num); i += 2) {
            while(num % i == 0) {
                if(i > largestFactor) {
                    largestFactor = i;
                }

                System.out.printf("%d ", i);
                num = num / i;
            }
        }

        if(num > 2) {
            if(num > largestFactor) {
                largestFactor = num;
            }

            System.out.printf("%d ", num);
        }

        System.out.printf("\nLargest factor: %d\n", largestFactor);
    }
}
