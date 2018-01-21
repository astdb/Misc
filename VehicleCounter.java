
import java.io.*;
import java.util.*;

public class VehicleCounter {
    public static void main(String[] args) {
        // read input data file from command line and process options
        if( args.length <= 0 || args[0].trim().length() <= 0 ) {
            System.out.println("Usage: $> java VehicleCounter data.txt");
            return;
        }

        String dataFileName = args[0].trim();

        // iterate through counter data and process required metric
        Scanner vehicleData = null;

		try {
			vehicleData = new Scanner(new File(dataFileName));
		} catch(FileNotFoundException e) {
			System.out.println("Input file not found.");
            return;
		}

        int day = 1;
        String prevDataPointSensor = null;
        int prevDataPointTime = 0;
        int j = 0;  // tem counter to help print first 10 data points for each day
		while( vehicleData.hasNextLine() ) {
            // read row
            String dataPoint = vehicleData.nextLine();

            if(dataPoint.length() >= 2) {
                // sensor this data point is from
                String dataPointSensor = dataPoint.substring(0,1);

                // timestamp this data point was generated
                int dataPointTime = 0;
                try {
                    dataPointTime = Integer.parseInt(dataPoint.substring(1));
                } catch(NumberFormatException e) {
                    // invalid timestamp - move to next data point
                    continue;
                }

                if(dataPointTime < prevDataPointTime) {
                    // switched into a new day of data gatheing
                    day++;
                    j = 0;
                    System.out.println("\nDAY " + day);
                } else if(j == 0) {
                    System.out.println("\nDAY " + day);
                }

                if(j < 10) {
                    System.out.printf("\tSensor: %s | Time: %d\n", dataPointSensor, dataPointTime);
                    j++;
                }               

                // completed processing - update previous with this data point data
                prevDataPointSensor = dataPointSensor;
                prevDataPointTime = dataPointTime;

            } else {
                // invalid data point - data points must be of ^[A-Za-z][0-9]+$ pattern

            }            
        }
    }
}
