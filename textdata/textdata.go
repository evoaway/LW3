package textdata

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
)

type Team struct {
	name            string
	city            string
	numberOfPlayers int
	points          int
}

func check(e error) {
	if e != nil {
		panic(e)
	}
}
func PrintData(filename string) {
	inputFile, err := os.Open(filename)
	check(err)
	defer inputFile.Close()
	fmt.Println("--- Вміст файлу ---")
	scanner := bufio.NewScanner(inputFile)
	for scanner.Scan() {
		line := scanner.Text()
		team := strings.Split(line, "|")
		fmt.Println("Назва: " + team[0] + " Місто: " + team[1] +
			" Кількість гравців: " + team[2] + " Кількість очок: " + team[3])
	}
}
func DeleteTeam(filename string, point int) error {
	var result string
	inputFile, err := os.Open(filename)
	check(err)
	defer inputFile.Close()
	scanner := bufio.NewScanner(inputFile)
	for scanner.Scan() {
		line := scanner.Text()
		words := strings.Split(line, "|")
		num, err := strconv.Atoi(words[3])
		if err != nil {
			return fmt.Errorf("Помилка конвертування рядка в  число: %v", err)
		}
		if num >= point {
			result += line + "\n"
		}
	}
	if err := scanner.Err(); err != nil {
		return fmt.Errorf("Помилка зчитування файлу: %v", err)
	}
	file, err := os.Create(filename)
	check(err)
	defer file.Close()
	file.WriteString(result)
	return nil
}
func TeamDataEnter() Team {
	var team Team
	fmt.Println("Введіть дані про команду")
	fmt.Print("Назва: ")
	if _, err := fmt.Scan(&team.name); err != nil {
		log.Print("  Scan for Name failed, due to ", err)
	}
	fmt.Print("Місто: ")
	if _, err := fmt.Scan(&team.city); err != nil {
		log.Print("  Scan for City failed, due to ", err)
	}
	fmt.Print("Кількість гравців: ")
	if _, err := fmt.Scan(&team.numberOfPlayers); err != nil {
		log.Print("  Scan for number of players failed, due to ", err)
	}
	fmt.Print("Кількість очок: ")
	if _, err := fmt.Scan(&team.points); err != nil {
		log.Print("  Scan for points failed, due to ", err)
	}
	return team
}
func ScanTeams() []Team {
	fmt.Printf("Введіть число команд: ")
	var size int
	if _, err := fmt.Scan(&size); err != nil {
		panic("Неправильно введена кількість!")
	}
	employers := make([]Team, size)
	for i := 0; i < size; i++ {
		employers[i] = TeamDataEnter()
	}
	return employers
}
func TeamToBytes(teams []Team) []byte {
	var result string
	for _, tm := range teams {
		str := tm.name + "|" + tm.city + "|" + strconv.Itoa(tm.numberOfPlayers) + "|" + strconv.Itoa(tm.points) + "\n"
		result += str
	}
	return []byte(result)
}
func AddToBegin(filename string, teams []Team) {
	data, err := os.ReadFile("text.txt")
	check(err)
	newTeams := TeamToBytes(teams)
	newData := append(newTeams, data...)
	file, err := os.Create(filename)
	check(err)
	defer file.Close()
	file.Write(newData)
}
