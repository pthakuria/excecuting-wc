package main

import (
	"bytes"
	"fmt"
	"log"
	"os/exec"
)

const ShellToUse = "bash"

func Shellout(command string) (error, string, string) {
	// 
	var stdout bytes.Buffer
	var stderr bytes.Buffer
	
	cmd := exec.Command(ShellToUse, "-c", command)
	cmd.Stdout = &stdout // standard output file
	cmd.Stderr = &stderr // this is the standard error file, 
						 // where all the exceptions are stored.
	err := cmd.Run()     // returns nil if execution complete,
						 // else returns 1.
	return err, stdout.String(), stderr.String()
}

func main() {
	err, out, errout := Shellout("ls | wc -l")
	if err != nil {
		log.Printf("error: %v\n", err)
	}
	fmt.Println("--- stdout ---")
	fmt.Println(out)
	fmt.Println("--- stderr ---")
	fmt.Println(errout)
}
