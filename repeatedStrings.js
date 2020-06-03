function repeatedString(s, n) {
    // number of full times length of s repeated in n
    let fullTimes = Math.floor(n/s.length);
 
    // number of 'a's in s
    let aCount = 0;
    for(let i = 0; i < s.length; i++) {
        if(s.charAt(i) === 'a') {
            aCount++;
        }
    }
 
    // number of characters to fill in any remainder to fill upto n characters and how many 'a's would be needed
    let fillIn = n % s.length;
    let fillInAs = 0;
    for(let j = 0; j < fillIn; j++) {
        if(s.charAt(j) === 'a') {
            fillInAs++;
        }
    }
 
    return ((aCount * fullTimes) + fillInAs);

}
