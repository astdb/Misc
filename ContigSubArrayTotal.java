
// find the contiguous subarray with the max total of elements, given an array of ints

import java.util.*;

public class ContigSubArrayTotal {
    public static void main(String[] args) {
        int[] input = {-2,1,-3,4,-1,2,1,-5,4};      // expect [4,-1,2,1]
        System.out.println(Arrays.toString(maxSubArray(input)));
    }

    public static int[] maxSubArray(int[] nums) {

        // array to store maximum subarray total and subarray boundaries e.g. max = {total, substart, subend+1}
        int[] maxSubTot = {0,0,0};

        // subarray size
        for(int i = 1; i < nums.length-1; i++) {
            // go through subarrays of i length, starting from index j
            for(int j = 0; j+i <= nums.length ; j++) {
                if(i == 1 && j == 0) {
                    // inititalize max
                    maxSubTot[0] = nums[0];
                    maxSubTot[1] = 0;
                    maxSubTot[2] = 1;
                } else {
                    int tot = 0;
                    for(int k = j; k < j+i; k++) {
                        tot =+ nums[k];
                    }

                    if(tot > maxSubTot[0]) {
                        maxSubTot[0] = tot;
                        maxSubTot[1] = j;
                        maxSubTot[2] = j+i;
                    }
                }
            }
        }

        int[] result = new int[maxSubTot[2]-maxSubTot[1]];
        for(int i = maxSubTot[1], j = 0; i < maxSubTot[2]; i++, j++){
            result[j] = nums[i];
        }

        return result;
    }
}
