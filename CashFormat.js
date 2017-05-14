<script> 
 var expamount = 45.56;
 var cents_num = expamount - Math.floor(expamount);
 if(cents_num == 0) {
    var cents = '00';
  } else {
    var cents_arr = (expamount + '').split('.');
    var cents = cents_arr[1];
  }

  var amt_str = '';
  var amt = Math.floor(expamount);

  var dgcount = 0;  // digit group count
  var quot = Math.floor(amt/10);
  var rem = amt%10;
  console.log('rem: ' + rem)
  amt = quot;
  amt_str = rem + amt_str;
  dgcount++;
  console.log(amt_str + '(' + dgcount + ')');

  while(quot > 0) {
    quot = Math.floor(amt/10);
    rem = amt%10;
    console.log('rem: ' + rem)
    amt = quot;

    if(dgcount >= 3) {
      amt_str = rem + ',' + amt_str;
      dgcount = 0;
    } else {
      amt_str = rem + amt_str;
      // dgcount++;
    }

    dgcount++;
    console.log(amt_str + '(' + dgcount + ')');
  }

  // amt_str = '$' + amt_str + '.' + cents;
  amt_str = '$' + amt_str;
  console.log(amt_str);
</script>
