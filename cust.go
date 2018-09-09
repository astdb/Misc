package main

import (
	"encoding/xml"
	"errors"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"os"
	"strconv"
	"strings"
)

func main() {
	http.HandleFunc("/", mainHandler)
	http.HandleFunc("/customer/", customerHandler)

	fmt.Println("Serving website at http://localhost:8080/..")
	log.Fatal(http.ListenAndServe(":8080", nil))
}

// display main page / customer list
func mainHandler(w http.ResponseWriter, r *http.Request) {
	// capture page ID from URL
	fmt.Println("Starting mainHandler()")
	urlVals := r.URL.Query()
	pageIDStr, exists := urlVals["page"]
	fmt.Println(pageIDStr)

	pageID := 1
	if !exists || len(pageIDStr) != 1 {
		pageID = 1
	} else {
		var err error = nil
		pageID, err = strconv.Atoi(pageIDStr[0])
		fmt.Println("PageID: ", pageID)
		if err != nil {
			fmt.Println("Redirecting..")
			http.Redirect(w, r, "/", 301) // or pageID = 1
			return
		}
	}

	// fmt.Println("PageID: ", pageID)
	customerList, err := GetCustomers(pageID)
	if err != nil {
		fmt.Fprintf(w, "<!DOCTYPE html><html><head><link rel='icon' type='image/png' href='data:image/png;base64,iVBORw0KGgo='></head><body>Error retrieving data | <a href='/'>home</a><body></html>")
		return
	}

	custMap := map[int]Customer{}
	for _, cust := range customerList {
		_, exists := custMap[cust.CustID]
		if !exists {
			custMap[cust.CustID] = cust
		}
	}

	ordersList, _ := GetOrders(1, 0)
	for _, ord := range ordersList {
		cust, exists := custMap[ord.CustomerID]
		if exists {
			cust.CustTotal++
			custMap[cust.CustID] = cust
		}
	}

	// check if before/after API pages exist and provide back/next buttons
	beforePage := APIPageValid(pageID - 1)
	nextPage := APIPageValid(pageID + 1)
	back := ""
	next := ""

	if beforePage {
		back = fmt.Sprintf("<a href='/page=%d'>back</a>", pageID-1)
	}

	if nextPage {
		next = fmt.Sprintf("<a href='/page=%d'>next</a>", pageID+1)
	}

	var b strings.Builder
	// Note: <link rel='icon' type='image/png' href='data:image/png;base64,iVBORw0KGgo='> part in the <head> prevents the HTTP handler function possibly getting called twice if the prowser decides to look for a favicon
	b.WriteString(fmt.Sprintf("<!DOCTYPE html><html><head><link rel='icon' type='image/png' href='data:image/png;base64,iVBORw0KGgo='></head><body><table><thead><tr><th>CUSTOMER</th><th>ORDERS PLACED</th></tr></thead><tbody><tr><td>%s</td><td>%s</td></tr>", back, next))
	for _, c := range custMap {
		b.WriteString(fmt.Sprintf("<tr><td><a href='/customer/?id=%d'>%s %s</a></td><td>%d</td></tr>", c.CustID, c.CustFN, c.CustLN, c.CustTotal))
	}

	b.WriteString("</tbody></table></body></html>")
	fmt.Fprintf(w, b.String())
	return
}

// display customer page
func customerHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Starting customerHandler()")
	w.Header().Set("Content-Type", "text/html; charset=utf-8") // send hypertext to the browser (responsewriter headers are plaintext by default)

	// capture Customer ID from URL
	urlVals := r.URL.Query()
	custIDStr, exists := urlVals["id"]
	// fmt.Println(custIDStr)
	if !exists || len(custIDStr) != 1 {
		fmt.Fprintf(w, "Invalid request | <a href='/'>back</a>")
		return
	}

	custID, err := strconv.Atoi(custIDStr[0])
	if err != nil {
		fmt.Fprintf(w, "Invalid customer ID | <a href='/'>back</a>")
		return
	}

	thisCust, err := GetCustomer(custID)
	if err != nil {
		fmt.Fprintf(w, "Invalid customer ID | <a href='/'>back</a>")
		return
	}

	var custLifeTimeVal float64 = 0
	orders, err := GetOrders(1, custID)

	noOrders := false
	if err != nil {
		noOrders = true
	} else {
		for _, order := range orders {
			custLifeTimeVal += order.TotalExTax
		}
	}	

	var b strings.Builder
	b.WriteString(fmt.Sprintf("<h2>%s %s</h2><b>Lifetime Value: $%.2f</b> | <a href='/'>back</a><br /><br /><table border='1'><thead><tr><th>ORDER ID</th><th>PLACED</th><th>STATUS</th><th>TOTAL</th></tr></thead><tbody>", thisCust.CustFN, thisCust.CustLN, custLifeTimeVal))
	
	if noOrders {
		b.WriteString(fmt.Sprintf("<tr><td colspan='4'><i>No orders found for this customer.</i></td></tr>"))
	} else {
		for _, c := range orders {
			b.WriteString(fmt.Sprintf("<tr><td>%d</td><td>%s</td><td>%s</td><td>$%.2f</td></tr>", c.OrderID, c.DateCreated, c.OrderStatus, c.TotalExTax))
		}
	}

	b.WriteString("</tbody></table>")
	fmt.Fprintf(w, b.String()) //	send HTML to browser

	return
}

