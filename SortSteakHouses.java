// IMPORT LIBRARY PACKAGES NEEDED BY YOUR PROGRAM
// SOME CLASSES WITHIN A PACKAGE MAY BE RESTRICTED
// DEFINE ANY CLASS AND METHOD NEEDED
import java.util.List;
import java.util.Arrays;
import java.util.ArrayList;
// CLASS BEGINS, THIS CLASS IS REQUIRED
public class SortSteakHouses
{
    // METHOD SIGNATURE BEGINS, THIS METHOD IS REQUIRED
    List<List<Integer>> nearestXsteakHouses(int totalSteakhouses, 
                                         List<List<Integer>> allLocations, 
                                         int numSteakhouses)
	{
        // WRITE YOUR CODE HERE
        
        // results collection
        List<List<Integer>> result = new ArrayList<List<Integer>>();
        
        // collection of steakhouse objects with distances calculated
        SteakHouse[] distances = new SteakHouse[allLocations.size()];
        
        // populate list of steakhouse objects with distances
        int i = 0;
        for(List<Integer> d: allLocations) {
            if(i < distances.length) {
                int x = d.get(0);
                int y = d.get(1);
                distances[i] = new SteakHouse(x, y, Math.sqrt((x*x) + (y*y)));
                
                i++;
            }
        }
        
        // sort steakhouses by distance to customer
        Arrays.sort(distances);
        
        // extract the first numSteakhouses locations from distances collection and populate 
        // results collection
        for(int j = 0; j < numSteakhouses; j++) {
            if(j < distances.length) {
                List<Integer> thisSH = new ArrayList<Integer>();
                thisSH.add(new Integer(distances[j].getX()));
                thisSH.add(new Integer(distances[j].getY()));
                
                result.add(thisSH);
            }
        }
        
        return result;
        
    }
    // METHOD SIGNATURE ENDS
}

// helper class representing a steakhouse with distance and coordinates, implementing 
// Comparable interface so a collection of SteakHouse objects can be sorted using Arrays.sort()
class SteakHouse implements Comparable<SteakHouse> {
    private int X;
    private int Y;
    private double distance;
    
    public SteakHouse(int x, int y, double dist) {
        this.X = x;
        this.Y = y;
        this.distance = dist;
    }
    
    protected double getDistance() {
        return this.distance;   
    }
    
    protected int getX() {
        return this.X;
    }
    
    protected int getY() {
        return this.Y;
    }
    
    @Override
    public int compareTo(SteakHouse o) {
        if(this.distance < o.distance) {
            return -1;
        }
        
        if(this.distance > o.distance) {
            return 1;
        }
        
        return 0;
    }
}
