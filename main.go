package main

import (
	"fmt"
	"os"
	"path/filepath"
)

func main() {
	if len(os.Args) < 3 {
		fmt.Println("Usage: govault <file> <password>")
		return
	}

	target := os.Args[1]
	password := os.Args[2]

	// check file existence
	if _, err := os.Stat(target); os.IsNotExist(err) {
		fmt.Printf("❌ Error: The file '%s' does not exist.\n", target)
		return
	}

	ext := filepath.Ext(target)

	if ext == ".govault" {
		if err := Decrypt(target, password); err != nil {
			fmt.Printf("❌ Decrypt Error: %v\n", err)
		}
	} else {
		// don't overwrite if vault exists already
		vaultPath := target + ".govault"
		if _, err := os.Stat(vaultPath); err == nil {
			fmt.Printf("❌ Error: A locked version ('%s') already exists. Unlock it first.\n", vaultPath)
			return
		}

		if err := Encrypt(target, password); err != nil {
			fmt.Printf("❌ Encrypt Error: %v\n", err)
		}
	}
}
