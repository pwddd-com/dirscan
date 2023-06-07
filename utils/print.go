package utils

import "github.com/gookit/color"

func PError(message string) {
	color.Red.Printf("[X] %s\n", message)
}

func PInfo(message string) {
	color.Blue.Printf("[-] %s\n", message)
}

func PWarn(message string) {
	color.Yellow.Printf("[!] %s\n", message)
}
