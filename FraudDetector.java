
// This program takes in a list of credit card transactions and outputs a list of credit card numbers 
// associated with suspected fraudulent transactions according to given guidelines.

import java.util.*;
import java.io.*;
import java.text.*;
import java.util.regex.*;

public class FraudDetector {
    public static void main(String[] args) {
        if( args.length <= 0 || args[0].trim().length() <= 0 ){
            System.out.println("Usage: $> java FraudDetector inputfile");
            return;
        }

        String inputFile = args[0];         // read transactions input file from command line
        Long priceThreshold = 11000L;       // in cents
        String date = "2017-09-14";         // in "yyyy-MM-dd" format
        

        // ------------------------- Test 01 -------------------------
        // a single card, with one exceeding transaction
        ArrayList<String> expected_test01 = new ArrayList<String>();
        expected_test01.add("10d7ce2f43e35fa57d1bbf8b1e2");

        if(equalLists(expected_test01, filterTransactions("test01.file", date, priceThreshold))) {
            System.out.println("Test01: Pass");
        } else {
            System.out.println("Test01: Fail");
        }

        // ------------------------- Test 02 -------------------------
        // a single card, with one non-exceeding transaction
        ArrayList<String> expected_test02 = new ArrayList<String>();
        
        if(equalLists(expected_test02, filterTransactions("test02.file", date, priceThreshold))) {
            System.out.println("Test02: Pass");
        } else {
            System.out.println("Test02: Fail");
        }

        // ------------------------- Test 03 -------------------------
        // an empty input
        ArrayList<String> expected_test03 = new ArrayList<String>();
        
        if(equalLists(expected_test03, filterTransactions("test03.file", date, priceThreshold))) {
            System.out.println("Test03: Pass");
        } else {
            System.out.println("Test03: Fail");
        }

        // ------------------------- Test 04 -------------------------
        // a garbage input
        ArrayList<String> expected_test04 = new ArrayList<String>();
        
        if(equalLists(expected_test04, filterTransactions("test04.file", date, priceThreshold))) {
            System.out.println("Test04: Pass");
        } else {
            System.out.println("Test04: Fail");
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
        }

        // ------------------------- Test 06 -------------------------
        // Two cards, both exceeding limit for different days
        ArrayList<String> expected_test06 = new ArrayList<String>();
        expected_test06.add("10d7ce2f43e35fa57d1bbf8b1e2");
        
        if(equalLists(expected_test06, filterTransactions("test06.file", date, priceThreshold))) {
            System.out.println("Test06: Pass");
        } else {
            System.out.println("Test06: Fail");
        }

        // print out suspected card hashes
        ArrayList<String> suspect_cards_list = filterTransactions("test06.file", date, priceThreshold);
        if(suspect_cards_list.size() == 0) {
            System.out.println("No cards with fraudulent transactions found.");
            return;
        } else {
            for (String cardHash: suspect_cards_list) {
                System.out.println(cardHash);
                return;
            }
        }

        // ------------------------- Test 07 -------------------------
        // Non-existent input file
        ArrayList<String> expected_test07 = null;
        
        if(equalLists(expected_test07, filterTransactions("test07.file", date, priceThreshold))) {
            System.out.println("Test07: Pass");
        } else {
            System.out.println("Test07: Fail");
        }

        // ------------------------- Test 08 -------------------------
        // Two cards, multiple transcations, single day, both exceeding
        ArrayList<String> expected_test08 = new ArrayList<String>();
        expected_test08.add("10d7ce2f43e35fa57d1bbf8b1e2");
        expected_test08.add("10d7ce2f43e35fa57d1bbf8b1e3");        
        
        if(equalLists(expected_test08, filterTransactions("test08.file", date, priceThreshold))) {
            System.out.println("Test08: Pass");
        } else {
            System.out.println("Test08: Fail");
        }

        // ------------------------- Test 09 -------------------------
        // Two cards, multiple transcations, single day, both non-exceeding
        ArrayList<String> expected_test09 = new ArrayList<String>();        
        
        if(equalLists(expected_test09, filterTransactions("test09.file", date, priceThreshold))) {
            System.out.println("Test09: Pass");
        } else {
            System.out.println("Test09: Fail");
        }

        // ------------------------- Test 10 -------------------------
        // Two cards, multiple transcations, single day, both non-exceeding
        ArrayList<String> expected_test10 = new ArrayList<String>();
        expected_test08.add("10d7ce2f43e35fa57d1bbf8b1e4");
        
        if(equalLists(expected_test10, filterTransactions("test10.file", date, priceThreshold))) {
            System.out.println("Test10: Pass");
        } else {
            System.out.println("Test10: Fail");
        }
        
        // print out suspected card hashes
        // ArrayList<String> suspect_cards_list = filterTransactions("test05.file", date, priceThreshold);
        // if(suspect_cards_list.size() == 0) {
        //     System.out.println("No cards with fraudulent transactions found.");
        //     return;
        // } else {
        //     for (String cardHash: suspect_cards_list) {
        //         System.out.println(cardHash);
        //         return;
        //     }
        // }
    }

