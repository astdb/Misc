/*
AvgCubicWeightCalculator connects to a remote API responding in JSON via HTTP. It implements data structures required to marshal
the JSON data into memory and perform subsequent processing. The program is invoked via command line and accepts arguments specifing API endpoint
URL, product category name, and a cubic weight conversion factor - and outputs the average cubic weight for all products found in the
specified product category.

*/

package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"net/url"
	"os"
	"strconv"
)

func main() {
	// verify correct program invocation - API URL, product category name and cubic weight conversion factor expected
	if len(os.Args) < 4 {
		fmt.Println("Usage: $> go run AvgCubicWeighCalculator.go <api_endpoint_url> <category_name> <cubic_weight_cnonversion_factor>")
		return
	}

	// read in command line arguments
	productAPILocation := os.Args[1]
	categoryName := os.Args[2]
	cubicWeightConversionFactor, err := strconv.ParseFloat(os.Args[3], 64) // convert into float (and handle any resulting error) - command line args come as strings
	if err != nil {
		log.Fatalf("Invalid cubic weigth conversion factor: %s", err)
	}

	// parse API endpoint URL
	apiURL, err := url.Parse(productAPILocation)
	if err != nil {
		log.Fatalf("Invalid API URL: %s", err)
	}
	apiProtocol := apiURL.Scheme // e,g, http
	apiHost := apiURL.Host       // e.g. http://wp8m3he1wt.s3-website-ap-southeast-2.amazonaws.com
	nextPath := apiURL.Path      // e.g. /api/products/1

	// declare map to track visited API pages - this will be used to detect if a previous page is provided as a next and starts a loop
	visitedEndPoints := map[string]bool{}

	productTotal := 0.0     // total number of product items found in the specified category
	cubicWeightTotal := 0.0 // total cubic weight of products in specified category

	// visit each page of the API
	for nextPath != "" {
		// check if this path was previously visited - if yes, exit. if not, store in list of pages visited
		_, visited := visitedEndPoints[nextPath]
		if visited {
			break
		}
		visitedEndPoints[nextPath] = true

		// get JSON data response from specified/next API location
		apiDataResp, err := http.Get(productAPILocation)
		if err != nil {
			// cannot get data - exit and show results so far
			fmt.Fprintf(os.Stderr, "Error fetching data from API URL (%s): %v\n", productAPILocation, err)
			apiDataResp.Body.Close()
			break
		}

		if apiDataResp.StatusCode != http.StatusOK {
			fmt.Fprintf(os.Stderr, "API request to URL %s failed: %s\n", productAPILocation, apiDataResp.Status)
			apiDataResp.Body.Close()
			break
		}

		// get JSON data response
		var productsPage ProductAPIPage                               // struct for storing this API page
		err = json.NewDecoder(apiDataResp.Body).Decode(&productsPage) // marshal data into struct
		if err != nil {
			fmt.Fprintf(os.Stderr, "JSON marshalling failed: %s\n", err)
			apiDataResp.Body.Close()
			break
		}

		// get list of products on this page and process data for any products found in the specified category
		products := productsPage.Objects
		for _, v := range products {
			if v.Category == categoryName {
				productTotal++
				cubicWeightTotal += cubicWeight(v.Size, cubicWeightConversionFactor)
			}
		}

		// get next API page path and construct location URL
		nextPath = productsPage.Next
		if nextPath != "" {
			productAPILocation = fmt.Sprintf("%s://%s%s", apiProtocol, apiHost, nextPath)
			// fmt.Printf("Visiting %s\n", productAPILocation)
		}

		apiDataResp.Body.Close()
	}

	avgCubicWeight := cubicWeightTotal / productTotal

	// format output
	fmt.Printf("Product Category: \"%s\"\n", categoryName)
	fmt.Printf("Total Products in Category: %.0f\n", productTotal)
	fmt.Printf("Average Cubic Weight: %.3fkg\n", avgCubicWeight)
}

// struct type definition representing an API page
type ProductAPIPage struct {
	Objects []*Product // list of products
	Next    string     // next page
}

// struct type representing a product
type Product struct {
	Category string
	Title    string
	Weight   float64
	Size     *ProductSize
}

// struct type representing a composite product size (length, width and height)
type ProductSize struct {
	Width  float64
	Length float64
	Height float64
}

func cubicWeight(pz *ProductSize, cubicWeightConversionFactor float64) float64 {
	return (pz.Width / 100) * (pz.Length / 100) * (pz.Height / 100) * cubicWeightConversionFactor
}
