import java.util.*;

public class BaseBallHits {
    public static void main(String args[]) {
        // String[] input = {"5", "-2", "4", "Z", "X", "9", "+", "+"};
        String[] input = {"1", "2", "+", "Z"};
        int input_size = 4;
        System.out.println(totalScore(input, input_size));
    }

    public static int totalScore(String[] blocks, int n) {
        String BLOCK_X = "X";
        String BLOCK_PLUS = "+";
        String BLOCK_Z = "Z";

        int blocks_len = blocks.length;
        int prev_score_1 = 0;
        int prev_score_2 = 0;
        int total_score = 0;

        for(int i = 0; i < n; i++) {
            if(i < blocks_len) {
                String current_block = blocks[i];

                if(isInt(current_block)) {
                    int thisScore = Integer.parseInt(current_block);
                    total_score = total_score + thisScore;
                    prev_score_2 = prev_score_1;
                    prev_score_1 = thisScore;

                } else if(BLOCK_X.equalsIgnoreCase(current_block)) {
                    int thisScore = prev_score_1 * 2;
                    total_score = total_score + thisScore;
                    prev_score_2 = prev_score_1;
                    prev_score_1 = thisScore;

                } else if(BLOCK_PLUS.equalsIgnoreCase(current_block)) {
                    int thisScore = prev_score_1 + prev_score_2;
                    total_score = total_score + thisScore;
                    prev_score_2 = prev_score_1;
                    prev_score_1 = thisScore;
                    
                } else if(BLOCK_Z.equalsIgnoreCase(current_block)) {
                    int thisScore = -2;
                    total_score = total_score - prev_score_1;
                    prev_score_2 = prev_score_1;
                    prev_score_1 = thisScore;
                    
                }
            } else {
                return total_score;
            }
        }

        return total_score;
    }

    public static boolean isInt(String intString) {
        try{
            Integer.parseInt(intString);
            return true;
        } catch (Exception e) {
            return false;
        }
    }
}
