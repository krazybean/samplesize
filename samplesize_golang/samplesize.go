package samplesize_golang

import (
	"errors"
	"math"
)

var z_score = map[float64]float64{0.80: 1.28, 0.85: 1.44, 0.90: 1.65, 0.95: 1.96, 0.99: 2.56}

func SampleSize(confidence float64, margin float64, distribution float64) (float64, error) {
	zScoreIndex := z_score[confidence]
	if zScoreIndex <= 0 {
		return 0.0, errors.New("usage of confidence should be [80, 85, 90, 95, 99]")
	}
	sample := distribution * (1 - distribution) / (math.Pow(margin/zScoreIndex, 2))
	return sample, nil
}

func TrueSample(population int, confidence float64, margin float64, distribution float64) (float64, error) {
	ss, err := SampleSize(confidence, margin, distribution)
	if err != nil {
		return 0.0, err
	}
	ts := (ss * float64(population)) / (ss + float64(population) - 1)
	return ts, nil
}
