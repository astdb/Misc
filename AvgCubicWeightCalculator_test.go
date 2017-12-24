package main

import (
	"math"
	"testing"
)

func TestAvgCubicWeight(t *testing.T) {
	F64_TOLERANCE := 0.000001

	// func avgCubicWeight(weightTotal, productTotal float64) float64
	if avgCubicWeight(200, 0) != 200 {
		t.Error(`avgCubicWeight(200, 0) != 200`)
	}

	if avgCubicWeight(0, 0) != 0 {
		t.Error(`avgCubicWeight(0, 0) != 0`)
	}

	if avgCubicWeight(-1, -1) != -1 {
		t.Error(`avgCubicWeight(-1, -1) != -1`)
	}

	val1 := avgCubicWeight(200, 2)
	if val1 != 100 {
		t.Errorf(`avgCubicWeight(200, 2) != 100 (got %f)`, val1)
	}

	val2 := avgCubicWeight(350, 35)
	if val2 != 10 {
		t.Errorf(`avgCubicWeight(350, 35) != 10 (got %f)`, val2)
	}

	val3 := avgCubicWeight(3216841, 354)
	if math.Abs(val3-9087.121468927) > F64_TOLERANCE {
		// if val3 != 9087.121468927 {
		t.Errorf(`avgCubicWeight(3216841, 354) - 9087.121468927 > %f (got %f)`, F64_TOLERANCE, val3)
	}

	val4 := avgCubicWeight(321986454, 321825)
	if math.Abs(val4-1000.501682591) > F64_TOLERANCE {
		// if val3 != 9087.121468927 {
		t.Errorf(`avgCubicWeight(321986454, 321825) - 1000.501682591 > %f (got %f)`, F64_TOLERANCE, val3)
	}
}

func TestCubicWeight(t *testing.T) {
	var cubicWeightConversionFactor float64
	cubicWeightConversionFactor = 250
	F64_TOLERANCE := 0.000001

	val1, err := cubicWeight(&ProductSize{1.0, 1.0, 1.0}, cubicWeightConversionFactor)
	if err != nil {
		t.Errorf(`cubicWeight(&ProductSize{1.0, 1.0, 1.0}, %f) returned error %v`, err)
	}

	if math.Abs(val1 - 0.00025) > F64_TOLERANCE {
		t.Errorf(`cubicWeight(&ProductSize{1.0, 1.0, 1.0}, %f) - 1000.501682591 > %f (got %f)`, cubicWeightConversionFactor, F64_TOLERANCE, val1)
	}

	_, err = cubicWeight(&ProductSize{-20.5, 30.0, 40.0}, cubicWeightConversionFactor)
	if err == nil {
		t.Errorf(`cubicWeight(&ProductSize{-20.5, 30.0, 40.0}, %f) did not return an error`)
	}

	_, err = cubicWeight(&ProductSize{0, 0, 0}, cubicWeightConversionFactor)
	if err == nil {
		t.Errorf(`cubicWeight(&ProductSize{0, 0, 0}, %f) did not return an error`)
	}

	val2, err := cubicWeight(&ProductSize{40, 20, 30}, cubicWeightConversionFactor)
	if err != nil {
		t.Errorf(`cubicWeight(&ProductSize{40, 20, 30}, %f) returned error %v`, err)
	}

	if math.Abs(val2 - 6.0) > F64_TOLERANCE {
		t.Errorf(`cubicWeight(&ProductSize{40, 20, 30}, %f) - 6.0 > %f (got %f)`, cubicWeightConversionFactor, F64_TOLERANCE, val2)
	}
}
