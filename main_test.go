package main

import (
	"reflect"
	"testing"
)

type Int2IntTestPair struct {
	input1 int
	input2 int
	output int
}

func TestSoma(t *testing.T) {
	var tests = []Int2IntTestPair{
		{0, 1, 2},
		{1, 1, 2},
	}

	for _, pair := range tests {
		result := soma(pair.input1, pair.input2)
		if !reflect.DeepEqual(result, pair.output) {
			t.Errorf(
				"For sum of %v and %v:\n expected: %v\n but got: %v",
				pair.input1, pair.input2, pair.output, result)
		}
	}
}
