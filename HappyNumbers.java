
public class HappyNumbers {
    public static void main(String[] args) throws InterruptedException {
        System.out.println("Result: " + numbers_mood(10));
        System.out.println("Result: " + numbers_mood(12));
        System.out.println("Result: " + numbers_mood(19));
    }

    private static String numbers_mood(int number) throws InterruptedException {
        int total = 0;
        int input = number;
        int remainder = 0;

        // per squared total
        do {
            // per digit
            do {
                remainder = input % 10;
                input = input / 10;

                total += Math.pow(remainder, 2);
            } while(input > 0);

            if(total == 1) {
                return "happy";
            }

            if(total == 4) {
                return "sad";
            }

            input = total;
            total = 0;
        } while(true);
    }
}
