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
	// declare handlers for main paths of the web app (/ for main and /customers/ for viewing specific customer detail)
	http.HandleFunc("/", mainHandler)
	http.HandleFunc("/customer/", customerHandler)

	// start server
	fmt.Println("Serving website at http://localhost:8080/..")
	log.Fatal(http.ListenAndServe(":8080", nil))
}

// ---------------------------------- HTTP handler functions ----------------------------------------
// display main page / customer list
func mainHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "text/html; charset=utf-8") // set headers to send hypertext to the browser (responsewriter headers are plaintext by default)

	//-----------------------------
	// capture page ID from URL- set to 1 if not specified
	// Customers API will be read in 50-item pages - pageID (e.g. http://localhost:8080/?page=3) specifies what page to read

	urlVals := r.URL.Query()
	custIDStr, exists := urlVals["page"]
	var pageID int = 1
	if !exists || len(custIDStr) != 1 {
		pageID = 1
	}

	var err error = nil
	if len(custIDStr) > 0 {
		pageID, err = strconv.Atoi(custIDStr[0])
	}

	if err != nil {
		pageID = 1
	}
	//-----------------------------

	// get customers list for specified page
	customerList, err := GetCustomers(pageID)
	noCust := false
	if err != nil {
		noCust = true // error retrieving customer list / no customers
	}

	// load customers into a hashmap indexed by customer ID - this helps to locate customer objects faster to update order totals
	custMap := map[int]Customer{}
	for _, cust := range customerList {
		_, exists := custMap[cust.CustID]
		if !exists {
			custMap[cust.CustID] = cust
		}
	}

	// get list of orders : GetOrders(APIPageNo, CustomerID)
	ordersList, err := GetOrders(1, 0)

	// check if orders are returned
	noOrds := false
	if err != nil {
		noOrds = true
	} else {
		for _, ord := range ordersList {
			cust, exists := custMap[ord.CustomerID]
			if exists {
				cust.CustTotal++
				custMap[cust.CustID] = cust
			}
		}
	}

	// check if before/after API pages exist and provide back/next navigation buttons on webpage
	beforePage := APIPageValid(pageID - 1)
	nextPage := APIPageValid(pageID + 1)
	back := ""
	next := ""

	if beforePage {
		back = fmt.Sprintf("<a href='/?page=%d'>[<< previous page]</a>", pageID-1)
	}

	if nextPage {
		next = fmt.Sprintf("<a href='/?page=%d'>[next page >>]</a>", pageID+1)
	}

	// build customer browser webpage
	var b strings.Builder
	// Note: <link rel='icon' type='image/png' href='data:image/png;base64,iVBORw0KGgo='> part in the <head> prevents the HTTP handler function possibly getting called twice if the browser decides to look for a favicon
	b.WriteString(fmt.Sprintf("<!DOCTYPE html><html><head><link rel='icon' type='image/png' href='data:image/png;base64,iVBORw0KGgo='></head><body><table><thead><tr><th colspan='4'><a href='/'>[home]</a></th></tr><tr><th>CUSTOMER</th><th>ORDERS PLACED</th><th>%s</th><th>%s</th></th></thead><tbody>", back, next))

	if noCust || noOrds {
		b.WriteString(fmt.Sprintf("<tr><td colspan='2'><i>No customers found.</i></td><td colspan='2'></td></tr>"))
	} else {
		for _, c := range custMap {
			b.WriteString(fmt.Sprintf("<tr><td><a href='/customer/?id=%d'>%s %s</a></td><td>%d</td><td colspan='2'></td></tr>", c.CustID, c.CustFN, c.CustLN, c.CustTotal))
		}
	}

	b.WriteString("</tbody></table></body></html>")
	fmt.Fprintf(w, b.String()) // send data to browser
	return
}

// display customer specific page (e.g. http://localhost:8080/customer/?id=10)
func customerHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "text/html; charset=utf-8") // set headers to send hypertext to the browser (responsewriter headers are plaintext by default)

	// capture Customer ID from URL (customer/?id=10)
	//-----------------------------
	urlVals := r.URL.Query()
	custIDStr, exists := urlVals["id"]
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
	//-----------------------------

	var custLifeTimeVal float64 = 0     // order total for this customer
	orders, err := GetOrders(1, custID) // get orders list for this customer

	// check if orders are returned
	noOrders := false
	if err != nil {
		noOrders = true
	} else {
		for _, order := range orders {
			custLifeTimeVal += order.TotalExTax
		}
	}

	// build customer webpage content
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
	fmt.Fprintf(w, b.String()) // send data to browser
	return
}

// ---------------------------------- utility functions ----------------------------------------

