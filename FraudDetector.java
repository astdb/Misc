
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

        // read input file
        String inputFile = args[0];
        Scanner input = null;

        try {
            input = new Scanner(new File(inputFile));
        } catch( FileNotFoundException e ) {
            System.out.println("No such file: " + inputFile);
        }

        // declare map to collect creditcards and expense totals
        // this collection will be a set of creditcard numbers mapping to a set of maps, which in turn map dates to totals
        HashMap<String, HashMap<Integer, Long>> creditCardTotals = new HashMap<String, HashMap<Integer, Long>>();
        
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

            // check if a record exists for this credit card
            HashMap<Integer, Integer> cardExists = creditCardTotals.get(transactionCardHash);

            if( cardExists != null ) {
                // card known in this dataset before - check if a total eists for this date and update it
                
                // extract day component from input e.g. "2014-04-29" from "2014-04-29T13:15:54"
                String dayComp = transactionDate.split('T');

                // calculate timestamp for the date
                try {
                    DateFormat dateFormat = new SimpleDateFormat("yyyy-MM-dd");        
                    Date date = dateFormat.parse(dayComp);
                    long time = date.getTime();
                    // new Timestamp(time);
                    // System.out.println(time);
                } catch(ParseException e) {
                    // malformed input, move to next
                    System.out.println("Error parsing date: " + dayComp);
                    continue;
                }

            } else {
                // new card to this dataset - create record

            }
        }

        // create list of suspicious cards

        // return

    }

    // filterTransactions function takes in a list of transactions, a date and a threshold spend amount 
    // and returns a list of credit card numbers associated with transactions exceeding the threshold.
    public static void filterTransactions(String inputFile, Date date, long amountThreshold) {
        
    }
}
