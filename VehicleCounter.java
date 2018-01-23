
import java.io.*;
import java.util.*;

public class VehicleCounter {
    public static void main(String[] args) {
        // read input data file from command line and process options
        // an invocation must have a minimum of two command arguments provided: [inputfile] [progmode]
        if( args.length < 2 || args[0].trim().length() <= 0 ) {
            System.out.println("Usage: $> java VehicleCounter inputfile programmode [options]");
            return;
        }

        String dataFileName = args[0].trim();

        // program mode
        boolean MODE_TOTALS = false;    // compute total counts, as other options specify
        boolean MODE_PEAKS = false;     // compute peak traffic times
        boolean MODE_SPEEDS = false;    // compute speed distribution
        boolean MODE_DIST = false;      // compute intercar distances

        String TOTAL_BOUND = null;  // direction the totals are requied for  i.e. "north", "south"
        String TOTAL_TIME = null;   // time of day to compute totals for i.e. "morning", "evening", "hourly" etc
        String TOTAL_TERM = null;   // time period to calculate totals per i.e. "daily", "avg"

        String progMode = args[1].trim();
        if (progMode.equalsIgnoreCase("totals")) {
            MODE_TOTALS = true;

            // read options required for TOTALS mode. five command line arguments required in total for invocation in TOTAL mode
            // e.g. $> java VehicleCounter inputdata.file totals south evening average    // output the average of southbound evening totals
            if(args.length < 5) {
                // invalid invocation 
                System.out.println("Invalid total count invocation.");
                return
            }

            TOTAL_BOUND = args[2].trim();
            TOTAL_TIME = args[3].trim();
            TOTAL_TERM = args[4].trim();


        } else if (progMode.equalsIgnoreCase("peaks")) {
            MODE_PEAKS = true;

        } else if (progMode.equalsIgnoreCase("speeds")) {
            MODE_SPEEDS = true;

        } else if (progMode.equalsIgnoreCase("distance")) {
            MODE_DIST = true;
        } else {
            // invalid program mode specifier
            System.out.println("Invalid program mode.");
            return;
        }

        // iterate through counter data and process required metric
        Scanner vehicleData = null;

		try {
			vehicleData = new Scanner(new File(dataFileName));
		} catch(FileNotFoundException e) {
			System.out.println("Input file not found.");
            return;
		}

        int day = 1;
        String prevDataPointSensor = null;      // previously read datapoint
        int prevDataPointTime = 0;
        int j = 0;  // tem counter to help print first j data points for each day
        boolean southBoundDetected = false;
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

                if(j < 100) {
                    System.out.printf("\t%d. Sensor: %s | Time: %d\n", j+1, dataPointSensor, dataPointTime);
                    j++;
                }

                if(prevDataPointSensor != null && dataPointSensor != null) {
                    if(prevDataPointSensor.equalsIgnoreCase("A") && dataPointSensor.equalsIgnoreCase("A")) {
                        if(dataPointTime - prevDataPointTime <= 160) {  // at 60kmh, a car with 2.5m wheelbase would take about 150ms to go over a single sensor
                            // this data point and the previous one are probably made by the same car (going northbound, over sensor A)
                            
                            // TODO - at this point, the vehicle will be counted if the time is in the time period required for analysis
                        }
                    }

                    if(prevDataPointSensor.equalsIgnoreCase("A") && dataPointSensor.equalsIgnoreCase("B")) {
                        if(dataPointTime - prevDataPointTime <= 6) {
                            // southbound car, first set of axles going over A and then B
                            if(southBoundDetected) {
                                // TODO - at this point, the vehicle will be counted if the time is in the time period required for analysis
                                
                            } else {
                                southBoundDetected = true;
                            }
                        }
                    }
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
