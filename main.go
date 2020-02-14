package main

// Using the modified Simple Size calculation based on:
// http://fluidsurveys.com/university/calculating-right-survey-sample-size/

import (
	"fmt"
	"github.com/krazybean/samplesize/samplesize_golang"
)

const MARGIN_OF_ERROR = 0.05
const CONFIDENCE = 0.95
const POPULATION = 1000
const DISTRIBUTION = 0.5

func main() {
	fmt.Println(samplesize_golang.TrueSample(POPULATION, CONFIDENCE, MARGIN_OF_ERROR, DISTRIBUTION))
}
