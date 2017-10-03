
import java.util.*; 
import java.io.*;
import java.util.regex.*;

// CommentTracker reads a source file and prints out comments (both // and /* */ styles)
public class CommentTracker {
    public static void main(String args[] ) throws Exception {
        /* Enter your code here. Read input from STDIN. Print output to STDOUT */
        String inputFile = args[0];
        Scanner input = null;

        try {
            input = new Scanner(new File(inputFile));
        } catch( FileNotFoundException e ) {
            System.out.println("No such file: " + inputFile);
            return;
        }
      
        // literal patterns for single line comment opener, multiline comment opener and multiline comment closer
        Pattern singleLineComment = Pattern.compile("//");
        Pattern multiLineCommentStart = Pattern.compile("/\\*");
        Pattern multiLineCommentEnd = Pattern.compile("\\*/");
      
         boolean inMultiLineComment = false;        // flag indicating if in a multiline comment (e.g. read a '/*' and looking for a '*/')
         String multiLineCommentContent = "";       // content of the multiline comment being read
      
         // for each line in source file
         while(input.hasNextLine()) {
            String sourceLine = input.nextLine();   /* // read */   // [that's just a weird comment for dogfooding (;]
            if(sourceLine != null){
                sourceLine = sourceLine.trim();
              
                if(!inMultiLineComment) {
                  // check for single/multiline comment starts in current source line
                  Matcher singleOpenMatch = singleLineComment.matcher(sourceLine);
                  Matcher multiOpenMatch = multiLineCommentStart.matcher(sourceLine);

                  // if a line contains both // and /*, find which comes first
                  if(singleOpenMatch.find() && multiOpenMatch.find()){
                      int singleStart = singleOpenMatch.start();
                      int multiStart = multiOpenMatch.start();

                      if(singleStart < multiStart) {
                          // single line comment
                          System.out.println(sourceLine.substring(singleOpenMatch.start()));
                      } else {
                          // multiline comment start
                          inMultiLineComment = true;
                          multiLineCommentContent  = "\n" + sourceLine.substring(multiOpenMatch.start());

                          // does this multiline comment close on this line?
                          while(sourceLine.length() > multiStart+2) {
                              String sourceLine = sourceLine.substring(multiStart+2);
                              
                              Matcher multiCloseMatch = multiLineCommentEnd.matcher(restOfLine);

                              if(multiCloseMatch.find()) {
                                  
                              }

                          }
                      }
                  } else if(singleOpenMatch.find()){
                      // a single-line comment starter was found in this source line - print line from that index
                      System.out.println(sourceLine.substring(singleOpenMatch.start()));
                  } else if(multiOpenMatch.find()){
                      // a multiline comment starter was found in this source line - print line from that index
                      // multiline comment start
                      inMultiLineComment = true;
                      multiLineCommentContent  = "\n" + sourceLine.substring(multiOpenMatch.start());
                  }
                  
                } else {
                  // already within a multiline source comment - look for multiline comment close
                  Matcher multiCloseMatch = multiLineCommentEnd.matcher(sourceLine);

                  if(multiCloseMatch.find()) {
                      inMultiLineComment = false;
                      multiLineCommentContent += "\n" + sourceLine.substring(0, multiCloseMatch.start());

                      // TODO: what if a multi/single comment starts after the */ ?
                      // need to run line processor for rest of this line
                  }
                  
                }
                
            }
         }
    }
}
