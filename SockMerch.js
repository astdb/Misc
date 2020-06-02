function sockMerchant(n, ar) {
    ar.sort(function(a, b) {
        return a - b;
    });

    let pairs = 0;
    let i = 0;
    while(i < ar.length) {
        if((i + 1) < ar.length) {
            if(ar[i] == ar[i+1]) {
                pairs++;
                i += 2;
            } else {
                i++;
            }
        } else {
            i++;
        }
    }

    return pairs;
}
