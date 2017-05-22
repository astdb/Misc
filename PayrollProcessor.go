package main

import (
	"encoding/csv"
	"fmt"
	"io"
	"math"
	"os"
	"regexp"
	"strconv"
	"strings"
)

var TAX_CONFIG_FILE = "TAX_CONFIG.csv"
var TAX_BRACKETS []*IncomeTaxBracket

func main() {
	if len(os.Args) <= 1 {
		fmt.Println("Usage: > go run PayrollProcessor.go <inputfile>")
		return
	}

	TAX_BRACKETS, err := readTaxBracketsConfig(TAX_CONFIG_FILE)
	if err != nil {
		// error reading tax bracket config - quit
		fmt.Printf("Error reading tax brackets config: %v", err)
		return
	}

	for _, v := range TAX_BRACKETS {
		v.Print();
	}

	payrollRecords, err := readPayrollRecords(os.Args[1])
	if err != nil {
		fmt.Printf("Error reading payroll records: %v", err)
	}

	for k, v := range payrollRecords {
		fmt.Printf("%2d. ", k)

		if v != nil {
			v.Print()
		} else {
			fmt.Printf("Nil record object\n")
		}
	}
}

func readPayrollRecords(inputFile string) ([]*PayrollRecord, error) {
	fileHandle, err := os.Open(inputFile)
	if err != nil {
		return nil, err
	}
	defer fileHandle.Close()

	csvReader := csv.NewReader(fileHandle)
	records := []*PayrollRecord{}

	for {
		row, err := csvReader.Read()
		if err != nil {
			if err == io.EOF {
				err = nil
			}

			// TODO: returning at the first erroneous row? can we skip over til EOF?
			return records, err
		}

		// TODO: consider sending entire input row to recod creation function
		if len(row) < 5 { // need at least five fields in a valid record, also check for empty fields
			// return or collect error and move to next record
			// return records, fmt.Errorf("Minimum number of fields not met in input <%s>\n", row)
			continue // read next row
		}

		newPayrollRecord, err := createPayrollRecord(row[0], row[1], row[2], row[3], row[4])
		if err != nil {
			// throw error
			// fmt.Errorf("Error creating payroll record with input %s\n", row)
			// return records, fmt.Errorf("Minimum number of fields not met in input <%s>\n", row)
			continue // read next row
		}

		records = append(records, newPayrollRecord)
	}
}

type PayrollRecord struct {
	FirstName    string
	LastName     string
	AnnualSalary float64
	SuperRate    float64
	PaymentDate  string
	Valid        bool   //	indicates if the record object is valid
	ErrorStr     string // if Valid == false, contains the input data from the input file leading to invalid object
	TaxBrackets  []*IncomeTaxBracket
}

func (rec *PayrollRecord) fullName() string {
	return rec.FirstName + " " + rec.LastName
}

func (rec *PayrollRecord) payPeriod() string {
	return rec.PaymentDate
}

func (rec *PayrollRecord) grossIncome() float64 {
	return round(rec.AnnualSalary / 12)
}

func (rec *PayrollRecord) incomeTax() (float64, error) {
	// iterate through tax brackets and find the right tax percentage for this salary
	perc := 0.0
	lump := 0.0
	abv := 0.0
	set := false
	for _, brac := range TAX_BRACKETS {
		if brac.Upper == 0 {
			// top tax bracket
			if rec.AnnualSalary >= brac.Lower {
				perc = brac.Percent
				lump = brac.Lump
				abv = brac.Above
				set = true
			}

		} else {
			if rec.AnnualSalary >= brac.Lower && rec.AnnualSalary <= brac.Upper {
				perc = brac.Percent
				lump = brac.Lump
				abv = brac.Above
				set = true
			}
		}
	}

	if !set {
		// no fitting bracket found for this salary
		return -1.0, fmt.Errorf("No fittnig tax bracket was found for salary amount %f", rec.AnnualSalary)
	}

	percentageTax := ((rec.AnnualSalary - abv) * perc) / 100
	return round(percentageTax + lump), nil
}

func (rec *PayrollRecord) netIncome() (float64, error) {
	gross := rec.grossIncome()
	tax, err := rec.incomeTax()

	if err != nil {
		return -1.0, err
	}

	if tax > gross {
		return -1.0, fmt.Errorf("Taxed amount (%f) larger than gross income (%f)", tax, gross)
	}

	return round(gross - tax), nil
}

func (rec *PayrollRecord) superAmount() (float64, error) {
	if rec.SuperRate < 0 || rec.SuperRate > 50 {
		return -1.0, fmt.Errorf("Invalid super rate (%f)", rec.SuperRate)
	}

	super := (rec.grossIncome() * rec.SuperRate) / 100
	return round(super), nil
}

func (rec *PayrollRecord) Print() {
	if rec.Valid {
		fmt.Printf("%s, %s, %.2f, %.2f%%, %s, %f\n", rec.FirstName, rec.LastName, rec.AnnualSalary, rec.SuperRate, rec.PaymentDate, rec.grossIncome())
	} else {
		fmt.Printf("Invalid record %s\n", rec.ErrorStr)
	}
}

