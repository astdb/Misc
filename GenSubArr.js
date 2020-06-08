'use strict';

let tests = [[], [0], [1,2], [1,2,3,4], [2,2,2,2]];   // basic tests

// generate large subarray test case
// let largeTestCase = [];
// for(let i = 1; i <= 1000; i++) {
//   largeTestCase.push(i)
// }
// tests.push(largeTestCase);

for(let i = 0; i < tests.length; i++) {
  // console.log("Test #" + (i+1) + ": subArr(" + tests[i] + ") == " + printSubArr(subArr(tests[i])));
  // console.log("Test #" + (i+1) + ": subArr(" + tests[i] + ") == " + subArr(tests[i]));

  let res = [];
  subArr2(tests[i], 0, 0, res);
  console.log("Test #" + (i+1) + ": subArr(" + tests[i] + ") == " + printSubArr(res));
}


// recursive version of subarray generation function
function subArr2(arr, start, end, res) {
  if(end >= arr.length) {
    return

  } else if(start > end) {
    subArr2(arr, 0, (end+1), res);

  } else {
    let thisSubArr = [];

    for(let i = start; i < end; i++) {
      thisSubArr.push(arr[i]);
    }

    thisSubArr.push(arr[end]);
    res.push(thisSubArr);

    subArr2(arr, (start+1), end, res)
  }
}

// subArr() generates all sub-arrays of arr
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
