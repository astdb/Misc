func sockMerchant(n int32, ar []int32) int32 {
    // sort ar
    sort.Slice(ar, func(i, j int) bool { return ar[i] < ar[j] })
 
    // count pairs
    i := 0
    var pairs int32
    pairs = 0
    for i < len(ar) {
            if i+1 < len(ar) {
                if ar[i] == ar[i+1] {
                    pairs++
                    i += 2
                } else {
                    i++
                }
            } else {
                i++
            }
    }

    return pairs
}
