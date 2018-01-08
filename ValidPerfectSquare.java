
public class ValidPerfectSquare {
    public static void main(String[] args) {
        int[] tests = {0,1,2,3,4,5,6,7,8,9,16,14,100000000, 2147483647};

        for(int testcase: tests) {
            System.out.printf("%d\t%s\n", testcase, isPerfectSquare(testcase));
        }
    }

    public static boolean isPerfectSquare(int num) {
        // addition of first n odd numbers is a perfect square
        for(int i = 1, sum = 0; sum <= num; i += 2) {
            sum += i;
            
            if(sum == num) {
                return true;
            }
        }

        return false;
    }

    // public static boolean isPerfectSquare(int num) {
    //     for(int i = 0; i <= num; i++) {
    //         if(i*i == num) {
    //             return true;
    //         }

    //         if(i*i > num) {
    //             return false;
    //         }
    //     }

    //     return false;
    // }
}