// check if a given customer API page is valid - used to generate back/forward buttons to navigate customer browser webpage
func APIPageValid(page int) bool {
	if page <= 0 {
		return false
	}

	pageSize := 50 // customer API is iterated in 50-item pages

	client := &http.Client{} // HTTP client to make API reqs

	// prep request
	APILocation := fmt.Sprintf("https://store-velgoi8q0k.mybigcommerce.com/api/v2/customers?limit=%d&page=%d", pageSize, page)
	req, _ := http.NewRequest("GET", APILocation, nil)

	// set authorization / content type headers
	req.Header.Add("Authorization", `Basic dGVzdDoyNTI1ZGY1NjQ3N2Y1OGU1ODY4YzI0MGVlNTIyOGIwYjVkNDM2N2M0`)
	req.Header.Add("Content-Type", `application/json`)

	// make request
	apiDataResp, err := client.Do(req)
	defer apiDataResp.Body.Close()

	if err == nil && apiDataResp.StatusCode == http.StatusOK {
		// valid API page
		return true
	} else {
		return false
	}
}

// Get Customer object for a given ID
func GetCustomer(customerID int) (Customer, error) {
	client := &http.Client{} // HTTP client to make API reqs

	// prep request
	customerAPILocation := fmt.Sprintf("https://store-velgoi8q0k.mybigcommerce.com/api/v2/customers/%d", customerID)
	req, _ := http.NewRequest("GET", customerAPILocation, nil)

	// set authorization / content type headers
	req.Header.Add("Authorization", `Basic dGVzdDoyNTI1ZGY1NjQ3N2Y1OGU1ODY4YzI0MGVlNTIyOGIwYjVkNDM2N2M0`)
	req.Header.Add("Content-Type", `application/json`)

	// make request
	apiDataResp, err := client.Do(req)
	defer apiDataResp.Body.Close()

	var cust Customer // placeholder for customer object

	if err == nil && apiDataResp.StatusCode == http.StatusOK {
		// request successful - get XML data from response body
		xmlData, err := ioutil.ReadAll(apiDataResp.Body)
		if err != nil {
			return cust, errors.New(fmt.Sprintf("Error reading customer API response for ID %d: %v\n", customerID, err))
		}

		// read XML data into Customer object
		unmarshalError := xml.Unmarshal(xmlData, &cust)
		if unmarshalError != nil {
			return cust, unmarshalError
		}

		return cust, nil // return Customer object
	} else {
		// request to get Customer data unsuccessful - return suitable error
		if err != nil {
			// request error
			fmt.Fprintf(os.Stderr, "Error fetching data from API URL (%s): %v\n", customerAPILocation, err)
		} else {
			// HTTP non-200
			err = errors.New(fmt.Sprintf("Non OK HTTP status %v retriving %s\n", apiDataResp.StatusCode, customerAPILocation))
		}
		return cust, err
	}
}

// Get Customers list from a given page of customers API
func GetCustomers(page int) ([]Customer, error) {
	if page <= 0 {
		page = 1
	}

	pageSize := 50

	client := &http.Client{} // HTTP client to make API reqs

	// prep request
	customerAPILocation := fmt.Sprintf("https://store-velgoi8q0k.mybigcommerce.com/api/v2/customers?limit=%d&page=%d", pageSize, page)
	req, _ := http.NewRequest("GET", customerAPILocation, nil)

	// set authorization / content type headers
	req.Header.Add("Authorization", `Basic dGVzdDoyNTI1ZGY1NjQ3N2Y1OGU1ODY4YzI0MGVlNTIyOGIwYjVkNDM2N2M0`)
	req.Header.Add("Content-Type", `application/json`)

	// make request
	apiDataResp, err := client.Do(req)
	defer apiDataResp.Body.Close()

	var customerList Customers // placeholder for Customers object encompassing customer list

	if err == nil && apiDataResp.StatusCode == http.StatusOK {
		// request successful - get XML data from response body
		xmlData, err := ioutil.ReadAll(apiDataResp.Body)
		if err != nil {
			return customerList.CustomersList, errors.New(fmt.Sprintf("Error reading customers API response: %v\n", err))
		}

		// read XML data into Customers object
		unmarshalError := xml.Unmarshal(xmlData, &customerList)
		if unmarshalError != nil {
			return customerList.CustomersList, unmarshalError
		}
	} else {
		// request to get customer data unsuccessful - return suitable error
		if err != nil {
			// request error
			return customerList.CustomersList, errors.New(fmt.Sprintf("Error fetching data from API URL (%s): %v\n", customerAPILocation, err))
		} else {
			// HTTP non-200
			return customerList.CustomersList, errors.New(fmt.Sprintf("Non OK HTTP status %v retriving %s\n", apiDataResp.StatusCode, customerAPILocation))
		}

		return nil, err
	}

	return customerList.CustomersList, nil // return customer list
}

