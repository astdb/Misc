
// CalcPrintCosts accepts a list of print jobs from an input file and calculates the cost of each job according to a given price list. each job has a total number of pages, number of colour pages and whether printing is double sided. 

import java.util.*;
import java.io.*;

public class CalcPrintCosts {
    public static void main(String[] args) {
        // read print jobs input file name from command line 
        String filename = readInputfileName(args);

        // read input data from input file
        Scanner fileinput = null;
        try {
			fileinput = new Scanner(new File(filename));
		} catch( FileNotFoundException e ) {
			System.out.println("No file found called <" + filename + ">");
		}

        // total cost of all print jobs
        long totalPrintCost = 0;

        // iterate through the print jobs
        while( fileinput.hasNextLine() ) {
            // read print job
            String this_row = fileinput.nextLine().trim();

            if(validInputJob(this_row)){
                // split job string into individual components
                String[] currentPrintJob = this_row.split(",");
                long totalPages = Long.parseLong(currentPrintJob[0].trim());
                long totalColor = Long.parseLong(currentPrintJob[1].trim());
                String sidedNess = currentPrintJob[2].trim();

                // create and print job object
                PrintJob printJob = new PrintJob(totalPages, totalColor, doubleSided(sidedNess));
                printJob.print();

                // add cost to total print cost
                totalPrintCost += printJob.cost();
            }
                      
        }

        // print total cost
        System.out.println(formattedPrice(totalPrintCost));
        return;        
    }

    // ---------------------------------- end main() function ----------------------------------

    // read input file name from command line 
    private static String readInputfileName(String[] inputArgs) {
        if( inputArgs.length <= 0 || inputArgs[0].trim().length() <= 0 ) {
            System.out.println("Please enter an input file name on command-line (i.e. $> java CalcPrintCosts input.csv)");
            System.exit(1);
        }

        return inputArgs[0].trim();
    }

    // takes a string representing an unprocessed row from  the input file and determine if the data is valid
    private static boolean validInputJob(String inputJob) {
        // ---------- check number of fields
        String[] currentPrintJobRaw = inputJob.split(",");

        if(currentPrintJobRaw.length != 3) {
            // valid job row required to have three fields i.e. [total pages, color pages, single/double-sided]
            System.out.println("Row doesn't split into three components");
            return false;
        }

        // --------- check field types
        // first two fields must contain numeric values (representing total and colored page numbers)
        long total_pages;
        long total_color;

        try {
            total_pages = Long.parseLong(currentPrintJobRaw[0].trim());
            total_color = Long.parseLong(currentPrintJobRaw[1].trim());

        } catch (NumberFormatException e) {
            System.out.println("Exception formatting page numbers into longs");
            return false;
        }


        // --------- check page counts
        // page numbers must be positive and number of color pages must be same or less than total pages
        if( (total_pages <= 0) || (total_color < 0) || (total_color > total_pages) ) {
            System.out.println("Page numbers mismatch");
            return false;
        }

        // --------- check sided-ness indicator - must be 'true' or 'false', literally
        String sided_ness = currentPrintJobRaw[2].trim();

        if( !(sided_ness.equalsIgnoreCase("true") || sided_ness.equalsIgnoreCase("false")) ) {
            System.out.println("Sided-ness error");
            return false;
        }

        return true;
    }

    // function to pretty-print a cent amount in $xx.cc format
    private static String formattedPrice(long priceInCents) {
        long dollars = priceInCents/100;
        long cents = priceInCents - (dollars * 100);
        String formattedCents = String.format("%02d", cents);

        return "$" + dollars + "." + formattedCents;
    }

    // function to return boolean value for doublesided 'true'/'false' string from inputfile
    private static boolean doubleSided(String sidedness) {
        if(sidedness.equalsIgnoreCase("true")) {
            return true;
        }

        return false;
    }

    // class representing a print job input from file
    private static class PrintJob {
        // properties of printjob
        private long totalPages;
        private long totalColorPages;
        private boolean doubleSided;

        // prices for print job types (in cents ¢) - prices are tracked in cents to alleviate floating-point errors
        public static final long A4_NONCOLOR_SINGLESIDED = 15;
        public static final long A4_COLOR_SINGLESIDED = 25;
        public static final long A4_NONCOLOR_DOUBLESIDED = 10;
        public static final long A4_COLOR_DOUBLESIDED = 20;

        protected PrintJob(long totalPages, long totalColorPages, boolean doubleSided){
            this.totalPages = totalPages;
            this.totalColorPages = totalColorPages;
            this.doubleSided = doubleSided;
        }

        // calculate cost of job
        protected long cost() {
            long total_noncolor = 0;
            long total_color = 0;

            // if double-sided
            if(this.doubleSided) {
                total_noncolor = (totalPages-totalColorPages) * A4_NONCOLOR_DOUBLESIDED;
                total_color = totalColorPages * A4_COLOR_DOUBLESIDED;

                return (total_noncolor + total_color);
            }

            // if single-sided
            total_noncolor = (totalPages-totalColorPages) * A4_NONCOLOR_SINGLESIDED;
            total_color = totalColorPages * A4_COLOR_SINGLESIDED;

            return (total_noncolor + total_color);
        }

        // pretty-print job to console
        protected void print() {
            System.out.println("Total pages:\t" + this.totalPages);
            System.out.println("Total color pages:\t" + this.totalColorPages);

            if(this.doubleSided) {
                System.out.println("Double-sided:\tYes");
            } else {
                System.out.println("Double-sided:\tNo");
            }

            System.out.println("Cost:\t" + formattedPrice(this.cost()));

            System.out.println("--------------------\n");     
        }
    }
}
