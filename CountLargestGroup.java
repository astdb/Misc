
import java.util.*;

class CountLargestGroup {
  public static void main(String[] args) {
    int[] tests = {13, 2, 15, 24};

    for(int test: tests) {
      System.out.printf("countLargestGroup(%d) = %d\n", test, countLargestGroup(test));
    }
  }

  private static int countLargestGroup(int n) {
    HashMap<Integer, Integer> digitTotals = new HashMap<Integer, Integer>();

    for(int i = 0; i <= n; i++) {
      int tot = digitTotal(i);

      boolean counted = digitTotals.containsKey(tot);

      if(counted == true) {
        int curVal = digitTotals.get(tot);
        curVal++;
        digitTotals.put(tot, curVal);
      } else {
        digitTotals.put(tot, 1);
      }
    }

    int largestCount = 0;
    int i = 0;

    Iterator dtIterator = digitTotals.entrySet().iterator();
    while (dtIterator.hasNext()) { 
      Map.Entry mapElement = (Map.Entry)dtIterator.next(); 
      // int marks = ((int)mapElement.getValue() + 10); 
      // System.out.println(mapElement.getKey() + " : " + marks); 
      int thisVal = (int)mapElement.getValue();
      
      if(i == 0) {
        largestCount = thisVal;
      } else if(thisVal > largestCount) {
        largestCount = thisVal;
      }
    }

    int groupCount = 0;

    Iterator dtIterator2 = digitTotals.entrySet().iterator();
    while (dtIterator2.hasNext()) { 
      Map.Entry mapElement = (Map.Entry)dtIterator2.next();

      int thisVal = (int)mapElement.getValue();

      if(thisVal == largestCount) {
        groupCount++;
      }
    }

    return groupCount;
 
  }

  private static int digitTotal(int x) {
    int tot = 0;
    int rem = 0;

    while(x > 0) {
      rem = x % 10;
      x = x / 10;

      tot = tot + rem;
    }

    return tot;
  }
}
