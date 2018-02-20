package main

import (
	"math/rand"
	"time"

	"github.ibm.com/Content-Delivery-Org/goverfitting/examples"
)

func main() {
	// Seed global random
	rand.Seed(time.Now().UTC().UnixNano())

	//run example
	examples.RunLinearRegression()
}
