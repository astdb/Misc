// A set of customer/account pairs indicates which customers owns which accounts.

// Task
// In your choice of programming language, write a function that finds all customers who share all of their accounts.

// Example
// The following example indicates which customers own which accounts. For instance, the customer with id 1 owns the account with id 10 and the account with id 11.

// Cust Account
// 1    10
// 1    11
// 2    13
// 3    11
// 4    14
// 3    10
// 4    13

// Customers 1 and 3 share all of their accounts, 10 and 11, so they are a match.
// Customers 2 and 4 share account 13 but customer 4 also owns account 14, which customer 2 doesn't. They are not a match.


// build a map of customers (int) to accounts they are associated with ([]int).
// to compute the result, determine which customers have similar account lists. 
func SharedAccounts(custAcc []*CustAccMapping) []*AccountCustomer {
    // map to track account-customer relations from input
    accCust := map[int][]int{}
    
    for _, custAccInstance := range custAcc {
        // check if this account is in map
        _, accExists := accCust[custAccInstance.Account]
        if accExists {
            // add to its customer list
            accCust[custAccInstance.Account] = append(accCust, custAccInstance.Customer)
        } else {
            accCust[custAccInstance.Account] = []int{custAccInstance.Customer}
        }
    }
    
    // determine which account numbers have the same list of customers
    result := []AccountCustomer{}
    
    // built a hash-map structure mapping customer lists to account numbers they are associated with
    
    customerListToAccounts := map[hashOfCustList]*AccountCustomer{}
    for accountNumber, customerList := range accCust {
        // check if cust list hash exists
        _, hashExists := customerListToAccounts[hashIntList(customerList)]
        if hashExists {
            // add to current AccCust object's account list
            customerListToAccounts[hashIntList(customerList)].Accounts = append(customerListToAccounts[hashIntList(customerList)].Accounts, accountNumber)
            
        } else {
            // create new
            customerListToAccounts[hashIntList(customerList)] = &AccountCustomer{Accounts: []int{accountNumber}, Customers: customerList}
        }
        
    }
    
    
}

func hashIntList(custList []int) string {
    
    
}

type AccountCustomer struct {
    Accounts []int
    Customers []int
}

type CustAccMapping struct {
    Account int
    Customer int
}
