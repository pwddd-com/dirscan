package utils

import (
	"fmt"
	"github.com/gookit/color"
)

func PError(message string) {
	color.Red.Printf("[X] %s\n", message)
}

func PInfo(message string) {
	color.Blue.Printf("[-] %s\n", message)
}

func PWarn(message string) {
	color.Yellow.Printf("[!] %s\n", message)
}

func PDebug(message string) {
	color.Gray.Printf("[DEBUG] %s\n", message)
}

func PBanner() {
	banner := `
 ██████████   █████ ███████████    █████████    █████████    █████████   ██████   █████
░░███░░░░███ ░░███ ░░███░░░░░███  ███░░░░░███  ███░░░░░███  ███░░░░░███ ░░██████ ░░███ 
 ░███   ░░███ ░███  ░███    ░███ ░███    ░░░  ███     ░░░  ░███    ░███  ░███░███ ░███ 
 ░███    ░███ ░███  ░██████████  ░░█████████ ░███          ░███████████  ░███░░███░███ 
 ░███    ░███ ░███  ░███░░░░░███  ░░░░░░░░███░███          ░███░░░░░███  ░███ ░░██████ 
 ░███    ███  ░███  ░███    ░███  ███    ░███░░███     ███ ░███    ░███  ░███  ░░█████ 
 ██████████   █████ █████   █████░░█████████  ░░█████████  █████   █████ █████  ░░█████
░░░░░░░░░░   ░░░░░ ░░░░░   ░░░░░  ░░░░░░░░░    ░░░░░░░░░  ░░░░░   ░░░░░ ░░░░░    ░░░░░ 
`
	color.Magenta.Println(banner)

	color.Magenta.Printf("[+] code by %s v%s\n", "沐风", "0.0.1")
	color.Magenta.Printf("[+] https://github.com/marmufeng/dirscan\n")
	fmt.Println()
}
