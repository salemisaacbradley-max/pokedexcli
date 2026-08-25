package main

import ("fmt" 
		"bufio"
		"os")

func main() {
    scanner := bufio.NewScanner(os.Stdin)
	initCommands()
	for ;; {
		fmt.Print("Pokedex > ")
		result := scanner.Scan()
		if result != true {
			fmt.Errorf("Hey, dingbat. Write something next time.")
		}
		scanned := scanner.Text()
		cleaned := cleanInput(scanned)
		if len(cleaned) == 0 {
			continue
		}
		command, ok := commands[cleaned[0]]
		if ok == true {
			command.callback()
		} else {
			fmt.Print("Unknown Command")
		}
	}
}
