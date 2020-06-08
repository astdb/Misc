'use strict';

let tests = [[], [0], [1,2], [1,2,3,4], [2,2,2,2]];

for(let i = 0; i < tests.length; i++) {
  console.log("subArr(" + tests[i] + ") == " + printSubArr(subArr(tests[i])));
}

// subArr() generates all sub arrays of arr
// start with an outer loop of subarray lengths (starting 1). In the inner loop, set a sub array start index and increment
// by 1 until end of subarray is at the end of arr. 
function subArr(arr) {
  let result = [];

  for(let subArrLen = 1; subArrLen <= arr.length; subArrLen++) {
    for(let subArrStart = 0; subArrStart+subArrLen <= arr.length; subArrStart++) {
      result.push(arr.slice(subArrStart, subArrStart+subArrLen));
    }
  }

  return result;
}

// pretty printer for nested arrays
function printSubArr(arr) {
  let result = "[";
  for(let i = 0; i < arr.length; i++) {
    if(i > 0) {
      result += ", [";
    } else {
      result += "[";
    }

    for(let j = 0; j < arr[i].length; j++) {
      if(j > 0) {
        result += ", " + arr[i][j];
      } else {
        result += arr[i][j];
      }
    }
    
    result += "]";
  }

  result += "]";

  return result;
}
