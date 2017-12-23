// This program connects to a remote API responding in JSON via HTTP. The code implements data structures required to marshal the JSON data
// into memory and subsequent processing. The program is invoked via command line and accepts arguments specifing API endpoint URL,
// product category name, and cubic weight conversion factor - and outputs the average cubic weight for all products found in the
// specified product category.

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
	if len(os.Args) < 4 {
		fmt.Println("Usage: $> go run AvgCubicWeighCalculator.go <api_endpoint_url> <category_name> <cubic_weight_cnonversion_factor>")
		return
	}

	// read in command line arguments: API URL, category, and cubic weight conversion factor
	productAPILocation := os.Args[1]
	categoryName := os.Args[2]
	cubicWeightConversionFactor, err := strconv.ParseFloat(os.Args[3], 64)
	if err != nil {
		log.Fatalf("Invalid cubic weigth conversion factor: %s", err)
	}

	// parse URL
	apiURL, err := url.Parse(productAPILocation)
	if err != nil {
		log.Fatalf("Invalid API URL specified: %s", err)
	}
	// fmt.Printf("API Host: %s\n", apiURL.Host)
	apiProtocol := apiURL.Scheme
	apiHost := apiURL.Host
	nextURL := apiURL.Path
	visitedEndPoints := map[string]bool{}

	productTotal := 0.0     // total of product found in the specified category
	cubicWeightTotal := 0.0 // total of cubic weight of products in spcified category

	for nextURL != "" {
		_, visited := visitedEndPoints[nextURL]
		if visited {
			break
		}
		visitedEndPoints[nextURL] = true

		// get JSON data response from specified/next API location
		apiDataResp, err := http.Get(productAPILocation)
		if err != nil {
			fmt.Fprintf(os.Stderr, "Error fetching JSON works data from specified API URL (%s): %v\n", productAPILocation, err)
			return
		}

		defer apiDataResp.Body.Close()

		if apiDataResp.StatusCode != http.StatusOK {
			log.Fatalf("API request failed: %s", apiDataResp.Status)
		}

		// decode read JSON data response
		var productsPage ProductAPIPage
		err = json.NewDecoder(apiDataResp.Body).Decode(&productsPage)

		if err != nil {
			log.Fatalf("JSON marshalling failed: %s", err)
		}
		products := productsPage.Objects

		for _, v := range products {
			if v.Category == categoryName {
				productTotal++
				cubicWeightTotal += cubicWeight(v.Size, cubicWeightConversionFactor)
			}
		}

		nextURL = productsPage.Next
		if nextURL != "" {
			productAPILocation = fmt.Sprintf("%s://%s%s", apiProtocol, apiHost, nextURL)
			fmt.Println(productAPILocation)
		}
	}

	avgCubicWeight := cubicWeightTotal / productTotal
	fmt.Println(avgCubicWeight)
}

type ProductAPIPage struct {
	Objects []*Product
	Next    string
}

type Product struct {
	Category string
	Title    string
	Weight   float64
	Size     *ProductSize
}

type ProductSize struct {
	Width  float64
	Length float64
	Height float64
}

func printSize(pz *ProductSize) string {
	return fmt.Sprintf("[Width: %f, Length: %f, Height: %f]", pz.Width, pz.Length, pz.Height)
}

func cubicWeight(pz *ProductSize, cubicWeightConversionFactor float64) float64 {
	return (pz.Width / 100) * (pz.Length / 100) * (pz.Height / 100) * cubicWeightConversionFactor
}
