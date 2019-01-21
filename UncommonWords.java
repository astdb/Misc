/*
We are given two sentences A and B.  (A sentence is a string of space separated words.  Each word consists only of lowercase letters.)

A word is uncommon if it appears exactly once in one of the sentences, and does not appear in the other sentence.

Return a list of all uncommon words. 

You may return the list in any order. 

Example 1:

Input: A = "this apple is sweet", B = "this apple is sour"
Output: ["sweet","sour"]
Example 2:

Input: A = "apple apple", B = "banana"
Output: ["banana"]
*/

import java.util.*;

public class UncommonWords {
    public static void main(String[] args) {
        String[][] tests = {{"this apple is sweet", "this apple is sour"}, {"apple apple", "banana"}};

        for(String[] test: tests) {
            System.out.println(Arrays.toString(uncommonFromSentences(test[0], test[1])));
        }
    }

    public static String[] uncommonFromSentences(String strA, String strB) {
        // convert sentences into word arrays
        String[] strAArr = strA.split(" ");
        String[] strBArr = strB.split(" ");

        // build word maps for each sentence with words as keys and values as their occurance counts
        Map<String, Integer> strAWordMap = new HashMap<String, Integer>();
        Map<String, Integer> strBWordMap = new HashMap<String, Integer>();

        // populate maps
        for(String word: strAArr) {
            if(strAWordMap.containsKey(word)) {
                int count = strAWordMap.get(word);
                strAWordMap.put(word, count++);
            } else {
                strAWordMap.put(word,1);
            }
        }

        for(String word: strBArr) {
            if(strBWordMap.containsKey(word)) {
                int count = strBWordMap.get(word);
                strBWordMap.put(word, count++);
            } else {
                strBWordMap.put(word,1);
            }
        }

        // go through each map and build collection of uncommon words
        List<String> uncommon = new ArrayList<String>();
        
        Iterator it1 = strAWordMap.entrySet().iterator();
        while (it1.hasNext()) {
            Map.Entry pair = (Map.Entry)it1.next();
            String word = (String)pair.getKey();
            int count = (int)pair.getValue();
            if(count == 1 && !strBWordMap.containsKey(word)) {
                uncommon.add(word);
            }

            // System.out.println(pair.getKey() + " = " + pair.getValue());
            it1.remove(); // avoids a ConcurrentModificationException
        }

        Iterator it2 = strBWordMap.entrySet().iterator();
        while (it2.hasNext()) {
            Map.Entry pair = (Map.Entry)it2.next();
            String word = (String)pair.getKey();
            int count = (int)pair.getValue();
            if(count == 1 && !strAWordMap.containsKey(word)) {
                uncommon.add(word);
            }

            // System.out.println(pair.getKey() + " = " + pair.getValue());
            it2.remove(); // avoids a ConcurrentModificationException
        }

        // String[] bar = uncommon.toArray(new String[uncommon.size()]);
        return uncommon.toArray(new String[uncommon.size()]);
    }
}
