package main

import "fmt"

func HackPassword(hash string) string {
	pswd, _ := GuessDigit(hash)
	return fmt.Sprintf("%04d", pswd)
}
