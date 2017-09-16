
// This program takes in a list of credit card transactions and outputs a list of credit card numbers 
// associated with suspected fraudulent transactions according to given guidelines.

import java.util.*;
import java.io.*;
import java.text.*;
import java.util.regex.*;

public class FraudDetector {
    public static void main(String[] args) {
        // if( args.length <= 0 || args[0].trim().length() <= 0 ){
        //     System.out.println("Usage: $> java FraudDetector inputfile");
        //     return;
        // }

        // String inputFile = args[0];         // read transactions input file from command line
        Long priceThreshold = 11000L;       // in cents (set here to $110 for test purposes)
        String date = "2017-09-14";         // in "yyyy-MM-dd" format (for test purposes)
        

        // ------------------------- Test 01 -------------------------
        // a single card, with one exceeding transaction
        ArrayList<String> expected_test01 = new ArrayList<String>();
        expected_test01.add("10d7ce2f43e35fa57d1bbf8b1e2");

        if(equalLists(expected_test01, filterTransactions("test01.file", date, priceThreshold))) {
            System.out.println("Test01: Pass");
        } else {
            System.out.println("Test01: Fail");
            return;
        }

        // // print out suspected card hashes
        // ArrayList<String> suspect_cards_list = filterTransactions("test01.file", date, priceThreshold);
        // if(suspect_cards_list.size() == 0) {
        //     System.out.println("No cards with fraudulent transactions found.");
        //     return;
        // } else {
        //     for (String cardHash: suspect_cards_list) {
        //         System.out.println(cardHash);
        //         return;
        //     }
        // }

        // ------------------------- Test 02 -------------------------
        // a single card, with one non-exceeding transaction
        ArrayList<String> expected_test02 = new ArrayList<String>();
        
        if(equalLists(expected_test02, filterTransactions("test02.file", date, priceThreshold))) {
            System.out.println("Test02: Pass");
        } else {
            System.out.println("Test02: Fail");
            return;
        }

        // ------------------------- Test 03 -------------------------
        // an empty input
        ArrayList<String> expected_test03 = new ArrayList<String>();
        
        if(equalLists(expected_test03, filterTransactions("test03.file", date, priceThreshold))) {
            System.out.println("Test03: Pass");
        } else {
            System.out.println("Test03: Fail");
            return;
        }

        // ------------------------- Test 04 -------------------------
        // a garbage input
        ArrayList<String> expected_test04 = new ArrayList<String>();
        
        if(equalLists(expected_test04, filterTransactions("test04.file", date, priceThreshold))) {
            System.out.println("Test04: Pass");
        } else {
            System.out.println("Test04: Fail");
            return;
        }

        // ------------------------- Test 05 -------------------------
        // Two cards, both exceeding limit for the same day
        ArrayList<String> expected_test05 = new ArrayList<String>();
        expected_test05.add("10d7ce2f43e35fa57d1bbf8b1e2");
        expected_test05.add("10d7ce2f43e35fa57d1bbf8b1e3");
        
        if(equalLists(expected_test05, filterTransactions("test05.file", date, priceThreshold))) {
            System.out.println("Test05: Pass");
        } else {
            System.out.println("Test05: Fail");
            return;
        }

        // ------------------------- Test 06 -------------------------
        // Two cards, both exceeding limit for different days
        ArrayList<String> expected_test06 = new ArrayList<String>();
        expected_test06.add("10d7ce2f43e35fa57d1bbf8b1e2");
        
        if(equalLists(expected_test06, filterTransactions("test06.file", date, priceThreshold))) {
            System.out.println("Test06: Pass");
        } else {
            System.out.println("Test06: Fail");
            return;
        }

        // ------------------------- Test 07 -------------------------
        // Non-existent input file
        ArrayList<String> expected_test07 = null;
        
        if(equalLists(expected_test07, filterTransactions("test07.file", date, priceThreshold))) {
            System.out.println("Test07: Pass");
        } else {
            System.out.println("Test07: Fail");
            return;
        }

        // ------------------------- Test 08 -------------------------
        // Two cards, multiple transactions, single day, both exceeding
        ArrayList<String> expected_test08 = new ArrayList<String>();
        expected_test08.add("10d7ce2f43e35fa57d1bbf8b1e2");
        expected_test08.add("10d7ce2f43e35fa57d1bbf8b1e3");        
        
        if(equalLists(expected_test08, filterTransactions("test08.file", date, priceThreshold))) {
            System.out.println("Test08: Pass");
        } else {
            System.out.println("Test08: Fail");
            return;
        }

        // ------------------------- Test 09 -------------------------
        // Two cards, multiple transactions, single day, both non-exceeding
        ArrayList<String> expected_test09 = new ArrayList<String>();        
        
        if(equalLists(expected_test09, filterTransactions("test09.file", date, priceThreshold))) {
            System.out.println("Test09: Pass");
        } else {
            System.out.println("Test09: Fail");
            return;
        }

        // ------------------------- Test 10 -------------------------
        // Three cards, multiple transactions, single day, one exceeding
        ArrayList<String> expected_test10 = new ArrayList<String>();
        expected_test10.add("10d7ce2f43e35fa57d1bbf8b1e4");
        
        if(equalLists(expected_test10, filterTransactions("test10.file", date, priceThreshold))) {
            System.out.println("Test10: Pass");
        } else {
            System.out.println("Test10: Fail");
            return;
        }
    }

