package main

import (
	"fmt"
	"io"
	"os"
	"strings"

	"filippo.io/age"
	"golang.org/x/term"
)

// Encrypt locks a file using a password
func Encrypt(filename string, password string) error {

	if password == "" {
		password = getPassword("Enter password: ")
	}

	recipient, _ := age.NewScryptRecipient(password)

	out, err := os.Create(filename + ".govault")
	if err != nil {
		return err
	}
	defer out.Close()

	// Wrap our output file in an age writer
	w, err := age.Encrypt(out, recipient)
	if err != nil {
		return err
	}

	in, _ := os.Open(filename)
	defer in.Close()

	// Stream the file through the encryption engine
	if _, err := io.Copy(w, in); err != nil {
		return err
	}
	w.Close()

	return os.Remove(filename)
}

// Decrypt unlocks a .govault file
func Decrypt(filename string, password string) error {

	if password == "" {
		password = getPassword("Enter password: ")
	}

	identity, _ := age.NewScryptIdentity(password)

	f, _ := os.Open(filename)
	defer f.Close()

	// Decrypt the stream
	r, err := age.Decrypt(f, identity)
	if err != nil {
		return err
	}

	outPath := strings.TrimSuffix(filename, ".govault")
	out, _ := os.Create(outPath)
	defer out.Close()

	if _, err := io.Copy(out, r); err != nil {
		return err
	}

	return os.Remove(filename)
}

func getPassword(prompt string) string {
	fmt.Print(prompt)
	// This hides the password as you type
	bytePassword, _ := term.ReadPassword(int(os.Stdin.Fd()))
	fmt.Println()
	return string(bytePassword)
}
