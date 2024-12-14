/*
Copyright © 2021-2024 Cameron Esfahani
*/

package utilities

import (
	"reflect"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestPrimeFactors(t *testing.T) {
	type testCase struct {
		number          int
		expectedFactors []int
	}

	testCases := []testCase{
		{2, []int{2}},
		{3, []int{3}},
		{4, []int{2, 2}},
		{7, []int{7}},
		{800, []int{2, 2, 2, 2, 2, 5, 5}},
	}

	for _, test := range testCases {
		assert.Equal(t, test.expectedFactors, PrimeFactors(test.number))
	}
}

func TestIntPow(t *testing.T) {
	type testCase struct {
		base          int
		exponent      int
		expectedPower int
	}

	testCases := []testCase{
		{2, 0, 1},
		{3, 3, 27},
		{2, 16, 65536},
		{8, 4, 64 * 64},
	}

	for _, test := range testCases {
		assert.Equal(t, test.expectedPower, IntPow(test.base, test.exponent))
	}
}

func TestHistogram(t *testing.T) {
	type testCase struct {
		list        []int
		expectedMap map[int]int
	}

	testCases := []testCase{
		{[]int{1, 2, 3, 4}, map[int]int{1: 1, 2: 1, 3: 1, 4: 1}},
		{[]int{1, 2, 3, 4, 3, 5}, map[int]int{1: 1, 2: 1, 3: 2, 4: 1, 5: 1}},
		{[]int{1, 2, 3, 4, 3, 5, -1, -1, -1}, map[int]int{1: 1, 2: 1, 3: 2, 4: 1, 5: 1, -1: 3}},
	}

	for _, test := range testCases {
		assert.True(t, reflect.DeepEqual(MakeHistogram(test.list), test.expectedMap))
	}
}

func TestConcatenate(t *testing.T) {
	type testCase struct {
		a                     int64
		b                     int64
		expectedConcatenation int64
	}
	testCases := []testCase{
		{
			a:                     1,
			b:                     3,
			expectedConcatenation: 13,
		},
		{
			a:                     1234,
			b:                     5678,
			expectedConcatenation: 12345678,
		},
		{
			a:                     98,
			b:                     1,
			expectedConcatenation: 981,
		},
		{
			a:                     70,
			b:                     0,
			expectedConcatenation: 700,
		},
		{
			a:                     1,
			b:                     0,
			expectedConcatenation: 10,
		},
		{
			a:                     100,
			b:                     0,
			expectedConcatenation: 1000,
		},
		{
			a:                     0,
			b:                     1,
			expectedConcatenation: 1,
		},
		{
			a:                     0,
			b:                     101,
			expectedConcatenation: 101,
		},
		{
			a:                     1,
			b:                     1,
			expectedConcatenation: 11,
		},
		{
			a:                     800,
			b:                     8,
			expectedConcatenation: 8008,
		},
		{
			a:                     10,
			b:                     10,
			expectedConcatenation: 1010,
		},
		{
			a:                     100,
			b:                     100,
			expectedConcatenation: 100100,
		},
		{
			a:                     11,
			b:                     111222333,
			expectedConcatenation: 11111222333,
		},
	}

	for _, test := range testCases {
		assert.Equal(t, test.expectedConcatenation, Concatenate(test.a, test.b))
	}
}

func TestGenerateSlopeIntercept(t *testing.T) {
	type testCase struct {
		pointA    Point2D
		pointB    Point2D
		expectedM float64
		expectedB float64
	}
	testCases := []testCase{
		{
			pointA:    NewPoint2D(0, 0),
			pointB:    NewPoint2D(2, 2),
			expectedM: 1.0,
			expectedB: 0.0,
		},
		{
			pointA:    NewPoint2D(4, 3),
			pointB:    NewPoint2D(5, 5),
			expectedM: 2.0,
			expectedB: -5.0,
		},
	}

	for _, test := range testCases {
		m, b := CalculateSlopeIntercept(test.pointA, test.pointB)
		assert.Equal(t, test.expectedM, m)
		assert.Equal(t, test.expectedB, b)
	}
}

func TestDigitCount(t *testing.T) {
	type testCase struct {
		number             int
		expectedDigitCount int
	}

	testCases := []testCase{
		{0, 1},
		{1, 1},
		{9, 1},
		{10, 2},
		{99, 2},
		{123456789, 9},
	}

	for _, test := range testCases {
		assert.Equal(t, test.expectedDigitCount, DigitCount(test.number))
	}
}
