package main

import "testing"

func BenchmarkIntegerAdd20(b *testing.B) {
	IntegerAdd20()
}

func BenchmarkFloatAdd20(b *testing.B) {
	FloatAdd20()
}

func BenchmarkFloatMultiply20(b *testing.B) {
	FloatMultiply20()

}
