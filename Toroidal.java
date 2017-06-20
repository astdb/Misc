
public class Toroidal {
    public static void main(String[] args) {
        System.out.println(toroidalMatrix("1"));
        System.out.println(toroidalMatrix("1 38"));
        System.out.println(toroidalMatrix("1 38 11 63"));
        System.out.println(toroidalMatrix("1 2 3 4 5 6 7 8 9"));
        System.out.println(toroidalMatrix("21 22 23 14 13"));
        System.out.println(toroidalMatrix("1 2 3 4 5 6 7 8 9 10 11 12 13 14 15 16"));
    }

    private static String toroidalMatrix(String input) {
        if(input == null || input.trim().equals("")) {
            return "-";
        }

        // read string into array     
        String[] input_arr = input.trim().split(" ");
        long input_len = input_arr.length;

        // check if input length is square, allowing a square matrix
        long root = (long) Math.sqrt(input_len);
        if(input_len != root * root) {
            // cannot form square matrix with given numbers string
            return "-";
        }

        // can form square matrix - declare multi-dimensional array to store matrix
        String[][] matrix = new String[(int)root][(int)root];

        // populate matrix with input data
        int p = 0;
        for(int i = 0; i < root; i++) {
            for(int j = 0; j < root; j++) {
                matrix[i][j] = input_arr[p];
                p++;
            }
        }

       int i, k = 0, l = 0;     // iterator, starting row/col
       int m = (int)root, n = (int)root;     // end row/col
       String traversal = "";
        
        // traverse matrix
        while (k < m && l < n) {
            for (i = l; i < n; ++i) {
                traversal += matrix[k][i] + " ";
            }
            k++;
  
           for (i = k; i < m; ++i) {
                traversal += matrix[i][n-1] + " ";
            }
            n--;
  
            if ( k < m) {
                for (i = n-1; i >= l; --i) {
                    traversal += matrix[m-1][i] + " ";
                }
                m--;
            }
  
            if (l < n) {
                for (i = m-1; i >= k; --i) {
                    traversal += matrix[i][l] + " ";
                }
                l++;    
            }        
        }

        return traversal;
    }
}
