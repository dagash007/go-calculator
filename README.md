# Go Calculator Project

## Project Overview
This project implements a simple calculator in Go, focusing on basic arithmetic operations like division and squaring. The goal is to familiarize with Go's unit testing and benchmarking capabilities.

## Instructions

### How to Run the Program
1. Clone the repository https://github.com/dagash007/go-calculator.git
2. Navigate to the project directory https://github.com/dagash007/go-calculator

### How to Execute Tests (gitbash)
Run the following command in the terminal:

go test


### How to Execute Benchmarks
go test -bench=


### Expected Output of Tests
Tests should validate the correctness of the Divide and Square functions.
Expected test output should indicate if the functions pass all cases.


### Test and Benchmark Design
TestDivide: Checks for normal division and error handling when dividing by zero.
TestSquare: Verifies the correctness for positive and negative integers.
BenchmarkDivide: Benchmarks division with a consistent non-zero denominator.
BenchmarkSquare: Benchmarks squaring of a representative integer input.


### Contributions
Team Member Peter: Implemented calculator.go.
Team Member Abdul: Wrote calculator_test.go and benchmarks.
Team Member Abdul: Prepared the documentation.
