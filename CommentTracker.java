import java.io.*;
import java.util.regex.*;

public class CommentTracker {
    public static void main(String args[] ) throws Exception {
        /* Enter your code here. Read input from STDIN. Print output to STDOUT */
        String inputFile = args[0];
        Scanner input = null;

        try {
            input = new Scanner(new File(inputFile));
        } catch( FileNotFoundException e ) {
            // System.out.println("No such file: " + inputFile);
            return null;
        }
      
        Pattern singleLineComment = Pattern.compile("//");
        Pattern multiLineCommentStart = Pattern.compile("/*");
        Pattern multiLineCommentEnd = Pattern.compile("*/");
      
         boolean inMultiLineComment = false;
         String multiLineCommentContent = "";
      
         while(input.hasNextLine()) {
            String sourceLine = input.nextLine();
            if(sourceLine != null){
                sourceLine = sourceLine.trim();
              
                if(!inMultiLineComment){
                  // check for single/multiline comment starts in current source line
                  Matcher singleOpenMatch = singleLineComment.matcher(sourceLine);
                  Matcher multiOpenMatch = multiLineCommentStart.matcher(sourceLine);
                  
                  
                } else {
                  // already within a multiline source comment - look for multiline comment close
                  Matcher multiCloseMatch = multiLineCommentStart.matcher(sourceLine);
                  
                }
                
            }
    }
}
