package main

import (
	"bufio"
	"fmt"
	"io/ioutil"
	"math/rand"
	"os"
	"strconv"
	"strings"
)

type Word struct {
	polish  string
	english string
}

type Words []Word

var wordsList Words

func (words *Words) Add(word Word) {
	*words = append(*words, word)
}
func (words *Words) DeleteByIndex(index int) {
	*words = append((*words)[:index], (*words)[index+1:]...)
}

func (words *Words) ReadByIndex(index int) (word Word) {
	word = (*words)[index]
	return
}
func (words *Words) ReadAll() Words {
	return *words
}

func (words *Words) GetLength() (length int) {
	length = len(*words)
	return
}

func handleAddWord(english string, polish string) {
	newWord := Word{polish, english}
	wordsList.Add(newWord)
	fmt.Println("New word added!")
}

func handleShowAllWords() {
	for i, word := range wordsList.ReadAll() {
		fmt.Println(fmt.Sprintf("%v. %v - %v", i+1, word.english, word.polish))
	}
}

func handleShowRandomWord() {
	word := getRandomWord()

	fmt.Println(fmt.Sprintf("%v - %v", word.english, word.polish))
}

func handleDeleteWord(indexAsString string) {
	index, err := strconv.Atoi(indexAsString)
	if err != nil {
		handleError()
	}
	wordsList.DeleteByIndex(index - 1)
}

func handleExit() {
	defer os.Exit(0)
	fmt.Println("Bye!")
}

func handleError() {
	fmt.Println("Error occurred")
	os.Exit(2)
}

func getRandomWord() (word Word) {
	max, min := wordsList.GetLength(), 0
	randomIndex := rand.Intn(max-min) + min
	word = wordsList.ReadByIndex(randomIndex)
	return
}

func loadWordsFromFile() {
	data, err := ioutil.ReadFile("words.txt")

	if err != nil {
		handleError()
	}
	stringWords := strings.Split(strings.ReplaceAll(string(data), "\r\n", "\n"), "\n")

	for _, sw := range stringWords {
		if sw == "" {
			break
		}
		splitWord := strings.Split(sw, " ")
		word := Word{splitWord[1], splitWord[0]}
		wordsList.Add(word)
	}

}

func saveToFile() {
	file, err := os.Create("words.txt")

	if err != nil {
		handleError()
	}

	defer func(file *os.File) {
		err := file.Close()
		if err != nil {
			handleError()
		}
	}(file)

	wordsString := ""

	for _, word := range wordsList.ReadAll() {
		wordsString += word.english + " " + word.polish + "\n"
	}

	_, err = file.WriteString(wordsString)

	if err != nil {
		handleError()
	}

}

var scanner = bufio.NewScanner(os.Stdin)

func playQuiz() {
	fmt.Println("Game started! Type end to stop playing")
	var input string
	for input != "end" {
		word := getRandomWord()
		fmt.Print("Translate this word - " + word.polish + " : ")
		scanner.Scan()
		input = scanner.Text()
		if input == word.english {
			fmt.Println("Correct!")
		} else if input != "end" {
			fmt.Println("Incorrect. Correct answer: " + word.english)
		}
	}

}

func main() {
	loadWordsFromFile()

	for {
		fmt.Print("\n>>> ")
		scanner.Scan()
		input := scanner.Text()

		inputArr := strings.Split(input, " ")

		args := inputArr[1:]

		command := inputArr[0]

		switch command {
		case "add":
			handleAddWord(args[0], args[1])
		case "showall":
			handleShowAllWords()
		case "showrandom":
			handleShowRandomWord()
		case "exit":
			handleExit()
		case "playquiz":
			playQuiz()
		case "delete":
			handleDeleteWord(args[0])
		default:
			fmt.Println("Command doesn't exist")
		}
		saveToFile()
	}
}
