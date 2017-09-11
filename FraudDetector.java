
// This program takes in a list of credit card transactions and outputs a list of credit card numbers 
// associated with suspected fraudulent transactions.

import java.util.*;
import java.io.*;

public class FraudDetector {
    public static void main(String[] args) {
        if( args.length <= 0 || args[0].trim().length() <= 0 ){
            System.out.println("Usage: $> java FraudDetector inputfile");
            System.exit(1);
        }

        // read input file
        String inputFile = args[1];
        Scanner input = null;

        try {
            input = new Scanner(new File(inputFile));
        } catch( FileNotFoundException e ) {
            System.out.println("No such file: " + inputFile);
        }

        // declare map to collect creditcards and expense totals
        // collection will be a set of creditcard numbers mapping to a set of maps, which in turn map dates to totals
        HashMap<String, HashMap<Integer, Integer>> creditCardTotals = new HashMap<String, HashMap<Integer, Integer>>();
        
        // iterate through the transaction input file
        // e.g. 10d7ce2f43e35fa57d1bbf8b1e2, 2014-04-29T13:15:54, 10.00
        while( input.hasNextLine() ) {
            String transaction = input.nextLine();

            // split into components
            String[] transactionComponents = transaction.split(",");

            if(transactionComponents.length < 3) {
                // malformed transaction input, move to next
                continue;
            }

            HashMap<Integer, Integer> cardExists = map.get(transactionComponents[0].trim());

            if( cardExists != null ) {
                // card known in this dataset before

            } else {
                // new card to this dataset

            }
        }

        // create list of suspicious credit card numbers

        // return

    }

    // filterTransactions function takes in a list of transactions, a date and a threshold spend amount 
    // and returns a list of credit card numbers associated with transactions exceeding the threshold.
    public static void filterTransactions(String inputFile, Date date, long amountThreshold) {
        
    }
}
