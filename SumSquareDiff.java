/*
    The sum of the squares of the first ten natural numbers is,

    12 + 22 + ... + 102 = 385
    The square of the sum of the first ten natural numbers is,

    (1 + 2 + ... + 10)2 = 552 = 3025
    Hence the difference between the sum of the squares of the first ten natural numbers and the square of the sum is 3025 − 385 = 2640.

    Find the difference between the sum of the squares of the first one hundred natural numbers and the square of the sum.
*/

public class SumSquareDiff {
    public static void main(String[] args) {
        int sumSquares = 0;     // sum of sqaures of every number from 1-100
        int sum = 0;            // sum of all numbers from 1-100

        for(int i = 1; i <= 100; i++) {
            sum += i;
            sumSquares += (i*i);
        }

        int squareSum = (sum*sum);      // square of sum of all numbers from 1-100
        int sumSquareDiff = squareSum - sumSquares;

        System.out.println(sumSquareDiff);

    }
}
