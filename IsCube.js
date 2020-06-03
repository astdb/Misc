
let tests = [27, 0, 9, 45];

for(let i = 0; i < tests.length; i++) {
    test = tests[i];

    console.log("isCube(" + test + ") == " + isCube(test));
}

function isCube(n) {
    console.log("\tisCube(): n == " + n);
    let thisCube = 0;
	
	let i = 0;
	while(n > thisCube) {
        thisCube = Math.pow(i, 3);
        console.log("\t\tisCube(): i == " + i + ", thiCube == " + thisCube);

		if(thisCube === n) {
            console.log("\t\tisCube(): thisCube === n, returning true.");
			return true;
		}

		i++;
	}

    console.log("\tisCube(): off while: n == " + n + ", thisCube == " + thisCube + ", returning false");
	return false;
}
