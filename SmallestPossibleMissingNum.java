
import java.util.*;

public class SmallestPossibleMissingNum {
  public static void main(String args[]) {
    System.out.println("Hello, World!");


    int[][] tests = {{1,3,6,4,1,2}, {1,2,3}, {-1, -3}};

    for(int[] test: tests) {
      // System.out.println(Arrays.toString(test));
      System.out.printf("smallestMissing(%s) = %d\n", Arrays.toString(test), smallestMissing(test));
    }
  }

  public static int smallestMissing(int[] nums) {
    Arrays.sort(nums);

    int j = 1;
    for(int i = 0; i < nums.length; i++) {
      // System.out.println("");
      if(nums[i] != j) {
        // return j;

      } else {
        if(i < nums.length-1) {
          if(nums[i] == nums[i+1]) {
            // next element should be compsred to un-incremented j
            // do nothing
            
          } else {
            // arr[i] == j, arr[i+1] must be compared to j+1
					  j++;
          }
        } else {
          // arr[i] == j, and its the last element in arr - j+1 must be the answer
				  j++;
        }
      }
    }      

    return j;
  }
}