// Get orders list for a given API page - if a >0 customer ID is provided, return orders list filtered by customer
func GetOrders(page int, customerID int) ([]Order, error) {
	if page <= 0 {
		page = 1
	}

	client := &http.Client{} // HTTP client to make API reqs

	// prep request
	var ordersAPILocation string
	if customerID > 0 {
		// get orders for a specific customer
		ordersAPILocation = fmt.Sprintf("https://store-velgoi8q0k.mybigcommerce.com/api/v2/orders?customer_id=%d&limit=250&page=%d", customerID, page)
	} else {
		ordersAPILocation = fmt.Sprintf("https://store-velgoi8q0k.mybigcommerce.com/api/v2/orders?limit=250&page=%d", page)
	}

	req, _ := http.NewRequest("GET", ordersAPILocation, nil)

	// set authorization / content type headers
	req.Header.Add("Authorization", `Basic dGVzdDoyNTI1ZGY1NjQ3N2Y1OGU1ODY4YzI0MGVlNTIyOGIwYjVkNDM2N2M0`)
	req.Header.Add("Content-Type", `application/json`)

	// make request
	apiDataResp, err := client.Do(req)
	defer apiDataResp.Body.Close()

	var orderList Orders       // placeholder for Orders object encompassing orders list
	fullOrderList := []Order{} // full list of orders from across ordes API pages

	if err == nil && apiDataResp.StatusCode == http.StatusOK {
		// request successful - get XML data from response body
		xmlData, err := ioutil.ReadAll(apiDataResp.Body)
		if err != nil {
			fmt.Fprintf(os.Stderr, "Error reading API response: %v\n", err)
			return orderList.OrdersList, errors.New(fmt.Sprintf("Error reading API response: %v\n", err))
		}

		// read XML data into Customers object
		unmarshalError := xml.Unmarshal(xmlData, &orderList)
		if unmarshalError != nil {
			fmt.Println("Error unmarshalling orders XML")
			fmt.Fprintf(os.Stderr, "%v", unmarshalError)
			return orderList.OrdersList, errors.New(fmt.Sprintf("Error unmarshalling order XML: %v\n", unmarshalError))

		}
	} else {
		// request to get customer data unsuccessful - return suitable error
		fmt.Fprintf(os.Stderr, "Error fetching data from API endpoint (%s): %v\n", ordersAPILocation, err)
		return orderList.OrdersList, errors.New(fmt.Sprintf("Error fetching data from API endpoint (%s): %v\n", ordersAPILocation, err))
	}

	// append orders found so far to full orders list, before reading next page of API
	fullOrderList = append(fullOrderList, orderList.OrdersList...)

	// read next pages of API
	// NOTE: if an error occurs in unmarshalling subsequent API pages, fullOrderList will be returned with a nil error as the order data read so far
	page++
	for err == nil && apiDataResp.StatusCode == http.StatusOK {
		// prep request
		var ordersAPILocation string
		if customerID > 0 {
			// get orders for a specific customer
			ordersAPILocation = fmt.Sprintf("https://store-velgoi8q0k.mybigcommerce.com/api/v2/orders?customer_id=%d&limit=250&page=%d", customerID, page)
		} else {
			ordersAPILocation = fmt.Sprintf("https://store-velgoi8q0k.mybigcommerce.com/api/v2/orders?limit=250&page=%d", page)
		}

		req, _ = http.NewRequest("GET", ordersAPILocation, nil)

		// set authorization / content type headers
		req.Header.Add("Authorization", `Basic dGVzdDoyNTI1ZGY1NjQ3N2Y1OGU1ODY4YzI0MGVlNTIyOGIwYjVkNDM2N2M0`)
		req.Header.Add("Content-Type", `application/json`)

		// make request
		apiDataResp, err = client.Do(req)
		defer apiDataResp.Body.Close()

		// request successful - get XML data from response body
		xmlData, err := ioutil.ReadAll(apiDataResp.Body)
		if err != nil {
			fmt.Fprintf(os.Stderr, "Error reading API response: %v\n", err)
			return fullOrderList, nil
		}

		// read XML data into Customers object
		unmarshalError := xml.Unmarshal(xmlData, &orderList)
		if unmarshalError != nil {
			fmt.Println("Error unmarshalling orders XML")
			fmt.Fprintf(os.Stderr, "%v", unmarshalError)
			return fullOrderList, nil

		}

		fullOrderList = append(fullOrderList, orderList.OrdersList...)
		page++
	}

	// return orderList.OrdersList, nil
	return fullOrderList, nil
}

// ---------------------------- define objects for Customers/Orders for reading in XML API data ------------------------
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
