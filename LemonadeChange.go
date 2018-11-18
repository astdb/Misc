}

func lemonadeChange(bills []int) bool {
        fivers := 0
        tens := 0
        twenties := 0

        for _, tx := range bills {
                if tx == 5 {
                        fivers++
                }

                if tx == 10 {
                        if fivers > 0 {
                                fivers--
                                tens++
                        } else {
                                return false
                        }
                }

                if tx == 20 {
                        if tens > 0 && fivers > 0 {
                                tens--
                                fivers--
                                twenties++
                        } else if fivers >= 3 {
                                fivers = fivers - 3
                                twenties++
                        } else {
                                return false
                        }
                }
        }

        return true
}
