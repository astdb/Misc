package main

import (
	"encoding/xml"
	"fmt"
	"io"
	"log"
	"net/http"
	"os"
	"strconv"
	"strings"
)

func main() {
	http.HandleFunc("/", mainHandler)
	http.HandleFunc("/customer/", customerHandler)
	log.Fatal(http.ListenAndServe(":8080", nil))
}

// display main page (customers list/order totals)
func mainHandler(w http.ResponseWriter, r *http.Request) {
	client := &http.Client{}       // HTTP client to make API reqs`
	page := 1                      // API pagination counter
	custMap := map[int]*Customer{} // all found customer objects are stored in  hashmap with customer IDs as keys

	CUSTOMER := "customer"
	ID := "id"
	FIRSTNAME := "first_name"
	LASTNAME := "last_name"
	var newCust *Customer

	productAPILocation := fmt.Sprintf("https://store-velgoi8q0k.mybigcommerce.com/api/v2/customers?limit=250&page=%d", page)
	req, _ := http.NewRequest("GET", productAPILocation, nil)
	req.Header.Add("Authorization", `Basic dGVzdDoyNTI1ZGY1NjQ3N2Y1OGU1ODY4YzI0MGVlNTIyOGIwYjVkNDM2N2M0`)
	req.Header.Add("Content-Type", `application/json`)
	apiDataResp, err := client.Do(req)

	for err == nil && apiDataResp.StatusCode == http.StatusOK {
		fmt.Println("Customer API call (main)")
		if err != nil {
			// cannot get data - exit and show results so far
			fmt.Fprintf(os.Stderr, "Error fetching data from API URL (%s): %v\n", productAPILocation, err)
			apiDataResp.Body.Close()
		}

		if apiDataResp.StatusCode != http.StatusOK {
			fmt.Fprintf(os.Stderr, "API request to URL %s failed: %s\n", productAPILocation, apiDataResp.Status)
			apiDataResp.Body.Close()
		}

		// decode read XML data body
		dec := xml.NewDecoder(apiDataResp.Body)
		// fmt.Println("Request successful")

		var stack []string // we'll use a string slice as a stack data structure to pop on/off start/end elements as we read through the XML data body's tokens

		for {
			// get next XML token
			token, err := dec.Token()

			// handle any errors
			if err == io.EOF {
				// reached end of data
				break
			} else if err != nil {
				fmt.Fprintf(os.Stderr, "Error reading XML data body token: %v\n", err)
				os.Exit(1)
			}

			// switch statement to take selective action based on the current token (start, end or data)
			switch token := token.(type) {
			case xml.StartElement:
				stack = append(stack, token.Name.Local)

				if len(stack) > 0 && stack[len(stack)-1] == CUSTOMER {
					// start of a new <work> in XML data: create a new Work instance and pop in to the list of all customers
					newCust = createCustomer()
					// customers = append(customers, newCust)
				}

			case xml.EndElement:
				// XML end element: pop off stack, and finalize current in-memory work object

				// check if there are already XML opening tags stored in stack - if not, we've encountered a closing tag without an opening tag
				if len(stack) <= 0 {
					fmt.Fprintf(os.Stderr, "Attempting to pop an element(%s) without any on stack - possibly malformed XML\n", token.Name.Local)
					os.Exit(1)
				}

				elementPopped := stack[len(stack)-1]
				stack = stack[:len(stack)-1]

				// check for XML consistency - if every end element should have had a corresponding start element
				if elementPopped != token.Name.Local {
					fmt.Fprintf(os.Stderr, "Closing element %s without matching opener (%s) - possibly malformed XML\n", elementPopped, token.Name.Local)
					os.Exit(1)
				}

				if elementPopped == CUSTOMER {
					custMap[newCust.ID] = newCust
					newCust = nil
				}

			case xml.CharData:
				// XML data - populate the current work object based on XML data token (e.g. ID, model, make etc.)
				if len(stack) > 0 && stack[len(stack)-1] == ID {
					IDData, err := strconv.Atoi(strings.TrimSpace(string(token)))

					if err != nil {
						fmt.Fprintf(os.Stderr, "Error converting Customer ID: %v\n", err)
						os.Exit(1)
					}

					if newCust != nil {
						newCust.ID = IDData

					} else {
						fmt.Fprintf(os.Stderr, "ID data(%d) detected without an active current Work struct instance. Possibly malformed XML.", IDData)
						os.Exit(1)
					}
				}

				if len(stack) > 0 && stack[len(stack)-1] == FIRSTNAME {
					FName := strings.TrimSpace(string(token))

					if newCust != nil {
						newCust.FirstName = FName
					} else {
						fmt.Fprintf(os.Stderr, "FirstName(%s) detected without an active current Work struct instance. Possibly malformed XML.", FName)
						os.Exit(1)
					}
				}

				if len(stack) > 0 && stack[len(stack)-1] == LASTNAME {
					LName := strings.TrimSpace(string(token))

					if newCust != nil {
						newCust.LastName = LName
					} else {
						fmt.Fprintf(os.Stderr, "LastName(%s) detected without an active current Work struct instance. Possibly malformed XML.", LName)
						os.Exit(1)
					}
				}
			}
		}

		page++
		productAPILocation = fmt.Sprintf("https://store-velgoi8q0k.mybigcommerce.com/api/v2/customers?limit=250&page=%d", page)
		req, _ = http.NewRequest("GET", productAPILocation, nil)
		req.Header.Add("Authorization", `Basic dGVzdDoyNTI1ZGY1NjQ3N2Y1OGU1ODY4YzI0MGVlNTIyOGIwYjVkNDM2N2M0`)
		req.Header.Add("Content-Type", `application/json`)
		apiDataResp, err = client.Do(req)
	}

	// -------------------------- get orders ------------------------------------------
	page = 1

	ORDER := "order"
	ID = "id"
	CUST_ID := "customer_id"
	VALUE := "total_ex_tax"
	var newOrd *Order

	productAPILocation = fmt.Sprintf("https://store-velgoi8q0k.mybigcommerce.com/api/v2/orders?limit=250&page=%d", page)
	req, _ = http.NewRequest("GET", productAPILocation, nil)
	req.Header.Add("Authorization", `Basic dGVzdDoyNTI1ZGY1NjQ3N2Y1OGU1ODY4YzI0MGVlNTIyOGIwYjVkNDM2N2M0`)
	req.Header.Add("Content-Type", `application/json`)
	apiDataResp, err = client.Do(req)

	for err == nil && apiDataResp.StatusCode == http.StatusOK {
		fmt.Println("Order API call (main)")
		if err != nil {
			// cannot get data - exit and show results so far
			fmt.Fprintf(os.Stderr, "Error fetching data from API URL (%s): %v\n", productAPILocation, err)
			apiDataResp.Body.Close()
		}

		if apiDataResp.StatusCode != http.StatusOK {
			fmt.Fprintf(os.Stderr, "API request to URL %s failed: %s\n", productAPILocation, apiDataResp.Status)
			apiDataResp.Body.Close()
		}

		// invalid orders page doesn't seem to give a non-OK HTTP response (unlike customer API), therefore ending iteration after first page
		// if page == 2 {
		// 	break
		// }

		// decode read XML data body
		dec := xml.NewDecoder(apiDataResp.Body)
		// fmt.Println("Request successful")

		var stack []string // we'll use a string slice as a stack data structure to pop on/off start/end elements as we read through the XML data body's tokens

		for {
			// get next XML token
			token, err := dec.Token()

			// handle any errors
			if err == io.EOF {
				// reached end of data
				break
			} else if err != nil {
				fmt.Fprintf(os.Stderr, "Error reading XML data body token: %v\n", err)
				os.Exit(1)
			}

			// switch statement to take selective action based on the current token (start, end or data)
			switch token := token.(type) {
			case xml.StartElement:
				stack = append(stack, token.Name.Local)

				if len(stack) > 0 && stack[len(stack)-1] == ORDER {
					// start of a new <work> in XML data: create a new Work instance and pop in to the list of all customers
					newOrd = createOrder()
				}

			case xml.EndElement:
				// XML end element: pop off stack, and finalize current in-memory work object

				// check if there are already XML opening tags stored in stack - if not, we've encountered a closing tag without an opening tag
				if len(stack) <= 0 {
					fmt.Fprintf(os.Stderr, "Attempting to pop an element(%s) without any on stack - possibly malformed XML\n", token.Name.Local)
					os.Exit(1)
				}

				elementPopped := stack[len(stack)-1]
				stack = stack[:len(stack)-1]

				// check for XML consistency - if every end element should have had a corresponding start element
				if elementPopped != token.Name.Local {
					fmt.Fprintf(os.Stderr, "Closing element %s without matching opener (%s) - possibly malformed XML\n", elementPopped, token.Name.Local)
					os.Exit(1)
				}

				if elementPopped == ORDER {
					_, ok := custMap[newOrd.CustID]
					if ok {
						custMap[newOrd.CustID].Total++
					}

					newOrd = nil
				}

			case xml.CharData:
				// XML data - populate the current work object based on XML data token (e.g. ID, model, make etc.)
				if len(stack) > 0 && stack[len(stack)-1] == ID {
					IDData, err := strconv.Atoi(strings.TrimSpace(string(token)))

					if err != nil {
						fmt.Fprintf(os.Stderr, "Error converting Order ID: %v\n", err)
						os.Exit(1)
					}

					if newOrd != nil {
						newOrd.ID = IDData

					} else {
						fmt.Fprintf(os.Stderr, "ID data(%d) detected without an active current order struct instance. Possibly malformed XML (page = %d).", IDData, page)
						os.Exit(1)
						continue
					}
				}

				if len(stack) > 0 && stack[len(stack)-1] == CUST_ID {
					IDData, err := strconv.Atoi(strings.TrimSpace(string(token)))

					if err != nil {
						fmt.Fprintf(os.Stderr, "Error converting Order Customer ID: %v\n", err)
						os.Exit(1)
					}

					if newOrd != nil {
						newOrd.CustID = IDData

					} else {
						fmt.Fprintf(os.Stderr, "Customer ID data(%d) detected without an active current order struct instance. Possibly malformed XML.", IDData)
						os.Exit(1)
					}
				}

				if len(stack) > 0 && stack[len(stack)-1] == VALUE {
					IDData, err := strconv.ParseFloat(strings.TrimSpace(string(token)), 64)

					if err != nil {
						fmt.Fprintf(os.Stderr, "Error converting Order Value: %v\n", err)
						os.Exit(1)
					}

					if newOrd != nil {
						newOrd.Value = IDData

					} else {
						fmt.Fprintf(os.Stderr, "value data(%d) detected without an active current order struct instance. Possibly malformed XML.", IDData)
						os.Exit(1)
					}
				}
			}
		}

		page++
		productAPILocation = fmt.Sprintf("https://store-velgoi8q0k.mybigcommerce.com/api/v2/orders?limit=250&page=%d", page)
		req, _ = http.NewRequest("GET", productAPILocation, nil)
		req.Header.Add("Authorization", `Basic dGVzdDoyNTI1ZGY1NjQ3N2Y1OGU1ODY4YzI0MGVlNTIyOGIwYjVkNDM2N2M0`)
		req.Header.Add("Content-Type", `application/json`)
		apiDataResp, err = client.Do(req)
	}

	var b strings.Builder
	b.WriteString(fmt.Sprintf("<table><thead><tr><th>CUSTOMER</th><th>ORDERS PLACED</th></tr></thead><tbody>"))
	for _, c := range custMap {
		b.WriteString(fmt.Sprintf("<tr><td><a href='/customer/?id=%d'>%s %s</a></td><td>%d</td></tr>", c.ID, c.FirstName, c.LastName, c.Total))
	}

	b.WriteString("</tbody></table>")
	fmt.Fprintf(w, b.String())
}

