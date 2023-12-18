package main

import (
	"fmt"
	"math"
	"os"
)

type Result struct {
	x float64
	i int
}

func f(x float64) float64 {
	return x*x - x - 5
}
func dt(x float64) float64 {
	return 2*x - 1
}
func solveEquationByIterations(function func(x float64) float64, derivative func(x float64) float64, a, b float64, c chan Result) {
	defer close(c)
	var xNext float64
	maxIter := 100
	finalIter := maxIter
	tolerance := 0.00001
	x := (a + b) / 2
	for i := 0; i < maxIter; i++ {
		xNext = x - function(x)/derivative(x)
		if math.Abs(xNext-x) < tolerance {
			finalIter = i
			break
		}
		x = xNext
	}
	c <- Result{x: xNext, i: finalIter}
}
func solveEquationByChord(function func(x float64) float64, a, b float64, c chan Result) {
	defer close(c)
	maxIter := 100
	finalIter := maxIter
	tolerance := 0.00001
	var tmp, xNext float64
	for i := 0; i < maxIter; i++ {
		tmp = xNext
		xNext = b - function(b)*(a-b)/(function(a)-function(b))
		a = b
		b = tmp
		if math.Abs(xNext-b) < tolerance {
			finalIter = i
			break
		}
	}
	c <- Result{x: xNext, i: finalIter}
}
func check(e error) {
	if e != nil {
		panic(e)
	}
}
func WriteToFile(filename string, resultIterMethod Result, resultChordMethod Result) {
	defer func() {
		if panicValue := recover(); panicValue != nil {
			fmt.Println(panicValue)
		}
	}()
	data := fmt.Sprintf("x1 = %.5f, кількість ітерацій = %d\nx1 = %.5f, кількість ітерацій = %d\n",
		resultIterMethod.x, resultIterMethod.i, resultChordMethod.x, resultChordMethod.i)
	if resultIterMethod.x-resultChordMethod.x > 0 {
		data += "Корінь при обчислений ітераційним метод більший, ніж методом хорд\n"
	} else if resultIterMethod.x-resultChordMethod.x < 0 {
		data += "Корінь при обчислений методом хорд більший, ніж ітераційним\n"
	} else {
		data += "Корені рівні\n"
	}
	if resultIterMethod.i-resultChordMethod.i > 0 {
		data += "Ітераційним метод виконав більшу кількість ітерацій\n"
	} else if resultIterMethod.x-resultChordMethod.x < 0 {
		data += "Метод хорд виконав більшу кількість ітерацій\n"
	} else {
		data += "Методи виконали однакову кількість іткрацій\n"
	}
	file, err := os.Create(filename)
	check(err)
	defer file.Close()
	file.WriteString(data)
}
