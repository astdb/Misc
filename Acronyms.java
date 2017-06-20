
public class Acronyms {
    public static void main(String[] args) {
        String[] w1 = {"What", "You", "See", "IS", "What", "You", "Get"};
        System.out.println("" + acronym(w1));

        String[] w2 = {"Zone", "Improvement", "Plan"};
        System.out.println("" + acronym(w2));

        String[] w3 = {"DataBase"};
        System.out.println("" + acronym(w3));

        String[] w4 = {"kitkat"};
        System.out.println("" + acronym(w4));

        String[] w5 = {"KitKat"};
        System.out.println("" + acronym(w5));
    }

    private static String acronym(String[] input) {
        int n = 0;
        String firstWord = "";
        String acronym = "";

        for (String str: input) {
            if(str != null) {
                str = str.trim();

                if(str.equals("")) {
                    continue;
                }

                // count valid word, cache it if it's the first in array
                if(n == 0) {
                    firstWord = str;
                }

                // get starting letter
                char first = str.charAt(0);

                boolean allUpper = true;
                if(str.length() <= 2) {
                    // check if length 1/2 string is all upper
                    for(int i = 0; i < str.length(); i++) {
                        if(!Character.isUpperCase(str.charAt(i))) {
                            allUpper = false;
                        }
                    }

                    // words of length 1 or 2 get an acronym place only if they're alll caps
                    if(allUpper) {
                        acronym += Character.toUpperCase(str.charAt(0)) + "";
                    }
                } else {
                    // word with length > 2  - record first char in uppercase in acronym
                    acronym += Character.toUpperCase(str.charAt(0)) + "";
                }

                n++;
            }
        }

        // if the input had only one valid word
        if(n == 1) {
          acronym = "";
          boolean caps = false;
          for(int i = 0; i < firstWord.length(); i++) {
                if(Character.isUpperCase(firstWord.charAt(i))) {
                    caps = true;
                    acronym += Character.toUpperCase(firstWord.charAt(i)) + "";
                }
          }

          // if the single input word did not have any uppercase letters
          if(!caps) {
            acronym = (firstWord.charAt(0) + "").toUpperCase();
          }
        }

        if(acronym.length() > 1) {
            // insert delimiter
            String tempAcr = "";

            for(int i = 0; i < acronym.length(); i++) {
                if((i + 1) < acronym.length()) {
                    tempAcr += acronym.charAt(i) + ".";
                } else {
                    tempAcr += acronym.charAt(i) + "";
                }
            }

            acronym = tempAcr;
        }

        return acronym;
    }
}