// handle customer page request for a specific customer (e.g. https://localhost/customer/<id>)
func customerHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "text/html; charset=utf-8") // send hypertext to the browser (responsewriter headers are plaintext by default)
	// custID, err := strconv.Atoi(r.URL.Path[10:])               // strip '/customer/' from the URL path name and acquire customer ID
	// if err != nil {
	// 	// invalid customer ID
	// 	fmt.Fprintf(w, "Invalid customer ID: %s | <a href='/'>back</a>", r.URL.Path[10:])
	// 	return
	// }

	// capture Customer ID from URL
	urlVals := r.URL.Query()
	custIDStr, exists := urlVals["id"]
	// fmt.Println(custIDStr)
	if !exists || len(custIDStr) != 1 {
		fmt.Fprintf(w, "Invalid request | <a href='/'>back</a>")
		return
	}

	custID, err := strconv.Atoi(custIDStr[0])
	if err !=  nil {
		fmt.Fprintf(w, "Invalid customer ID | <a href='/'>back</a>")
		return
	}

	client := &http.Client{} // HTTP client to make API requests
	
	// tokens of interest in the customer XML object
	CUST := "customer"
	ID := "id"
	CUSTFN := "first_name"
	CUSTLN := "last_name"

	var newCust *Customer
	customers := []*Customer{}

	customerAPILocation := fmt.Sprintf("https://store-velgoi8q0k.mybigcommerce.com/api/v2/customers/%d", custID)
	req, _ := http.NewRequest("GET", customerAPILocation, nil)
	req.Header.Add("Authorization", `Basic dGVzdDoyNTI1ZGY1NjQ3N2Y1OGU1ODY4YzI0MGVlNTIyOGIwYjVkNDM2N2M0`)
	req.Header.Add("Content-Type", `application/json`)
	apiDataResp, err := client.Do(req)
	fmt.Println(apiDataResp.StatusCode)

	if err == nil && apiDataResp.StatusCode == http.StatusOK {
		fmt.Println("Getting customer detail for CustID:", custID)
		if err != nil {
			// cannot get data - exit and show results so far
			fmt.Fprintf(os.Stderr, "Error fetching data from API URL (%s): %v\n", customerAPILocation, err)
			apiDataResp.Body.Close()
		}

		if apiDataResp.StatusCode != http.StatusOK {
			fmt.Fprintf(os.Stderr, "API request to URL %s failed: %s\n", customerAPILocation, apiDataResp.Status)
			apiDataResp.Body.Close()
		}

		// decode read XML data body
		dec := xml.NewDecoder(apiDataResp.Body)

		var stack []string // we'll use a string slice as a stack data structure to pop on/off start/end elements as we read through the XML data body's tokens

		for {
			// get next XML token
			token, err := dec.Token()

			// handle any errors
			if err == io.EOF {
				// reached end of data
				break
			} else if err != nil {
				fmt.Fprintf(os.Stderr, "Error reading XML data body token: %v\n", err)
				os.Exit(1)
			}

			// switch statement to take selective action based on the current token (start, end or data)
			switch token := token.(type) {
			case xml.StartElement:
				stack = append(stack, token.Name.Local)

				if len(stack) > 0 && stack[len(stack)-1] == CUST {
					// start of a new <work> in XML data: create a new Work instance and pop in to the list of all customers
					newCust = createCustomer()
				}

			case xml.EndElement:
				// XML end element: pop off stack, and finalize current in-memory work object

				// check if there are already XML opening tags stored in stack - if not, we've encountered a closing tag without an opening tag
				if len(stack) <= 0 {
					fmt.Fprintf(os.Stderr, "Attempting to pop an element(%s) without any on stack - possibly malformed XML\n", token.Name.Local)
					os.Exit(1)
				}

				elementPopped := stack[len(stack)-1]
				stack = stack[:len(stack)-1]

				// check for XML consistency - if every end element should have had a corresponding start element
				if elementPopped != token.Name.Local {
					fmt.Fprintf(os.Stderr, "Closing element %s without matching opener (%s) - possibly malformed XML\n", elementPopped, token.Name.Local)
					os.Exit(1)
				}

				if elementPopped == CUST {
					customers = append(customers, newCust)
					newCust = nil
				}

			case xml.CharData:
				// XML data - populate the current work object based on XML data token (e.g. ID, model, make etc.)
				if len(stack) > 0 && stack[len(stack)-1] == ID {
					IDData, err := strconv.Atoi(strings.TrimSpace(string(token)))

					if err != nil {
						fmt.Fprintf(os.Stderr, "Error converting Order ID: %v\n", err)
						os.Exit(1)
					}

					if newCust != nil {
						newCust.ID = IDData

						if IDData != custID {
							fmt.Fprintf(os.Stderr, "Received customer ID (%d) different from requested ID (%d)", IDData, custID)
							os.Exit(1)
						}

					} else {
						fmt.Fprintf(os.Stderr, "ID data(%d) detected without an active current customer struct instance. Possibly malformed XML.", IDData)
						os.Exit(1)
						continue
					}
				}

				if len(stack) > 0 && stack[len(stack)-1] == CUSTFN {
					fn := strings.TrimSpace(string(token))

					if newCust != nil {
						newCust.FirstName = fn
					} else {
						fmt.Fprintf(os.Stderr, "First name (customer) (%s) detected without an active current customer struct instance. Possibly malformed XML.", fn)
						os.Exit(1)
					}
				}

				if len(stack) > 0 && stack[len(stack)-1] == CUSTLN {
					ln := strings.TrimSpace(string(token))

					if newCust != nil {
						newCust.LastName = ln
					} else {
						fmt.Fprintf(os.Stderr, "Last name (order) (%s) detected without an active current order struct instance. Possibly malformed XML.", ln)
						os.Exit(1)
					}
				}
			}
		}
	} else {
		fmt.Fprintf(w, "Invalid request ksdaksdj | <a href='/'>back</a>")
		return
	}

	if len(customers) != 1 {
		fmt.Fprintf(w, "%d customers returned for one ID (%d) | <a href='/'>back</a>", len(customers), custID)
		return
	}

	page := 1                // pagination value for iterating through orders API
	custLifeTimeVal := 0.0   // order total for this customer

	// initialize XML token values for parsing orders XML
	ORDER := "order"
	ID = "id"
	CUST_ID := "customer_id"
	VALUE := "total_ex_tax"
	DATE := "date_created"
	STATUS := "status"	

	// placeholder for each order object read and list of orders
	var newOrd *Order
	orders := []*Order{}

	// get data from first page of orders for this customer
	ordersAPILocation := fmt.Sprintf("https://store-velgoi8q0k.mybigcommerce.com/api/v2/orders?customer_id=%d&limit=250&page=%d", custID, page)
	req, _ = http.NewRequest("GET", ordersAPILocation, nil)
	req.Header.Add("Authorization", `Basic dGVzdDoyNTI1ZGY1NjQ3N2Y1OGU1ODY4YzI0MGVlNTIyOGIwYjVkNDM2N2M0`)
	req.Header.Add("Content-Type", `application/json`)
	apiDataResp, err = client.Do(req)

	for err == nil && apiDataResp.StatusCode == http.StatusOK {
		fmt.Println("Order API call (customers)")
		if err != nil {
			// cannot get data - exit and show results so far
			fmt.Fprintf(os.Stderr, "Error fetching data from API URL (%s): %v\n", ordersAPILocation, err)
			apiDataResp.Body.Close()
		}

		if apiDataResp.StatusCode != http.StatusOK {
			fmt.Fprintf(os.Stderr, "API request to URL %s failed: %s\n", ordersAPILocation, apiDataResp.Status)
			apiDataResp.Body.Close()
		}

		// decode read XML data body
		dec := xml.NewDecoder(apiDataResp.Body)

		var stack []string // we'll use a string slice as a stack data structure to pop on/off start/end elements as we read through the XML data body's tokens

		for {
			// get next XML token
			token, err := dec.Token()

			// handle any errors
			if err == io.EOF {
				// reached end of data
				break
			} else if err != nil {
				fmt.Fprintf(os.Stderr, "Error reading XML data body token: %v\n", err)
				os.Exit(1)
			}

			// switch statement to take selective action based on the current token (start, end or data)
			switch token := token.(type) {
			case xml.StartElement:
				stack = append(stack, token.Name.Local)

				if len(stack) > 0 && stack[len(stack)-1] == ORDER {
					// start of a new <work> in XML data: create a new Work instance and pop in to the list of all customers
					newOrd = createOrder()
				}

			case xml.EndElement:
				// XML end element: pop off stack, and finalize current in-memory work object

				// check if there are already XML opening tags stored in stack - if not, we've encountered a closing tag without an opening tag
				if len(stack) <= 0 {
					fmt.Fprintf(os.Stderr, "Attempting to pop an element(%s) without any on stack - possibly malformed XML\n", token.Name.Local)
					os.Exit(1)
				}

				elementPopped := stack[len(stack)-1]
				stack = stack[:len(stack)-1]

				// check for XML consistency - if every end element should have had a corresponding start element
				if elementPopped != token.Name.Local {
					fmt.Fprintf(os.Stderr, "Closing element %s without matching opener (%s) - possibly malformed XML\n", elementPopped, token.Name.Local)
					os.Exit(1)
				}

				if elementPopped == ORDER {
					orders = append(orders, newOrd)
					newOrd = nil
				}

			case xml.CharData:
				// XML data - populate the current work object based on XML data token (e.g. ID, model, make etc.)
				if len(stack) > 0 && stack[len(stack)-1] == ID {
					IDData, err := strconv.Atoi(strings.TrimSpace(string(token)))

					if err != nil {
						fmt.Fprintf(os.Stderr, "Error converting Order ID: %v\n", err)
						os.Exit(1)
					}

					if newOrd != nil {
						newOrd.ID = IDData

					} else {
						fmt.Fprintf(os.Stderr, "ID data(%d) detected without an active current order struct instance. Possibly malformed XML (page = %d).", IDData, page)
						os.Exit(1)
						continue
					}
				}

				if len(stack) > 0 && stack[len(stack)-1] == CUST_ID {
					IDData, err := strconv.Atoi(strings.TrimSpace(string(token)))

					if err != nil {
						fmt.Fprintf(os.Stderr, "Error converting Order Customer ID: %v\n", err)
						os.Exit(1)
					}

					if newOrd != nil {
						newOrd.CustID = IDData

					} else {
						fmt.Fprintf(os.Stderr, "Customer ID data(%d) detected without an active current order struct instance. Possibly malformed XML.", IDData)
						os.Exit(1)
					}
				}

				if len(stack) > 0 && stack[len(stack)-1] == VALUE {
					IDData, err := strconv.ParseFloat(strings.TrimSpace(string(token)), 64)

					if err != nil {
						fmt.Fprintf(os.Stderr, "Error converting Order Value: %v\n", err)
						os.Exit(1)
					}

					if newOrd != nil {
						custLifeTimeVal += IDData
						newOrd.Value = IDData

					} else {
						fmt.Fprintf(os.Stderr, "value data(%d) detected without an active current order struct instance. Possibly malformed XML.", IDData)
						os.Exit(1)
					}
				}

				if len(stack) > 0 && stack[len(stack)-1] == DATE {
					orderDate := strings.TrimSpace(string(token))

					if newOrd != nil {
						newOrd.Date = orderDate
					} else {
						fmt.Fprintf(os.Stderr, "Order date(%s) detected without an active current order struct instance. Possibly malformed XML.", orderDate)
						os.Exit(1)
					}
				}

				if len(stack) > 0 && stack[len(stack)-1] == STATUS {
					status := strings.TrimSpace(string(token))

					if newOrd != nil {
						newOrd.Status = status
					} else {
						fmt.Fprintf(os.Stderr, "Oder status(%s) detected without an active current order struct instance. Possibly malformed XML.", status)
						os.Exit(1)
					}
				}

				if len(stack) > 0 && stack[len(stack)-1] == CUSTFN {
					fn := strings.TrimSpace(string(token))

					if newOrd != nil {
						newOrd.CustFN = fn
					} else {
						fmt.Fprintf(os.Stderr, "First name (order) (%s) detected without an active current order struct instance. Possibly malformed XML.", fn)
						os.Exit(1)
					}
				}

				if len(stack) > 0 && stack[len(stack)-1] == CUSTLN {
					ln := strings.TrimSpace(string(token))

					if newOrd != nil {
						newOrd.CustLN = ln
					} else {
						fmt.Fprintf(os.Stderr, "Last name (order) (%s) detected without an active current order struct instance. Possibly malformed XML.", ln)
						os.Exit(1)
					}
				}
			}
		}

		// read next page of orders API for this customer
		page++
		ordersAPILocation = fmt.Sprintf("https://store-velgoi8q0k.mybigcommerce.com/api/v2/orders?customer_id=%d&limit=250&page=%d", custID, page)
		req, _ = http.NewRequest("GET", ordersAPILocation, nil)
		req.Header.Add("Authorization", `Basic dGVzdDoyNTI1ZGY1NjQ3N2Y1OGU1ODY4YzI0MGVlNTIyOGIwYjVkNDM2N2M0`)
		req.Header.Add("Content-Type", `application/json`)
		apiDataResp, err = client.Do(req)
	}

	// prepare orders HTML for web interface display
	var b strings.Builder
	b.WriteString(fmt.Sprintf("<h2>%s %s</h2><b>Lifetime Value: $%.2f</b> | <a href='/'>back</a><br /><br /><table border='1'><thead><tr><th>ORDER ID</th><th>PLACED</th><th>STATUS</th><th>TOTAL</th></tr></thead><tbody>", customers[0].FirstName, customers[0].LastName, custLifeTimeVal))
	for _, c := range orders {
		b.WriteString(fmt.Sprintf("<tr><td>%d</td><td>%s</td><td>%s</td><td>$%.2f</td></tr>", c.ID, c.Date, c.Status, c.Value))
	}

	if len(orders) == 0 {
		b.WriteString(fmt.Sprintf("<tr><td colspan='4'><i>No orders found for this customer.</i></td></tr>"))
	}

	b.WriteString("</tbody></table>")
	fmt.Fprintf(w, b.String()) //	sned HTML to browser

	return
}

// type struct representing a customer
type Customer struct {
	ID        int
	FirstName string
	LastName  string
	Total     int
}

// type struct representing an order
type Order struct {
	ID     int
	CustID int
	Value  float64
	Date   string
	Status string
	CustFN string
	CustLN string
}

// constructor for orders struct
func createOrder() *Order {
	var w Order
	w.ID = -1
	w.CustID = -1
	w.Value = 0.0

	return &w
}

// constructor for customers struct
func createCustomer() *Customer {
	var w Customer
	w.ID = -1
	w.FirstName = ""
	w.LastName = ""

	return &w
}

/* ◬ */
