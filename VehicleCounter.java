
import java.io.*;
import java.util.*;

public class VehicleCounter {
    public static void main(String[] args) {
        // read and process input data file and operating options from command line
        // a valid invocation must have a minimum of two command arguments provided: [inputfile] [progmode]
        if( args.length < 2 || args[0].trim().length() <= 0 ) {
            System.out.println("Usage: $> java VehicleCounter inputfile programmode [options]");
            return;
        }

        String dataFileName = args[0].trim();

        /*
            The program can be invoked in four major modes: TOTALS, PEAKS, SPEEDS, or DISTANCE. 
            Detail on each mode's operation and invocation syntax as below: 

            TOTALS: 
                This mode provides total vehicle counts in each direction: morning versus evening, per hour, per half hour, 
                per 20 minutes, and per 15 minutes.
                Stats can be displayed for each day of the session, or as averages across all days.

                Syntax:
                ------
                $> java VehicleCounter input.file TOTALS [NORTH|SOUTH] [MORNING | EVENING | HOUR | HALFHOUR | 20MIN | 15MIN] [DAILY | AVG] [DAYX]
                    First arhument specifies the input file name.
                    Second argument specifies program mode.
                    Third argument specifies whether to count south- or north-bound traffic. 
                    Fourth argument specifies the time of day for the stats to be generated.
                    Fifth argument specifies whether the stats should be generated for each day or as averages across all days.
                    Sixtth argument (optional) specifies the day of session (as an integer, 1-indexed) should the fifth argument be "DAILY". 

            PEAKS: 

            SPEEDS: 

            DISTANCE: 
        */

        // program major mode
        boolean MODE_TOTALS = false;    // compute total counts, as other options specify
        boolean MODE_PEAKS = false;     // compute peak traffic times
        boolean MODE_SPEEDS = false;    // compute speed distribution
        boolean MODE_DIST = false;      // compute intercar distances

        String TOTAL_BOUND = null;  // direction the totals are requied for  i.e. "north", "south"
        String TOTAL_TIME = null;   // time of day to compute totals for i.e. "morning", "evening", "hourly" etc
        String TOTAL_TERM = null;   // time period to calculate totals per i.e. "daily", "avg"
        int TOTAL_DAY = 0;          // if TOTAL_TERM specifies DAILY stats, this will indicate which day of he session (e.g. 1, 2 etc)

        String progMode = args[1].trim();
        if (progMode.equalsIgnoreCase("TOTALS")) {
            MODE_TOTALS = true;

            // read options required for TOTALS mode. five command line arguments required in total for invocation in TOTAL mode
            // e.g. $> java VehicleCounter inputdata.file totals south evening average    // output the average of southbound evening totals
            if(args.length < 5) {
                // invalid invocation 
                System.out.println("Invalid total count invocation. \nUsage: $> java VehicleCount input.file TOTALS <DIRECTION> <TIME OF DAY> <DAILY or AVG> <DAY if AVG>");
                return;
            }

            TOTAL_BOUND = args[2].trim();   // NORTH/SOUTH
            TOTAL_TIME = args[3].trim();    // MORNING/EVENING/HOUR/HALFHOUR/15MIN/20MIN
            TOTAL_TERM = args[4].trim();    // DAILY/AVG
            
            try {
                TOTAL_DAY = Integer.parseInt(args[5]);
            } catch(NumberFormatException e) {
                System.out.println("Invalid DAY in TOTALS mode invocation. DAY must be an integer representing a day in the data gathering session.\nUsage: $> java VehicleCount input.file TOTALS <DIRECTION> <TIME OF DAY> <DAILY or AVG> <DAY if AVG>");
            }
            

        } else if (progMode.equalsIgnoreCase("PEAKS")) {
            MODE_PEAKS = true;

        } else if (progMode.equalsIgnoreCase("SPEEDS")) {
            MODE_SPEEDS = true;

        } else if (progMode.equalsIgnoreCase("DIST")) {
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

        int day = 1;    // counter for tracking days of session
        String prevDataPointSensor = null;      // previously read datapoint: sensor
        int prevDataPointTime = 0;              // previously read datapoint: timestamp
        int j = 0;  // tem counter to help print first j data points for each day
        boolean southBoundDetected = false;

        // int carCount = 0;

        // For counting totals for different time periods, we need to record totals for various time periods
        // e.g. morning, evening, hour, half hour etc. 
        // Upon completion of iterating though sensor data file, we will then have a set of totals for each time period
        // e.g 2 values for morning/evening, 24 values for hourly, 48 values for halhourly etc (in AVG mode, these would be totals across
        // days in the session, each of which will have to be divided by the number of days in the session). 
        List<Integer> carCounts = new ArrayList<Integer>();

		while(vehicleData.hasNextLine()) {
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
                    System.out.printf("\t%d. Sensor: %s | Time: %d\n", j+1, dataPointSensor, dataPointTime);
                    j++;
                }

                if(prevDataPointSensor != null && dataPointSensor != null) {
                    // Two "A"s  - possibly NORTH-bound
                    if(prevDataPointSensor.equalsIgnoreCase("A") && dataPointSensor.equalsIgnoreCase("A")) {
                        if((dataPointTime - prevDataPointTime) <= 160) {  // at 60kmh, a car with 2.5m wheelbase would take about 150ms to go over a single sensor
                            // this data point and the previous one are probably made by the same car (going northbound, over sensor A)
                            
                            // TODO - at this point, the vehicle will be counted if the time is in the time period required for analysis 
                            // and count is for Northbound traffic
                            if(TOTAL_TERM.equalsIgnoreCase("DAILY") && day == TOTAL_DAY) {
                                carCount++;
                            }

                            if(TOTAL_TERM.equalsIgnoreCase("AVG")) {
                                carCount++;       
                            }
                        }

                        // TODO: once a car is counted, should prevData be set to NULL? think yes. 
                    }

                    // "A" followed by "B" shortly - possibly SOUTH-bound
                    if(prevDataPointSensor.equalsIgnoreCase("A") && dataPointSensor.equalsIgnoreCase("B")) {
                        if(dataPointTime - prevDataPointTime <= 6) {    // if the wheels reached B sensor within 6ms of reaching A 
                            // southbound car, first set of axles going over A and then B
                            if(southBoundDetected) {
                                // TODO - at this point, the vehicle will be counted if the time is in the time period required for analysis
                                if(TOTAL_TERM.equalsIgnoreCase("DAILY") && day == TOTAL_DAY) {
                                    carCount++;
                                }

                                southBoundDetected = false;
                                
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
