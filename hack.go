package main

import "fmt"

func HackPassword(hash string) string {
	pswd, _ := GuessDigit(hash)
	if pswd == -1 {
		return fmt.Sprintf("%d", pswd)
	}
	return fmt.Sprintf("%04d", pswd)
}