    // filterTransactions function takes in a list of transactions, a date, and a threshold spend amount and returns
    // a list of credit card number hashes associated with a total of transactions exceeding the threshold for that day.
    public static ArrayList<String> filterTransactions(String inputFile, String dateThreshold, long amountThreshold) {
        // System.out.println("\nfilterTransactions() ------------------------------------------------------------ ");
        // read input file
        Scanner input = null;

        try {
            input = new Scanner(new File(inputFile));
        } catch( FileNotFoundException e ) {
            // System.out.println("No such file: " + inputFile);
            return null;
        }
        
        if(inputFile == null || dateThreshold == null) {
            return null;
        }

        // ensure that the date provided is of "yyyy-MM-dd" format
        dateThreshold = dateThreshold.trim();
        Pattern properDate = Pattern.compile("^[0-9][0-9][0-9][0-9]-[0-9][0-9]-[0-9][0-9]$");
        Matcher m = properDate.matcher(dateThreshold);
        if (!m.find()) {
            // System.out.println("Invalid date threshold parameter.");
            return null;
        }

        // declare map to collect creditcards and expense totals (mapping card hashes to transaction totals)
        HashMap<String, Long> creditCardTotals = new HashMap<String, Long>();

        // list of suspicious cards - this will be populated whenever a card's total for the given day is detected to be over threshold
        ArrayList<String> suspectCards = new ArrayList<String>();
        
        // iterate through each of the transaction input file
        while(input.hasNextLine()) {
            // System.out.println("filterTransactions() ----------------- ");
            String transaction = input.nextLine();      // e.g. "10d7ce2f43e35fa57d1bbf8b1e2, 2014-04-29T13:15:54, 10.00"
            // System.out.println("filterTransactions() - processing transaction record <" + transaction + ">");

            // ignore if line starts with # (comment)
            Pattern commentLine = Pattern.compile("^#");
            Matcher clm = commentLine.matcher(transaction.trim());
            if (clm.find()) {
                // System.out.println("filterTransactions() - comment, continuing to next..");
                continue;
            }

            // split transaction string into components (card hash, date, amount)
            String[] transactionComponents = transaction.split(",");

            // confirm transaction is complete
            if(transactionComponents.length < 3) {
                // malformed input, move to next transaction
                // System.out.println("Error parsing transaction: " + transaction);
                continue;
            }

            // components of this transaction
            String transactionCardHash = transactionComponents[0].trim();
            String transactionDate = transactionComponents[1].trim();
            String transactionAmount = transactionComponents[2].trim();
            // System.out.println("filterTransactions() - split into components: Hash: " + transactionCardHash + ", Date: " + transactionDate + ", Amount: " + transactionAmount);

            // extract day component from input e.g. "2014-04-29" from "2014-04-29T13:15:54"
            String[] dateSplit = transactionDate.split("T");
            String transactionDate_day = dateSplit[0];
            // System.out.println("filterTransactions() - Tx day component: " + transactionDate_day);

            // ensure that the date provided is of "yyyy-MM-dd" format
            transactionDate_day = transactionDate_day.trim();
            Matcher pdm = properDate.matcher(transactionDate_day);
            if (!pdm.find()) {
                // System.out.println("Invalid date threshold parameter.");
                // return null;

                // System.out.println("filterTransactions() - invalid date(" + transactionDate_day + "), continuing to next..");

                // possibly malformed transaction reord - move onto next
                continue;
            }

            // check if this transaction is one we'd be interested in - i.e. one for the date given
            // TODO - begs the question, can't you just string compare the day components?
            // Also, if the processor only sees transactions only for a given day, can't it be a simple hashmap of card hash-> total?
            if(!transactionDate_day.equals(dateThreshold)) {
                // System.out.println("filterTransactions() - tx date (" + transactionDate_day + ") not equal to threshold date (" + dateThreshold + "), continuing to next..");
                continue;
            }

            // System.out.println("filterTransactions() - dates match");

            Long amt = null;
            try {
                amt = new Long((long)(Float.parseFloat(transactionAmount)*100));
                // System.out.println("filterTransactions() - numeric transaction amount: " + amt);
            } catch(NumberFormatException e) {
                // malformed input - continue to next
                // System.out.println("filterTransactions() - error formatting numeric tx amt - continuing.. ");
                continue;
            }            

            // check if a record exists for this credit card
            // HashMap<Long, Long> cardExists = creditCardTotals.get(transactionCardHash);
            Long cardTotal = creditCardTotals.get(transactionCardHash);

            if(cardTotal != null) {
                // card known in this dataset before - update total
                // System.out.println("filterTransactions() - card (" + transactionCardHash + ") known in this dataset - updating");
                Long newTotal = cardTotal + amt;
                creditCardTotals.put(transactionCardHash, newTotal);

                if(newTotal > amountThreshold) {
                    // System.out.println("filterTransactions() - total (" + newTotal + ") exceeds threshold (" + amountThreshold + ")");
                    if(!suspectCards.contains(transactionCardHash)) {
                        // System.out.println("filterTransactions() - not known in suspected cards collection, adding..");
                        suspectCards.add(transactionCardHash);
                    }
                    
                } else {
                    // System.out.println("filterTransactions() - total (" + newTotal + ") below threshold (" + amountThreshold + ")");
                }
            } else {
                // System.out.println("filterTransactions() - card (" + transactionCardHash + ") NOT known in this dataset - creating..");
                // new card to this dataset - create record
                creditCardTotals.put(transactionCardHash, amt);

                if(amt > amountThreshold) {
                    // System.out.println("filterTransactions() - amount(" + amt + ") exceeds threshold (" + amountThreshold + ")");
                    if(!suspectCards.contains(transactionCardHash)) {
                        // System.out.println("filterTransactions() - not known in suspected cards collection, adding..");
                        suspectCards.add(transactionCardHash);
                    }
                } else {
                    // System.out.println("filterTransactions() - total (" + amt + ") below threshold (" + amountThreshold + ")");
                }
            }
        }

        return suspectCards;
    }

    // utility test function to check if two given lists of card number hashes contain the same elements
    public static boolean equalLists(ArrayList<String> expected, ArrayList<String> result) {
        // DEBUG - print out the two lists
        // System.out.println("\tExpected: " + expected);
        // System.out.println("\tResults: " + result);

        if (expected == null && result == null) {
            return true;
        }

        if ((expected == null && result!= null) || (expected != null && result== null) || (expected.size() != result.size())) {
            return false;
        }

        // Sort and compare the two lists          
        Collections.sort(expected);
        Collections.sort(result);

        return expected.equals(result);
    }
}
