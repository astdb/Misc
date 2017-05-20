
"use strict";

var testCases = new Map();

testCases.set(0, "$0.00");
testCases.set(.10, "$0.10");
testCases.set(.01, "$0.01");
testCases.set(1000000000000.00, "$1,000,000,000,000.00");
testCases.set(10.00, "$10.00");
testCases.set(-.10, "-$0.10");
testCases.set("Million bucks", "<invalid>");
testCases.set(3216848.45, "$3,216,848.45");
testCases.set("3216848.45", "$3,216,848.45");

var failures = false;
for(var [testcase, result] of testCases) {
    var thisRes = renderAmt(testcase);
    if(thisRes !== result) {
        failures = true;
        console.log("----FAIL: renderAmt(" + testcase + ") returned " + thisRes + ", expected " + result);
    } else {
        console.log("----PASS: renderAmt(" + testcase + ") == " + thisRes);
    }
}

if(!failures) {
    console.log("All tests have passed 🍺");
}

function renderAmt(expamount) {
    if(!isNumeric(expamount)){
        return "<invalid>";
    }

    var neg = false;
    if(expamount < 0) {
        expamount *= (-1);
        neg = true;
    }

    var cents_num = expamount - Math.floor(expamount);
    if(cents_num == 0) {
        var cents = '00';
    } else {
        var cents_arr = (expamount + '').split('.');
        var cents = cents_arr[1];

        if(cents.length < 2){
            cents = cents + "0";
        }
    }

    var amt_str = '';
    var amt = Math.floor(expamount);

    var dgcount = 0;  // digit group count
    var quot = Math.floor(amt/10);
    var rem = amt%10;
    amt = quot;
    amt_str = rem + amt_str;
    dgcount++;

    while(quot > 0) {
        quot = Math.floor(amt/10);
        rem = amt%10;
        amt = quot;

        if(dgcount >= 3) {
            amt_str = rem + ',' + amt_str;
            dgcount = 0;
        } else {
            amt_str = rem + amt_str;
        }
        
        dgcount++;
    }

    if(neg) {
        amt_str = '-$' + amt_str + '.' + cents;
    } else {
        amt_str = '$' + amt_str + '.' + cents;
    }
    
    return amt_str;
}

// via https://goo.gl/WTbW5g
function isNumeric(n) {
  return !isNaN(parseFloat(n)) && isFinite(n);
}