// check if a given customer API page is valid
func APIPageValid(page int) bool {
	if page <= 0 {
		return false
	}

	pageSize := 50
	client := &http.Client{} // HTTP client to make API reqs
	APILocation := fmt.Sprintf("https://store-velgoi8q0k.mybigcommerce.com/api/v2/customers?limit=%d&page=%d", pageSize, page)
	req, _ := http.NewRequest("GET", APILocation, nil)
	req.Header.Add("Authorization", `Basic dGVzdDoyNTI1ZGY1NjQ3N2Y1OGU1ODY4YzI0MGVlNTIyOGIwYjVkNDM2N2M0`)
	req.Header.Add("Content-Type", `application/json`)
	apiDataResp, err := client.Do(req)
	defer apiDataResp.Body.Close()

	if err == nil && apiDataResp.StatusCode == http.StatusOK {
		// valid API page
		return true
	} else {
		return false
	}
}

func GetCustomer(customerID int) (Customer, error) {
	client := &http.Client{} // HTTP client to make API reqs
	customerAPILocation := fmt.Sprintf("https://store-velgoi8q0k.mybigcommerce.com/api/v2/customers/%d", customerID)
	req, _ := http.NewRequest("GET", customerAPILocation, nil)
	req.Header.Add("Authorization", `Basic dGVzdDoyNTI1ZGY1NjQ3N2Y1OGU1ODY4YzI0MGVlNTIyOGIwYjVkNDM2N2M0`)
	req.Header.Add("Content-Type", `application/json`)
	apiDataResp, err := client.Do(req)
	defer apiDataResp.Body.Close()

	var cust Customer

	if err == nil && apiDataResp.StatusCode == http.StatusOK {
		xmlData, err := ioutil.ReadAll(apiDataResp.Body)
		if err != nil {
			fmt.Fprintf(os.Stderr, "Error reading API response: %v\n", err)
			os.Exit(1)
		}

		// fmt.Println(string(xmlData))
		unmarshalError := xml.Unmarshal(xmlData, &cust)
		if unmarshalError != nil {
			fmt.Println("Error unmarshalling")
			fmt.Fprintf(os.Stderr, "%v", unmarshalError)
			// os.Exit(1)
			return cust, unmarshalError
		}

		return cust, nil
	} else {
		if err != nil {
			// request error
			fmt.Fprintf(os.Stderr, "Error fetching data from API URL (%s): %v\n", customerAPILocation, err)
		} else {
			// HTTP non-200
			err = errors.New(fmt.Sprintf("Non OK HTTP status %v retriving %s\n", apiDataResp.StatusCode, customerAPILocation))
			fmt.Println(err)
		}
		return cust, err
	}
}