    // filterTransactions function takes in a list of transactions, a date, and a threshold spend amount and returns
    // a list of credit card number hashes associated with a total of transactions exceeding the threshold for that day.
    public static ArrayList<String> filterTransactions(String inputFile, String dateThreshold, long amountThreshold) {
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
            System.out.println("Invalid date threshold parameter.");
            return null;
        }

        // declare map structure to collect creditcards and expense totals
        // this collection will be a set of creditcard numbers mapping to a set of maps, which in turn map dates to totals
        // e.g 
        //  card1 -> {date1 -> total1, date2 -> total2}
        //  card2 -> {date3 -> total3}
        // HashMap<String, HashMap<Long, Long>> creditCardTotals = new HashMap<String, HashMap<Long, Long>>();
        HashMap<String, Long> creditCardTotals = new HashMap<String, Long>();

        // list of suspicious cards - this will be populated whenever a card's total for a given day is detected to be over threshold
        ArrayList<String> suspectCards = new ArrayList<String>();

        // get timestamp for the threshold date (the date the fraudulent transaction cards are required for)
        // Long dateThreshold = null;
        // try {
        //     DateFormat dateFormat = new SimpleDateFormat("yyyy-MM-dd");        
        //     Date dte = dateFormat.parse(date);
        //     dateThreshold = new Long(dte.getTime());
            
        // } catch(ParseException e) {
        //     // malformed input, move to next transaction
        //     // System.out.println("Error parsing date: " + dayComp);
        //     return null;
        // }
        
        // iterate through each of the transaction input file
        while(input.hasNextLine()) {
            String transaction = input.nextLine();      // e.g. "10d7ce2f43e35fa57d1bbf8b1e2, 2014-04-29T13:15:54, 10.00"

            // ignore if line starts with # (comment)
            Pattern commentLine = Pattern.compile("^#");
            Matcher clm = commentLine.matcher(transaction.trim());
            if (clm.find()) {
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

            // extract day component from input e.g. "2014-04-29" from "2014-04-29T13:15:54"
            String[] dateSplit = transactionDate.split("T");
            String transactionDate_day = dateSplit[0];

            // ensure that the date provided is of "yyyy-MM-dd" format
            transactionDate_day = transactionDate_day.trim();
            Matcher pdm = properDate.matcher(transactionDate_day);
            if (!pdm.find()) {
                // System.out.println("Invalid date threshold parameter.");
                // return null;

                // possibly malformed transaction reord - move onto next
                continue;
            }

            // get timestamp for the transaction date (day component)
            // Long day = null;
            // try {
            //     DateFormat dateFormat = new SimpleDateFormat("yyyy-MM-dd");        
            //     Date dte = dateFormat.parse(dayComp);
            //     day = new Long(dte.getTime());
                
            // } catch(ParseException e) {
            //     // malformed input, move to next transaction
            //     // System.out.println("Error parsing date: " + dayComp);
            //     continue;
            // }

            // check if this transaction is one we'd be interested in - i.e. one for the date given
            // TODO - begs the question, can't you just string compare the day components?
            // Also, if the processor only sees transactions only for a given day, can't it be a simple hashmap of card hash-> total?
            if(transactionDate_day != dateThreshold) {
                continue;
            }

            Long amt = null;
            try {
                amt = new Long((long)(Float.parseFloat(transactionAmount)*100));
            } catch(NumberFormatException e) {
                // malformed input - continue to next
                continue;
            }            

            // check if a record exists for this credit card
            // HashMap<Long, Long> cardExists = creditCardTotals.get(transactionCardHash);
            Long cardTotal = creditCardTotals.get(transactionCardHash);

            if(cardTotal != null) {
                // card known in this dataset before - update total
                Long newTotal = cardTotal + amt;

                if(newTotal > amountThreshold) {
                    if(!suspectCards.contains(transactionCardHash)) {
                        suspectCards.add(transactionCardHash);
                    }
                    
                }

                // // check if a total exists for this card for this date
                // Long totalExists = cardExists.get(day);

                // if(totalExists != null) {
                //     // a count is known for this card for this date - update                   
                //     Long newTot = cardExists.get(day) + amt;

                //     if(newTot > amountThreshold) {
                //         if(!suspectCards.contains(transactionCardHash)) {
                //             suspectCards.add(transactionCardHash);
                //         }
                //     }
                //     cardExists.put(day, newTot);
                // } else {
                //     // no count for this card for this date, add
                //     cardExists.put(day, amt);
                // }

            } else {
                // new card to this dataset - create record
                creditCardTotals.put(transactionCardHash, amt);

                if(amt > amountThreshold) {
                    if(!suspectCards.contains(transactionCardHash)) {
                        suspectCards.add(transactionCardHash);
                    }
                }

                // HashMap<Long, Long> thisAmount = new HashMap<Long, Long>();
                // thisAmount.put(day, amt);                

                // creditCardTotals.put(transactionCardHash, thisAmount);
            }
        }

        return suspectCards;
    }

    // utility test function to check if two given lists of card number hashes contain the same elements
    public static boolean equalLists(ArrayList<String> expected, ArrayList<String> result) {
        // DEBUG - print out the two lists
        System.out.println("\tExpected: " + expected);
        System.out.println("\tResults: " + result);

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
