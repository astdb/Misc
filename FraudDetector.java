// This program takes in a list of credit card transactions and outputs a list of credit card numbers 
// associated with suspected fraudulent transactions.

import java.util.*;
import java.io.*;
import java.text.*;

public class FraudDetector {
    public static void main(String[] args) {
        if( args.length <= 0 || args[0].trim().length() <= 0 ){
            System.out.println("Usage: $> java FraudDetector inputfile");
            System.exit(1);
        }

        Long priceThreshold = 10000L;       // in cents
        String date = "2017-09-12";         // in "yyyy-MM-dd" format
        
    }

    // filterTransactions function takes in a list of transactions, a date, and a threshold spend amount 
    // and returns a list of credit card numbers associated with transactions exceeding the threshold.
    public static ArrayList<String> filterTransactions(String inputFile, String date, long amountThreshold) {
        // read input file
        String inputFile = args[0];
        Scanner input = null;

        try {
            input = new Scanner(new File(inputFile));
        } catch( FileNotFoundException e ) {
            System.out.println("No such file: " + inputFile);
            return null;
        }

        // declare map to collect creditcards and expense totals
        // this collection will be a set of creditcard numbers mapping to a set of maps, which in turn map dates to totals
        HashMap<String, HashMap<Long, Long>> creditCardTotals = new HashMap<String, HashMap<Long, Long>>();

        // list of sus[icious cards
        ArrayList<String> suspectCards = new ArrayList<String>();
        
        // iterate through the transaction input file
        // e.g. 10d7ce2f43e35fa57d1bbf8b1e2, 2014-04-29T13:15:54, 10.00
        while( input.hasNextLine() ) {
            String transaction = input.nextLine();

            // split into components
            String[] transactionComponents = transaction.split(",");

            if(transactionComponents.length < 3) {
                // malformed input, move to next
                System.out.println("Error parsing transaction: " + transaction);
                continue;
            }

            String transactionCardHash = transactionComponents[0].trim();
            String transactionDate = transactionComponents[1].trim();
            String transactionAmount = transactionComponents[2].trim();

            // extract day component from input e.g. "2014-04-29" from "2014-04-29T13:15:54"
            String[] dateSplit = transactionDate.split("T");
            String dayComp = dateSplit[0];
            // TODO: validate day component

            // get timestamp for the date
            Long time = null;
            try {
                DateFormat dateFormat = new SimpleDateFormat("yyyy-MM-dd");        
                Date date = dateFormat.parse(dayComp);
                time = new Long(date.getTime());
                
            } catch(ParseException e) {
                // malformed input, move to next
                System.out.println("Error parsing date: " + dayComp);
                continue;
            }

            Long amt = new Long((long)(Float.parseFloat(transactionAmount)*100));

            // check if a record exists for this credit card
            HashMap<Long, Long> cardExists = creditCardTotals.get(transactionCardHash);

            if( cardExists != null ) {
                // card known in this dataset before - check if a total eists for this date and update it

                // check if a total exists for this card for this date
                Long totalExists = cardExists.get(time);                

                if(totalExists != null) {
                    // a count is known for this card for this date - update                   
                    // Long newTot = totalExists + amt;

                    cardExists.put(time, cardExists.get(time) + amt);
                } else {
                    // no count for this card for this date, add
                    cardExists.put(time, amt);
                }

            } else {
                // new card to this dataset - create record
                HashMap<Long, Long> thisAmount = new HashMap<Long, Long>();
                thisAmount.put(time, amt);
                creditCardTotals.put(transactionCardHash, thisAmount);
            }
        }

        return suspectCards;
    }
}
