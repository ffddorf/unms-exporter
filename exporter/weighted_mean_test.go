package exporter

import (
	"fmt"
	"math/rand"
	"sort"
	"testing"
	"time"

	"github.com/ffddorf/unms-exporter/models"
	"github.com/stretchr/testify/assert"
)

// assert that actual = expected ± ε
func assertWithin(t *testing.T, expected, ε, actual float64, msgAndArgs ...interface{}) {
	t.Helper()

	if actual <= expected-ε || expected+ε <= actual {
		assert.Fail(t, fmt.Sprintf("%f should be within %f of %f", actual, ε, expected), msgAndArgs...)
	}
}

func TestWeightedMean_trivialCases(t *testing.T) {
	assert.EqualValues(t, 0, weightedMean(0, 0, nil))
	assert.EqualValues(t, 0, weightedMean(10, 0, nil))
	assert.EqualValues(t, 0, weightedMean(0, 10, nil))
}

func TestWeightedMean_singleItem(t *testing.T) {
	assert.EqualValues(t, 0, weightedMean(0, 10, []*models.CoordinatesXY{{X: 0, Y: 10}}))
	assert.EqualValues(t, 10, weightedMean(0, 10, []*models.CoordinatesXY{{X: 10, Y: 10}}))

}

func weightedMeanTestList() models.ListOfCoordinates {
	list := models.ListOfCoordinates{
		{X: 20, Y: 10},
		{X: 30, Y: 10},
		{X: 40, Y: 10},
		{X: 50, Y: 10},
		{X: 60, Y: 10},
		{X: 70, Y: 10},
		{X: 80, Y: 10},
	}

	// shuffle list. order shall not matter
	rand.Seed(time.Now().Unix())
	sort.Slice(list, func(int, int) bool { return rand.Float64() < 0.5 })

	return list
}

func TestWeightedMean_largeList(t *testing.T) {
	list := weightedMeanTestList()

	// if all values are 10, then the weight is irrelevant
	assertWithin(t, 10, 1e-6, weightedMean(25, 25, list), "equal limits")
	assertWithin(t, 10, 1e-6, weightedMean(20, 80, list), "boring case")
	assertWithin(t, 10, 1e-6, weightedMean(80, 20, list), "inverted limits")
	assertWithin(t, 10, 1e-6, weightedMean(0, 100, list), "limits larger than data")
}

func TestWeightedMean_outlier(t *testing.T) {
	// add outlier
	list := append(weightedMeanTestList(), &models.CoordinatesXY{X: 10, Y: 1000})
	assertWithin(t, 10, 1e-6, weightedMean(25, 25, list[0:7]), "ignoring outlier")

	// m = 1/(55-5) = 0.02, tmin = 5
	// 	i  list[i]     weight             weighted value
	//	-  ---------   -----------------  --------------
	//	5  (70, 10)                  1.0    1.0*10 =  10
	//	6  (80, 10)                  1.0    1.0*10 =  10
	//	7  (10, 1000)  0.02*(10-5) = 0.1  0.1*1000 = 100
	//	                        sum: 2.1        sum: 120
	// mean = 120/2.1 = 57.1428(5)
	assertWithin(t, 57.14285, 5e-5, weightedMean(5, 55, list[5:]), "including outlier")
}
