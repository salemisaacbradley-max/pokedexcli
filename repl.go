package main

import ("strings"
		"fmt" 
		"bufio"
		"os")

func cleanInput(text string) []string {
	lowerText := strings.ToLower(text)
	mons := strings.Fields(lowerText)
	return mons
}

func REPL(cfg *config) {
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
			command.callback(cfg)
		} else {
			fmt.Print("Unknown Command\n")
		}
	}
}