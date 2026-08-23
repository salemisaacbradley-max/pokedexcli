package main

import "strings"

func cleanInput(text string) []string {
	lowerText := strings.ToLower(text)
	mons := strings.Fields(lowerText)
	return mons
}