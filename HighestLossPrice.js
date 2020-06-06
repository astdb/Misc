
/*
Problem: Given a list of daily prices, calculate the largest loss that can be incurred on buying an item on any day and selling  on any day afterwards.

Solution: An initial brute-force version of the solution would be for to traverse the list of prices, and for each day traverse the rest of the list to see the lowest future price. This would provide the correct result but would run at exponential runtime complexity. 

A better version of the solution depends on the insight that for each day, the program only needs to know the highest past price observed. This can then be used to calculate the potential loss for each day, and the highest loss would be the sought after result. This can be calculated in a single pass of the price list with linear runtime complexity and constant space complexity. Below implementation uses this method. 
*/

tests = [[7,4,2,9], [2,4,6,8], [7,4,8,2,9]];

for(let i = 0; i < tests.length; i++) {
    let test = tests[i];
    console.log("Test #"+ (i+1) +": worstLosingStreak(" + test + ") == " + worstLosingStreak(test));
}

// worstLosingStreak() computes the highest loss that can be incurred by buying on any day in the given range and selling on a subsequent day. 
function worstLosingStreak(nums) {
    let pastHighestPrice = 0;   // highest price encountered until now
    let highestLoss = 0;        // highest loss encountered until now

    // for each day or pricing
    for(let i = 0; i < nums.length; i++) {
        if(i == 0) {
            // initialize pastHighestPrice to first day's price
            pastHighestPrice = nums[i];
        } else {
            // if the current day's price incurs the highest loss observed so far (compared to the highest price observed so far), update the highest loss
            if((pastHighestPrice-nums[i]) > highestLoss) {
                highestLoss = pastHighestPrice-nums[i];
            }

            // if today's price is higher than the highest price observed so far, update the highest price
            if(nums[i] > pastHighestPrice) {
                pastHighestPrice = nums[i];
            }
        }
    }

    return highestLoss;
}
