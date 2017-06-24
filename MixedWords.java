
// find if a given string is a valid shuffle of two other given strings

public class MixedWords {
    public static void main(String[] args) {
        if(mixedWords("ab2", "cd", "d2abc").equals("yes")) {
            System.out.println("Pass");
        } else {
            System.out.println("Fail");
        }

        if(mixedWords("Gilgamesh", "Spacecraft", "Fail").equals("no")) {
            System.out.println("Pass");
        } else {
            System.out.println("Fail");
        }

        if(mixedWords("abcd", "efgh", "abcdefghh").equals("no")) {
            System.out.println("Pass");
        } else {
            System.out.println("Fail");
        }

        if(mixedWords("Carmina", "Burana", "Carmina Burana").equals("yes")) {
            System.out.println("Pass");
        } else {
            System.out.println("Fail");
        }
    }

    private static String mixedWords(String input1, String input2, String input3) {
        if(input1 == null || input2 == null || input3 == null) {
            System.out.println("Null input(s)");
            return "no";
        }

        // get rid of possible middle whitespace
        String[] input1_exp = input1.trim().split(" ");
        String[] input2_exp = input2.trim().split(" ");
        String[] input3_exp = input3.trim().split(" ");

        StringBuilder sb1 = new StringBuilder();
        for(int i = 0; i < input1_exp.length; i++) {
            sb1.append(input1_exp[i].trim());
        }

        StringBuilder sb2 = new StringBuilder();
        for(int i = 0; i < input2_exp.length; i++) {
            sb2.append(input2_exp[i].trim());
        }

        StringBuilder sb3 = new StringBuilder();
        for(int i = 0; i < input3_exp.length; i++) {
            sb3.append(input3_exp[i].trim());
        }

        // transform strings to char arrays
        char[] input1_charr = sb1.toString().trim().toLowerCase().toCharArray();
        char[] input2_charr = sb2.toString().trim().toLowerCase().toCharArray();
        char[] input3_charr = sb3.toString().trim().toLowerCase().toCharArray();

        if((input1_charr.length + input2_charr.length) != input3_charr.length) {
            System.out.println("String lengths don't add up.");
            return "no";
        }

        // check if all str3 characters come from one or other of str1 or str2
        for(int i = 0; i < input3_charr.length; i++) {
            char curChar = input3_charr[i];

            boolean found = false;
            for(int j = 0; j < input1_charr.length; j++) {
                if(curChar == input1_charr[j]) {
                    input1_charr[i] = ' ';
                    found = true;
                    break;
                }
            }

            if(found){
                continue;
            }

            for(int j = 0; j < input2_charr.length; j++) {
                if(curChar == input2_charr[j]) {
                    input2_charr[j] = ' ';
                    found = true;
                    break;
                }
            }

            if(!found){
                // string3 had a character not found in either string1 or string2
                System.out.println(curChar + " wasn't found in either string.");
                return "no";
            }
        }

        // if we're here
        //  - string3 length = str1 len + str2 len
        //  - every char in str3 was found in either str1 or two
        // i.e. valid shuffle
        return "yes";
    }
}
