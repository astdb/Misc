let tests = [121, -121, 10, 0, 1, -1, 11];

for(let i = 0; i < tests.length; i++) {
	console.log("isPalindromeNumber(" + tests[i] + ") == " + isPalindromeNumber(tests[i]) + "\n");
}

function isPalindromeNumber(x) {
	if(x < 0) {
		return false;
	}

	// turn number into string
  let xStr = x.toString();
  console.log("x == "+ x + ", xStr == " + xStr);

	let endIndex = xStr.length-1;
	let startIndex = 0;

	while(startIndex <= endIndex) {
    console.log("(xStr.charAt(startIndex) !== xStr.charAt(endIndex) => (" + xStr.charAt(startIndex) + " <=> " + xStr.charAt(endIndex));
		if(xStr.charAt(startIndex) !== xStr.charAt(endIndex)) {
			return false;
		}

		endIndex--;
		startIndex++;
	}

	return true;
}