func GetCustomers(page int) ([]Customer, error) {
	if page <= 0 {
		page = 1
	}

	pageSize := 50

	client := &http.Client{} // HTTP client to make API reqs
	customerAPILocation := fmt.Sprintf("https://store-velgoi8q0k.mybigcommerce.com/api/v2/customers?limit=%d&page=%d", pageSize, page)
	req, _ := http.NewRequest("GET", customerAPILocation, nil)
	req.Header.Add("Authorization", `Basic dGVzdDoyNTI1ZGY1NjQ3N2Y1OGU1ODY4YzI0MGVlNTIyOGIwYjVkNDM2N2M0`)
	req.Header.Add("Content-Type", `application/json`)
	apiDataResp, err := client.Do(req)
	defer apiDataResp.Body.Close()

	var customerList Customers

	if err == nil && apiDataResp.StatusCode == http.StatusOK {
		xmlData, err := ioutil.ReadAll(apiDataResp.Body)
		if err != nil {
			fmt.Fprintf(os.Stderr, "Error reading API response: %v\n", err)
			os.Exit(1)
		}

		// fmt.Println(string(xmlData))
		unmarshalError := xml.Unmarshal(xmlData, &customerList)
		if unmarshalError != nil {
			fmt.Println("Error unmarshalling")
			fmt.Fprintf(os.Stderr, "%v", unmarshalError)
			os.Exit(1)
		}
	} else {
		if err != nil {
			// request error
			fmt.Fprintf(os.Stderr, "Error fetching data from API URL (%s): %v\n", customerAPILocation, err)
		} else {
			// HTTP non-200
			err = errors.New(fmt.Sprintf("Non OK HTTP status %v retriving %s\n", apiDataResp.StatusCode, customerAPILocation))
			fmt.Println(err)
		}
		return nil, err
	}

	return customerList.CustomersList, nil
}

func GetOrders(page int, customerID int) ([]Order, error) {
	if page <= 0 {
		page = 1
	}

	client := &http.Client{} // HTTP client to make API reqs`
	// page := 1

	var productAPILocation string
	if customerID > 0 {
		// get orders for a sepcific customer
		productAPILocation = fmt.Sprintf("https://store-velgoi8q0k.mybigcommerce.com/api/v2/orders?customer_id=%d&limit=250&page=%d", customerID, page)
	} else {
		productAPILocation = fmt.Sprintf("https://store-velgoi8q0k.mybigcommerce.com/api/v2/orders?limit=250&page=%d", page)
	}

	req, _ := http.NewRequest("GET", productAPILocation, nil)
	req.Header.Add("Authorization", `Basic dGVzdDoyNTI1ZGY1NjQ3N2Y1OGU1ODY4YzI0MGVlNTIyOGIwYjVkNDM2N2M0`)
	req.Header.Add("Content-Type", `application/json`)
	apiDataResp, err := client.Do(req)
	defer apiDataResp.Body.Close()
	var orderList Orders

	if err == nil && apiDataResp.StatusCode == http.StatusOK {
		// dec := xml.NewDecoder(apiDataResp.Body)

		xmlData, err := ioutil.ReadAll(apiDataResp.Body)
		if err != nil {
			fmt.Fprintf(os.Stderr, "Error reading API response: %v\n", err)
			return orderList.OrdersList, errors.New(fmt.Sprintf("Error reading API response: %v\n", err))
			// os.Exit(1)
		}
		// fmt.Println(string(xmlData))
		unmarshalError := xml.Unmarshal(xmlData, &orderList)
		if unmarshalError != nil {
			fmt.Println("Error unmarshalling orders XML")
			fmt.Fprintf(os.Stderr, "%v", unmarshalError)
			return orderList.OrdersList, errors.New(fmt.Sprintf("Error unmarshalling order XML: %v\n", unmarshalError))
			// os.Exit(1)

		}
	} else {
		fmt.Fprintf(os.Stderr, "Error fetching data from API endpoint (%s): %v\n", productAPILocation, err)
		// apiDataResp.Body.Close()
		// os.Exit(1)
		return orderList.OrdersList, errors.New(fmt.Sprintf("Error fetching data from API endpoint (%s): %v\n", productAPILocation, err))
	}

	// for _, v := range orderList.OrdersList {
	// 	fmt.Println(v.OrderID, v.CustomerID, v.DateCreated, v.TotalExTax)
	// }

	return orderList.OrdersList, nil
}

type Customers struct {
	CustomersList []Customer `xml:"customer"`
}

type Customer struct {
	CustID    int    `xml:"id"`
	CustFN    string `xml:"first_name"`
	CustLN    string `xml:"last_name"`
	CustEmail string `xml:"email"`
	CustTotal int
}

type Orders struct {
	OrdersList []Order `xml:"order"`
}

type Order struct {
	OrderID     int     `xml:"id"`
	CustomerID  int     `xml:"customer_id"`
	DateCreated string  `xml:"date_created"`
	TotalExTax  float64 `xml:"total_ex_tax"`
	OrderStatus string  `xml:"status"`
}
