package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

type ClassInfo struct {
	name       string
	instructor string
	numCredits int
	isOnline   bool
}

func main() {
	inputReader := bufio.NewReader(os.Stdin)
	for {
		fmt.Println("Project 1 Menu")
		fmt.Println("1) Print every Xth line of a file")
		fmt.Println("2) Count bytes and characters in some text")
		fmt.Println("3) Struct demo")
		fmt.Println("4) List of structs demo")
		fmt.Println("5) Quit")
		fmt.Print("Pick an option (1-5): ")
		answer, _ := inputReader.ReadString('\n')
		answer = strings.TrimSpace(answer)

		switch answer {
		case "1":
			printEveryXLine(inputReader)
		case "2":
			countBytesAndChars(inputReader)
		case "3":
			structDemo()
		case "4":
			structListDemo()
		case "5":
			fmt.Println("Bye!")
			return
		}
	}
}

func getFile(name string) []string {
	contents, err := os.ReadFile(name)
	if err != nil {
		fmt.Println("File reading error", err)
		return nil
	}
	bigString := string(contents)
	allLines := strings.Split(bigString, "\n")
	return allLines
}

func printEveryXLine(inputReader *bufio.Reader) {
	fmt.Print("Enter the file name: ")
	fileName, _ := inputReader.ReadString('\n')
	fileName = strings.TrimSpace(fileName)
	allLines := getFile(fileName)
	fmt.Print("Enter a number X (prints every Xth line): ")
	xString, _ := inputReader.ReadString('\n')
	xString = strings.TrimSpace(xString)
	x, _ := strconv.Atoi(xString)

	for i, line := range allLines {
		if (i+1)%x == 0 {
			fmt.Println(line)
		}
	}
}

func countBytesAndChars(inputReader *bufio.Reader) {
	fmt.Print("Type or paste some text: ")
	userText, _ := inputReader.ReadString('\n')
	userText = strings.TrimSpace(userText)
	numBytes := len(userText)
	numChars := len([]rune(userText))
	fmt.Printf("Bytes: %d\n", numBytes)
	fmt.Printf("Characters: %d\n", numChars)
}

func structDemo() {
	myClass := ClassInfo{}
	myClass.name = "COMP 415"
	myClass.instructor = "Santore"
	myClass.numCredits = 3
	myClass.isOnline = false
	fmt.Println("Single struct demo")
	fmt.Print("Struct before calling the function: ", myClass, "\n")
	changeClass(myClass)
	fmt.Print("Struct after the function returns: ", myClass, "\n")
	fmt.Println("The original still says COMP 415 because the function only changed a copy of the struct.")
}

func changeClass(class ClassInfo) {
	class.name = "COMP 490"
	fmt.Print("  Struct inside the function after changing the name: ", class, "\n")
}

func structListDemo() {
	classList := make([]ClassInfo, 0)

	firstClass := ClassInfo{}
	firstClass.name = "COMP 415"
	firstClass.instructor = "Santore"
	firstClass.numCredits = 3
	firstClass.isOnline = false
	classList = append(classList, firstClass)

	secondClass := ClassInfo{}
	secondClass.name = "MATH 200"
	secondClass.instructor = "Puyu"
	secondClass.numCredits = 4
	secondClass.isOnline = true
	classList = append(classList, secondClass)

	fmt.Println("List of structs demo")
	fmt.Print("List before calling the function: ", classList, "\n")
	changeSecondClass(classList)
	fmt.Print("List after the function returns: ", classList, "\n")
	fmt.Println("The 2nd class stayed changed this time because the function edited the actual list, not a copy.")
}

func changeSecondClass(classes []ClassInfo) {
	classes[1].name = "COMP 350"
	classes[1].instructor = "Abdul"
	fmt.Print("  List inside the function after changing the 2nd class: ", classes, "\n")
}
