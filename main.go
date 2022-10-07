package main

import (
	"bufio"
	"errors"
	"fmt"
	"os"
)

func openFile() ([]byte, error) {
	if len(os.Args) < 2 {
		return nil, errors.New("please provide a file name")
	}
	data, err := os.ReadFile(os.Args[1])
	if err != nil {
		fmt.Println("Can't read file:", os.Args[1])
		panic(err)
	}
	return data, nil
}
func saveFile(data []byte) error {
	if len(os.Args) < 2 {
		return errors.New("please provide a file name")
	}
	err := os.WriteFile(os.Args[1], data, 0644)
	if err != nil {
		fmt.Println("Can't write file:", os.Args[1])
		panic(err)
	}
	return nil
}
func printFile() {
	data, _ := openFile()
	printTemp(data)

}

func printTemp(data []byte) {
	fmt.Println(string(data))
}
func edit(data []byte) []byte {
	for {
		fmt.Print("\033[H\033[2J")
		printTemp(data)
		scanner := bufio.NewScanner(os.Stdin)
		for scanner.Scan() {
			line := scanner.Text()
			if line == ":eq" {
				return data
			}
			if line == "\n" {
				data = append(data, line...)
				break
			}

		}
	}
}

func main() {
	data, err := openFile()

	if err != nil {
		panic(err)
	}
	for {
		fmt.Println("command mode")
		fmt.Print("exit - :q\n", "edit - :e\n")
		var input string
		fmt.Scanln(&input)
		if input == ":q" {
			fmt.Print("save file? [Y][n]:")
			fmt.Scanln(&input)
			if input == "n" || input == "N" {
				return
			}
			saveFile(data)
			fmt.Println("saved")
			return
		}
		if input == ":e" {

			data = edit(data)
		}
	}
}