func createPayrollRecord(firstname_str string, lastname_str string, annualsalary_str string, superrate_str string, paymentdate_str string) (*PayrollRecord, error) {
	var newRecord PayrollRecord

	// prepare input
	FirstName := strings.TrimSpace(firstname_str)
	LastName := strings.TrimSpace(lastname_str)
	AnnualSalary, err_sal := strconv.ParseFloat(annualsalary_str, 64)
	SuperRate := strings.TrimSpace(superrate_str)
	PaymentDate := strings.TrimSpace(paymentdate_str)

	// validate [TODO: may improve by returning specific error, consider ignoring erroneous record and reading ahead]
	// extract numeric super percentage
	rexp, _ := regexp.Compile("^[0-9]+")
	SuperRate_f, err_sr := strconv.ParseFloat(rexp.FindString(SuperRate), 64)

	if FirstName == "" || LastName == "" || (err_sal != nil) || (err_sr != nil) || AnnualSalary <= 0 || SuperRate_f < 0 || SuperRate_f > 50 {
		newRecord.Valid = false
		newRecord.ErrorStr = fmt.Sprintf("Invalid input record: [%s] [%s] [%s] [%s] [%s]", firstname_str, lastname_str, annualsalary_str, superrate_str, paymentdate_str)
		return &newRecord, fmt.Errorf("Invalid data in payroll input record\n")
	}

	newRecord.FirstName = FirstName
	newRecord.LastName = LastName
	newRecord.AnnualSalary = AnnualSalary
	newRecord.SuperRate = SuperRate_f
	newRecord.PaymentDate = PaymentDate
	newRecord.Valid = true

	return &newRecord, nil
}

func round(n float64) float64 {
	n_floor := math.Floor(n)
	n_dec := n - n_floor

	if n_dec >= .50 {
		// round up
		return n_floor + 1.0
	}

	// round down
	return n_floor
}

func createOutputFile(inFileName) error {
	if strings.TrimSpace(inFileName) == "" {
		return fmt.Errorf("Invalid input filename")
	}

	outFileName := strings.Split(strings.TrimSpace(inFileName), ".")[0] + "-out.txt"
	f, err := os.Create(outFileName)
	
	if err != nil {
		return fmt.Errorf("Error creating outputfile")
	}

	defer f.Close()

	for scanner.Scan() {
		line := scanner.Text()
		date := strings.TrimSpace(r1.FindString(line))
		ext := strings.TrimSpace(r2.FindString(line))

		if len(date) > 0 && len(ext) > 0 {
			// write line to file
			_, err := f.WriteString(r1.FindString(line) + "\t" + r2.FindString(line) + "\n")
			logErr(err)
		}
	}

	f.Sync()

	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}
}

// takes a salary amount and returns percentage, lump payment and value above which the percebntage should be calculated
// func getTaxPercent(sal float64) (float64, float64, float64) {
// 	taxmap := map[float64]float64{0: 18200, }
// }

type IncomeTaxBracket struct {
	Lower   float64 // lower salary limit of tax bracket
	Upper   float64 // upper salary limit of tax bracket
	Percent float64 // percentage tax to be levied above a certain threshold
	Lump    float64 // lump sum to be paid for this bracket, if any
	Above   float64 // threshold above which percentage tax has to be paid
}

func (itb *IncomeTaxBracket) Print()  {
	fmt.Printf("[Lower: $%.2f, Upper: $%.2f, Percent: %.2f%%, Lump Sum: $%.2f, Above: $%.2f]\n", itb.Lower, itb.Upper, itb.Percent, itb.Lump, itb.Above)
}

func readTaxBracketsConfig(inputFile string) ([]*IncomeTaxBracket, error) {
	fileHandle, err := os.Open(inputFile)
	if err != nil {
		return nil, err
	}
	defer fileHandle.Close()

	csvReader := csv.NewReader(fileHandle)
	brackets := []*IncomeTaxBracket{}

	i := 0
	prev_upper := 0.0
	for {
		row, err := csvReader.Read()
		if err != nil {
			if err == io.EOF {
				err = nil
			}

			return brackets, err
		}

		// TODO: consider sending entire input row to recod creation function
		if len(row) < 5 { // need at least five fields in a valid record, also check for empty fields
			// return or collect error and move to next record
			return nil, fmt.Errorf("readTaxBrackets(): Minimum number of fields not met in input <%s>\n", row)
		}

		lower, err_low := strconv.ParseFloat(strings.TrimSpace(row[0]), 64)

		top := false // flag indicating topmost tax bracket
		upp := strings.TrimSpace(row[1])
		upper := 0.0
		var err_upp error
		if upp != "" {
			// upper field could be empty if this is the top bracket
			upper, err_upp = strconv.ParseFloat(strings.TrimSpace(row[1]), 64)
		} else {
			// set flag indicating uppermost bracket
			top = true
		}

		percent, err_perc := strconv.ParseFloat(strings.TrimSpace(row[2]), 64)
		lump, err_lump := strconv.ParseFloat(strings.TrimSpace(row[3]), 64)
		threshold, err_thr := strconv.ParseFloat(strings.TrimSpace(row[4]), 64)

		if err_low != nil || err_upp != nil || err_perc != nil || err_lump != nil || err_thr != nil {
			fmt.Printf("%v\n", err_thr)
			return nil, fmt.Errorf("readTaxBrackets(): Error reading tax bracket config record: <%s>\n", row)
		}

		if lower >= upper && !top {
			return nil, fmt.Errorf("readTaxBrackets(): Lower limit >= upper limit in input <%s>\n", row)
		}

		if i == 0 {
			if lower != 0 {
				return nil, fmt.Errorf("readTaxBrackets(): First bracket lower limit != 0 in input <%s>\n", row)
			}
		} else {
			if lower <= prev_upper {
				return nil, fmt.Errorf("readTaxBrackets(): Current bracket's lower limit <= previous upper limit in input <%s>\n", row)
			}
		}

		prev_upper = upper
		newTaxBracket := &IncomeTaxBracket{lower, upper, percent, lump, threshold}
		brackets = append(brackets, newTaxBracket)
		i++

		if top {
			break
		}
	}

	if i == 0 {
		return nil, fmt.Errorf("readTaxBrackets(): No valid tax brackets found\n")
	}

	return brackets, nil
}
