function countingValleys(n, s) {
    let alt = 0;    // altitude in units (can be negative) - initialized to sea level
    let inValley = false; // flag denoting if hiker currently in valley - valley is a sequence of consecutive steps below sea level, starting with a step down from sea level and ending with a step up to sea level
    let valleyCount = 0;    // no. of valleys traversed
 
    // for each step taken
    for(let i = 0; i < s.length; i++) {
        let curStep = s.charAt(i);
        if(curStep == 'D') {
            // downhill step
            // if below sea level and not in a current valley, start valley
            if((alt == 0) && !inValley) {
                inValley = true;
            }
            
            alt--;    // decrease altitude
 
        } else if(curStep == 'U') {
            // uphill step
 
            // if currently in valley and one step below sea level, reset valley flag and count valley
            if(inValley && (alt == -1)) {
                inValley = false;
                valleyCount++;
            }
            
            alt++;    // increase altitude
        }
    }
 
    return valleyCount;


}
