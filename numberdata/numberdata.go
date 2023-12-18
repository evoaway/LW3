package numberdata

import (
	"bufio"
	"fmt"
	"math"
	"os"
	"sort"
	"strconv"
)

func IsOddSquare(num int) bool {
	root := int(math.Sqrt(float64(num)))
	if root*root == num && root%2 != 0 {
		return true
	}
	return false
}
func Abs(x int) int {
	if x < 0 {
		return -x
	}
	return x
}

func Task1_1(inputFilename, outputFilename string) (error, int) {
	count := 0
	inputFile, err := os.Open(inputFilename)
	if err != nil {
		return err, count
	}
	defer inputFile.Close()
	outputFile, err := os.Create(outputFilename)
	if err != nil {
		return err, count
	}
	defer outputFile.Close()
	scanner := bufio.NewScanner(inputFile)
	for scanner.Scan() {
		line := scanner.Text()
		num, err := strconv.Atoi(line)
		if err != nil {
			return fmt.Errorf("Помилка конвертування рядка в число: %v", err), count
		}
		if IsOddSquare(num) {
			count++
		}
		result := strconv.Itoa(Abs(num))
		_, err = outputFile.WriteString(result + "\n")
		if err != nil {
			return fmt.Errorf("Помилка запису result до файлу: %v", err), count
		}
	}
	if err := scanner.Err(); err != nil {
		return fmt.Errorf("Помилка зчитування файлу: %v", err), count
	}
	return nil, count
}
func Task1_2(inputFilename, outputFilename string) (error, int) {
	sum := 0
	inputFile, err := os.Open(inputFilename)
	if err != nil {
		return err, sum
	}
	defer inputFile.Close()
	outputFile, err := os.Create(outputFilename)
	if err != nil {
		return err, sum
	}
	defer outputFile.Close()
	scanner := bufio.NewScanner(inputFile)
	for scanner.Scan() {
		line := scanner.Text()
		num, err := strconv.Atoi(line)
		if err != nil {
			return fmt.Errorf("Помилка конвертування рядка в число: %v", err), Abs(sum)
		}
		sum += num
		result := fmt.Sprintf("%.2f", math.Sqrt(float64(num)))
		_, err = outputFile.WriteString(result + "\n")
		if err != nil {
			return fmt.Errorf("Помилка запису result до файлу: %v", err), Abs(sum)
		}
	}
	if err := scanner.Err(); err != nil {
		return fmt.Errorf("Помилка зчитування файлу: %v", err), Abs(sum)
	}
	return nil, Abs(sum)
}
func check(e error) {
	if e != nil {
		panic(e)
	}
}
func Task1_3(inputFilename, outputFilename string) (error, int) {
	prod := 1
	var nums []int
	inputFile, err := os.Open(inputFilename)
	check(err)
	defer inputFile.Close()
	outputFile, err := os.Create(outputFilename)
	check(err)
	defer outputFile.Close()
	scanner := bufio.NewScanner(inputFile)
	for scanner.Scan() {
		line := scanner.Text()
		num, err := strconv.Atoi(line)
		if err != nil {
			return fmt.Errorf("Помилка конвертування рядка в число: %v", err), int(math.Sqrt(float64(prod)))
		}
		nums = append(nums, num)
		prod *= num
	}
	if err := scanner.Err(); err != nil {
		return fmt.Errorf("Помилка зчитування файлу: %v", err), int(math.Sqrt(float64(prod)))
	}
	sort.Ints(nums)
	for _, data := range nums {
		result := fmt.Sprintf("%d\n", data)
		_, err = outputFile.WriteString(result)
		if err != nil {
			return fmt.Errorf("Помилка запису result до файлу: %v", err), int(math.Sqrt(float64(prod)))
		}
	}
	return nil, int(math.Sqrt(float64(prod)))
}
