
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
        String date = "2014-04-29";         // in "yyyy-MM-dd" format

        // call filterTransactions() with transactions, price threchold and date to get a list of cards w/ suspicious transactions
        ArrayList<String> suspect_cards_list = filterTransactions(inputFile, date, priceThreshold);

        // print out suspected card hashes
        if(suspect_cards_list.size() == 0) {
            System.out.println("No cards with fraudulent transactions found.");
        } else {
            for (String cardHash: suspect_cards_list) {
                System.out.println(cardHash);
            }
        }        
    }

    // filterTransactions function takes in a list of transactions, a date, and a threshold spend amount and returns
    // a list of credit card number hashes associated with a total of transactions exceeding the threshold for that day.
    public static ArrayList<String> filterTransactions(String inputFile, String date, long amountThreshold) {
        // read input file
        Scanner input = null;

        try {
            input = new Scanner(new File(inputFile));
        } catch( FileNotFoundException e ) {
            System.out.println("No such file: " + inputFile);
            return null;
        }

        // declare map structure to collect creditcards and expense totals
        // this collection will be a set of creditcard numbers mapping to a set of maps, which in turn map dates to totals
        // e.g 
        //  card1 -> {date1 -> total1, date2 -> total2}
        //  card2 -> {date3 -> total3}
        HashMap<String, HashMap<Long, Long>> creditCardTotals = new HashMap<String, HashMap<Long, Long>>();

        // list of suspicious cards - this will be populated whenever a card's total for a given day is detected to be over threshold
        ArrayList<String> suspectCards = new ArrayList<String>();
        
        // iterate through each of the transaction input file
        while( input.hasNextLine() ) {
            String transaction = input.nextLine();      // e.g. "10d7ce2f43e35fa57d1bbf8b1e2, 2014-04-29T13:15:54, 10.00"

            // ignore if line starts with # (comment)
            Pattern commentLine = Pattern.compile("^#");
            Matcher m = commentLine.matcher(transaction.trim());
            if (m.find()) {
                continue;
            }

            // split transaction string into components (card hash, date, amount)
            String[] transactionComponents = transaction.split(",");

            // confirm transaction is complete
            if(transactionComponents.length < 3) {
                // malformed input, move to next
                System.out.println("Error parsing transaction: " + transaction);
                continue;
            }

            // components of this transaction
            String transactionCardHash = transactionComponents[0].trim();
            String transactionDate = transactionComponents[1].trim();
            String transactionAmount = transactionComponents[2].trim();

            // extract day component from input e.g. "2014-04-29" from "2014-04-29T13:15:54"
            String[] dateSplit = transactionDate.split("T");
            String dayComp = dateSplit[0];

            // get timestamp for the date (day component)
            Long day = null;
            try {
                DateFormat dateFormat = new SimpleDateFormat("yyyy-MM-dd");        
                Date dte = dateFormat.parse(dayComp);
                day = new Long(dte.getTime());
                
            } catch(ParseException e) {
                // malformed input, move to next
                System.out.println("Error parsing date: " + dayComp);
                continue;
            }

            Long amt = new Long((long)(Float.parseFloat(transactionAmount)*100));

            // check if a record exists for this credit card
            HashMap<Long, Long> cardExists = creditCardTotals.get(transactionCardHash);

            if( cardExists != null ) {
                // card known in this dataset before - check if a total eists for this date

                // check if a total exists for this card for this date
                Long totalExists = cardExists.get(day);                

                if(totalExists != null) {
                    // a count is known for this card for this date - update                   
                    Long newTot = cardExists.get(day) + amt;

                    if(newTot > amountThreshold) {
                        if(!suspectCards.contains(transactionCardHash)){
                            suspectCards.add(transactionCardHash);
                        }
                    }
                    cardExists.put(day, newTot);
                } else {
                    // no count for this card for this date, add
                    cardExists.put(day, amt);
                }

            } else {
                // new card to this dataset - create record
                if(amt > amountThreshold) {
                    if(!suspectCards.contains(transactionCardHash)){
                        suspectCards.add(transactionCardHash);
                    }
                }

                HashMap<Long, Long> thisAmount = new HashMap<Long, Long>();
                thisAmount.put(day, amt);                

                creditCardTotals.put(transactionCardHash, thisAmount);
            }
        }

        return suspectCards;
    }

    // utility test function to check if two given lists of card number hashes contain the same elements
    public  boolean equalLists(ArrayList<String> expected, ArrayList<String> result) {
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
