package main

import (
	"LW3/numberdata"
	"LW3/textdata"
	"fmt"
)

func Task1_1Demo() {
	inputFilename := "numbers.txt"
	outputFilename := "modules.txt"

	err, count := numberdata.Task1_1(inputFilename, outputFilename)
	if err != nil {
		fmt.Println("Error:", err)
	}
	fmt.Println("Кількість квадратів непарних чисел:", count)
}
func Task1_2Demo() {
	inputFilename := "numbers.txt"
	outputFilename := "sqrt.txt"

	err, sum := numberdata.Task1_2(inputFilename, outputFilename)
	if err != nil {
		fmt.Println("Error:", err)
	}
	fmt.Println("Модуль суми компонент файлу:", sum)
}
func Task1_3Demo() {
	inputFilename := "numbers.txt"
	outputFilename := "sort.txt"

	err, sum := numberdata.Task1_3(inputFilename, outputFilename)
	if err != nil {
		fmt.Println("Error:", err)
	}
	fmt.Println("Квадрат добутку компонент файлу", sum)
}
func Task2Demo() {
	filename := "text.txt"
	teams := textdata.ScanTeams()
	textdata.AddToBegin(filename, teams)
	textdata.PrintData(filename)
	teams = textdata.ScanTeams()
	textdata.AddToBegin(filename, teams)
	textdata.PrintData(filename)
	fmt.Println("Після видалення деяких команд:")
	err := textdata.DeleteTeam(filename, 12)
	if err != nil {
		fmt.Println("Error:", err)
	}
	textdata.PrintData(filename)
}
func Task3Demo() {
	resultFile := "result.txt"
	c1 := make(chan Result)
	c2 := make(chan Result)
	go solveEquationByIterations(f, dt, -3, 3, c1)
	go solveEquationByChord(f, -3, 3, c2)
	res1, res2 := <-c1, <-c2
	WriteToFile(resultFile, res1, res2)
}
func main() {
	//Task1_1Demo()
	//Task1_2Demo()
	//Task1_3Demo()
	//Task2Demo()
	Task3Demo()
}
